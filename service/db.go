package service

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

// OpenDB opens the BoltDB database and returns the instance.
func OpenDB() (*bolt.DB, error) {
	var err error
	db, err = bolt.Open("../db/my.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CloseDB closes the BoltDB database.
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func GetValueFromDB(db *bolt.DB, key string) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("myBucket"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}
		value = bucket.Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}
