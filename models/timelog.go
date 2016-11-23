package models

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/tonyalaribe/monitor-server/constants"
)

//UserTimeLog basically represents a user. There are two password fields.
type UserTimeLog struct {
	UserID string `json:"username"`
}

type TimeLog struct {
	Day    time.Time
	Login  time.Time
	Logout time.Time
}

//LoginNow creates a key and value with the key representing the current day, and the value representing the current time, which would serve as login time
func (user UserTimeLog) LoginNow(db *bolt.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.LOGIN_BUCKET))
	log.Printf("%+v", user)
	log.Println("thats the user sttruct")
	//Generate a bcrypt encoded hash of the password

	bkt, err := b.CreateBucketIfNotExists([]byte(user.UserID))
	if err != nil {
		return err
	}
	now := time.Now()
	currentTime := now.Format(time.RFC3339)
	today := []byte(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format(time.RFC3339))

	err = bkt.Put([]byte(today), []byte(currentTime))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

//LogoutNow creates a key and value with the key representing the current day, and the value representing the current time, which would serve as logout time. It should be called after each db reated activity like a new screenshot, to signify user's last active time.
func (user UserTimeLog) LogoutNow(db *bolt.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.LOGOUT_BUCKET))
	log.Printf("%+v", user)
	//Generate a bcrypt encoded hash of the password

	bkt, err := b.CreateBucketIfNotExists([]byte(user.UserID))
	if err != nil {
		return err
	}
	now := time.Now()
	currentTime := now.Format(time.RFC3339)
	today := []byte(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format(time.RFC3339))

	err = bkt.Put([]byte(today), []byte(currentTime))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

//GetTimeline returns an array of the various days and when the user logged in and out.
func (user UserTimeLog) GetTimeline(db *bolt.DB) ([]TimeLog, error) {
	timelog := []TimeLog{}

	tx, err := db.Begin(true)
	if err != nil {
		return timelog, err
	}
	defer tx.Rollback()

	login := tx.Bucket([]byte(constants.LOGIN_BUCKET))

	logout := tx.Bucket([]byte(constants.LOGOUT_BUCKET))
	log.Printf("%+v", user)
	log.Println("thats the user sttruct")

	userlogin, err := login.CreateBucketIfNotExists([]byte(user.UserID))
	if err != nil {
		return timelog, err
	}
	userlogout, err := logout.CreateBucketIfNotExists([]byte(user.UserID))
	if err != nil {
		return timelog, err
	}

	loginCursor := userlogin.Cursor()

	i := 0
	for k, v := loginCursor.First(); k != nil && i < 31; k, v = loginCursor.Next() {
		i++
		//fmt.Printf("key=%s, value=%s\n", k, v)
		day, err := time.Parse(time.RFC3339, string(k))
		if err != nil {
			log.Println(err)
		}
		logoutTime := userlogout.Get(k)
		lo, err := time.Parse(time.RFC3339, string(logoutTime))
		if err != nil {
			log.Println(err)
		}
		li, err := time.Parse(time.RFC3339, string(v))
		if err != nil {
			log.Println(err)
		}
		timelog = append(timelog, TimeLog{day, li, lo})
	}
	return timelog, nil
}
