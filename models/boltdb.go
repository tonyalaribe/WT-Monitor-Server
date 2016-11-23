package models

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/tonyalaribe/monitor-server/constants"
)

// Client represents a client to the underlying BoltDB data store.
type Client struct {
	Path string   //Filename to the bolt db file
	DB   *bolt.DB //pointer to a bolt instance
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	// Open database file.
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println(c.Path)
		log.Println("jjhjkl")
		log.Println(err)
		return err
	}
	c.DB = db

	// Start writable transaction.
	tx, err := c.DB.Begin(true)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	// Initialize top-level buckets.
	if _, err = tx.CreateBucketIfNotExists([]byte(constants.USER_BUCKET)); err != nil {
		log.Println(err)
		return err
	}

	if _, err = tx.CreateBucketIfNotExists([]byte(constants.USERLIST)); err != nil {
		log.Println(err)
		return err
	}

	if _, err = tx.CreateBucketIfNotExists([]byte(constants.LOGIN_BUCKET)); err != nil {
		log.Println(err)
		return err
	}

	if _, err = tx.CreateBucketIfNotExists([]byte(constants.LOGOUT_BUCKET)); err != nil {
		log.Println(err)
		return err
	}

	if _, err = tx.CreateBucketIfNotExists([]byte(constants.SCREENSHOTS_BUCKET)); err != nil {
		log.Println(err)
		return err
	}

	if _, err = tx.CreateBucketIfNotExists([]byte(constants.ARCHIVE_BUCKET)); err != nil {
		log.Println(err)
		return err
	}

	// Save transaction to disk.
	return tx.Commit()
}
