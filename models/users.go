package models

import (
	"github.com/tonyalaribe/monitor-server/constants"

	"github.com/boltdb/bolt"
)

//User basically represents a user. There are two password fields.
type User struct {
	UserName string `json:"username"`
}

//Create adds a new user to the database
func (user *User) Create(db *bolt.DB) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.SCREENSHOTS_BUCKET))
	_, err = b.CreateBucketIfNotExists([]byte(user.UserName))
	if err != nil {
		return err
	}

	b = tx.Bucket([]byte(constants.LOGIN_BUCKET))
	_, err = b.CreateBucketIfNotExists([]byte(user.UserName))
	if err != nil {
		return err
	}

	b = tx.Bucket([]byte(constants.LOGOUT_BUCKET))
	_, err = b.CreateBucketIfNotExists([]byte(user.UserName))
	if err != nil {
		return err
	}

	b = tx.Bucket([]byte(constants.ARCHIVE_BUCKET))
	_, err = b.CreateBucketIfNotExists([]byte(user.UserName))
	if err != nil {
		return err
	}

	userlist := tx.Bucket([]byte(constants.USERLIST))
	err = userlist.Put([]byte(user.UserName), nil)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

//GetAll retrieves a user from the databse based on its key
func (user User) GetAll(db *bolt.DB) ([]string, error) {
	//log.Println("get all users")
	users := []string{}
	tx, err := db.Begin(false)
	if err != nil {
		return users, err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(constants.USERLIST))

	c := b.Cursor()

	for k, _ := c.First(); k != nil; k, _ = c.Next() {
		//fmt.Printf("key=%s, value=%s\n", k, v)
		users = append(users, string(k))
	}

	//log.Println(users)
	return users, nil
}
