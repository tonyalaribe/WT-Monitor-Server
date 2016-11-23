package models

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/constants"

	"path/filepath"

	"github.com/boltdb/bolt"
)

type Archive struct {
	UserID string
}

//Create adds a new user to the database
func (a Archive) Create(db *bolt.DB, date time.Time) error {

	day := date.Format(time.RFC3339)

	archivePath := filepath.Join("archive", a.UserID)
	err := os.MkdirAll(archivePath, os.ModePerm)
	if err != nil {
		return err
	}

	imagePath := filepath.Join("static", a.UserID, day)
	err = os.MkdirAll(imagePath, os.ModePerm)
	if err != nil {
		return err
	}

	completeImagePath := filepath.Join(imagePath, "%06d.png")
	log.Println(completeImagePath)

	completePath := filepath.Join(archivePath, day+".mp4")
	log.Println(completePath)

	ffmpegCmd := exec.Command("ffmpeg", "-framerate", "1/5", "-i", completeImagePath, "-c:v", "libx264", "-r", "30", "-pix_fmt", "yuv420p", completePath)
	err = ffmpegCmd.Run()
	if err != nil {
		return err
	}

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.ARCHIVE_BUCKET))
	aBkt, err := b.CreateBucketIfNotExists([]byte(a.UserID))
	if err != nil {
		return err
	}

	err = aBkt.Put([]byte(day), []byte(completePath))

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

//GetForDate returns the url to the video responsible for a particular day
func (a Archive) GetForDate(db *bolt.DB, date time.Time) (string, error) {
	var resultURL string
	day := date.Format(time.RFC3339)

	tx, err := db.Begin(true)
	if err != nil {
		return resultURL, err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.ARCHIVE_BUCKET))
	aBkt := b.Bucket([]byte(a.UserID))

	resultURLByte := aBkt.Get([]byte(day))

	return string(resultURLByte), nil
}

func worker(id int, jobs <-chan string, results chan<- error) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		archive := Archive{
			UserID: j,
		}
		c := Client{Path: config.Get().DB.TestFile}
		err := c.Open()
		if err != nil {
			results <- err
			return
		}
		defer c.DB.Close()
		now := time.Now().AddDate(0, 0, 0)
		err = archive.Create(c.DB, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
		if err != nil {
			results <- err
			return
		}
		results <- nil
	}
}

func DoArchiveForYesterday() error {
	c := Client{Path: config.Get().DB.TestFile}
	err := c.Open()
	if err != nil {
		log.Println(err)
		return err
	}
	defer c.DB.Close()
	allUsers, err := User{}.GetAll(c.DB)
	if err != nil {
		log.Println(err)
		return err
	}

	maxConcurrentWorkers := 5

	jobs := make(chan string, len(allUsers))
	results := make(chan error, len(allUsers))

	for w := 1; w <= maxConcurrentWorkers; w++ {
		go worker(w, jobs, results)
	}
	for _, user := range allUsers {
		jobs <- user
	}
	close(jobs)

	for a := 1; a <= len(allUsers); a++ {
		err = <-results
		log.Println(err)
	}
	return nil
}

func DoArchiveForYesterdayWrapper() {
	err := DoArchiveForYesterday()
	if err != nil {
		log.Println(err)
	}
}
