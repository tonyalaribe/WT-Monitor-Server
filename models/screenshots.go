package models

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/constants"

	"github.com/boltdb/bolt"
)

type Screenshot struct {
	Timestamp time.Time
	URL       string
	UserID    string
}

//Create adds a new user to the database
func (s *Screenshot) Create(db *bolt.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.SCREENSHOTS_BUCKET))
	sBkt, err := b.CreateBucketIfNotExists([]byte(s.UserID))
	if err != nil {
		return err
	}
	t := s.Timestamp.Format(time.RFC3339)

	err = sBkt.Put([]byte(t), []byte(s.URL))

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

//GetTodaysScreenshots returns all screenshots taken today
func (s *Screenshot) GetTodaysScreenshots(db *bolt.DB, user string) (screenshots []Screenshot, err error) {
	tx, err := db.Begin(false)
	if err != nil {
		return screenshots, err
	}
	defer tx.Rollback()

	now := time.Now()
	//currentTime := []byte(now.Format(time.RFC3339))
	today := []byte(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format(time.RFC3339))

	//log.Println(user + "_" + constants.SCREENSHOTS_BUCKET)

	b := tx.Bucket([]byte(constants.SCREENSHOTS_BUCKET))
	sBkt := b.Bucket([]byte(user))

	c := sBkt.Cursor()
	/*
		for k, v := c.Seek(today); k != nil && bytes.Compare(k, currentTime) <= 0; k, v = c.Next() {
			//fmt.Printf("%s: %s\n", k, v)
			x := Screenshot{}
			x.URL = string(v)
			var t time.Time
			t, err = time.Parse(time.RFC3339, string(k))
			if err != nil {
				log.Println(err)
			}

			x.Timestamp = t
			x.UserID = user

			screenshots = append(screenshots, x)

		}
	*/

	for k, v := c.Last(); k != nil && bytes.Compare(k, today) >= 0; k, v = c.Prev() {
		//fmt.Printf("%s: %s\n", k, v)
		x := Screenshot{}
		x.URL = string(v)
		var t time.Time
		t, err = time.Parse(time.RFC3339, string(k))
		if err != nil {
			log.Println(err)
		}

		x.Timestamp = t
		x.UserID = user

		screenshots = append(screenshots, x)

	}
	reversed := []Screenshot{}

	// reverse order
	// and append into new slice
	for i := range screenshots {
		n := screenshots[len(screenshots)-1-i]
		//fmt.Println(n) -- sanity check
		reversed = append(reversed, n)
	}

	return reversed, nil
}

//GetScreenshotsForADay returns all screenshots taken on a particular day
func (s *Screenshot) GetScreenshotsForADay(db *bolt.DB, user string, day time.Time) (screenshots []Screenshot, err error) {

	now := time.Now()
	currentTime := []byte(now.Format(time.RFC3339))
	dayByte := []byte(time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local).Format(time.RFC3339))

	//log.Println(user + "_" + constants.SCREENSHOTS_BUCKET)
	tx, err := db.Begin(false)
	if err != nil {
		return screenshots, err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.SCREENSHOTS_BUCKET))
	sBkt := b.Bucket([]byte(user))

	c := sBkt.Cursor()

	for k, v := c.Seek(dayByte); k != nil && bytes.Compare(k, currentTime) <= 0; k, v = c.Next() {
		//fmt.Printf("%s: %s\n", k, v)
		x := Screenshot{}
		x.URL = string(v)
		var t time.Time
		t, err = time.Parse(time.RFC3339, string(k))
		if err != nil {
			log.Println(err)
		}

		x.Timestamp = t
		x.UserID = user

		screenshots = append(screenshots, x)

	}
	return screenshots, nil
}

func CreateScreenshot(id string, buf []byte) error {

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	storageFolder := filepath.Join("static", id, today.Format(time.RFC3339))

	err := os.MkdirAll(storageFolder, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}

	files, err := ioutil.ReadDir(storageFolder)
	if err != nil {
		log.Println(err)
		return err
	}
	folderCountplusone := fmt.Sprintf("%06d", len(files)+1)

	filename := filepath.Join(storageFolder, folderCountplusone+".png")
	log.Println(err)

	err = ioutil.WriteFile(filename, buf, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	client := Client{Path: config.Get().DB.File}
	err = client.Open()
	if err != nil {
		log.Println(err)
		return err
	}

	defer client.DB.Close()

	data := Screenshot{
		UserID:    id,
		Timestamp: time.Now(),
		URL:       filename,
	}
	//log.Printf("%+v", data)

	err = data.Create(client.DB)
	if err != nil {
		log.Println(err)
		return err
	}

	timelog := UserTimeLog{
		UserID: id,
	}
	err = timelog.LogoutNow(client.DB)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
