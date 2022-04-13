package bolt

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type boltDB struct {
	db *bolt.DB
}

/*func NewBoltDB() boltDB {
	var bo boltDB
	bo.DBopen("base.db")
	return bo
}*/

func DBopen(filename string) boltDB {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return boltDB{db: db}
}

func (b boltDB) DBclose(db *bolt.DB) {
	db.Close()
}

func (b boltDB) DBpath() string {

	return b.db.Path()
}

func (b boltDB) DBput(bucketName string, key string, value string) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			log.Fatal()
		}

		bucket.Put([]byte(key), []byte(value))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}

func (b boltDB) DBget(bucketName string, key string) string {
	var value string
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			log.Fatal()
		}
		value = string(bucket.Get([]byte(key)))

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (b boltDB) DBgetAll(bucketName string) []string {

	var res []string
	err := b.db.View(func(tx *bolt.Tx) error {
		cursor := tx.Cursor()
		bucket := tx.Bucket([]byte(bucketName))
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Println("%s", v)
			res = append(res, fmt.Sprintf("%s", v))
		}

		if bucket == nil {
			log.Fatal(fmt.Sprintf("%s = nil", bucketName))
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (b boltDB) DBdelete(bucketName string, key string) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			log.Fatal()
		}
		bucket.Delete([]byte(key))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func DBcreateBucket(db *bolt.DB, bucketName string) {

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
