package bolt

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"log"

	"github.com/boltdb/bolt"
)

var database *bolt.DB

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
	database = db
	if err != nil {
		log.Fatal(err)
	}
	DBCreate(boltDB{db: db})
	return boltDB{db: db}
}

func GetboltDB() boltDB {

	if database != nil {
		return boltDB{db: database}
	}

	return DBopen("base.db")

}

func (b boltDB) DBclose(db *bolt.DB) {
	db.Close()
}

func DBCreate(b boltDB) {

	var buckets []string = []string{"Students", "Languages"}

	for _, name := range buckets {
		b.deleteBucket(name)
		b.DBcreateBucket(name)
	}

	b.insertDefaultDataLanguages()
	b.insertDefaultDataStudents()

}

func (b boltDB) insertDefaultDataLanguages() {

	var languages []entities.Language = []entities.Language{
		entities.NewLanguage("FR", "Français"),
		entities.NewLanguage("EN", "Anglais"),
		entities.NewLanguage("IT", "Italien"),
	}

	for _, language := range languages {

		res, _ := json.Marshal(language)

		b.DBput("Languages", language.Code, string(res))
	}

}

func (b boltDB) insertDefaultDataStudents() {
	var students []entities.Student = []entities.Student{
		entities.NewStudent(2, "Marin", "Marais", 18, "FR"),
		entities.NewStudent(3, "Arthur", "Dubeau", 20, "FR"),
		entities.NewStudent(4, "Sebastiano", "Sagese", 23, "IT"),
		entities.NewStudent(1, "Matilda", "Barton", 19, "EN"),
		entities.NewStudent(1, "Aaron", "Hardy", 20, "EN"),
	}

	for _, student := range students {

		res, _ := json.Marshal(student)

		b.DBput("Students", fmt.Sprintf("%d", student.Id), string(res))
	}

}

func (b boltDB) DBpath() string {

	return b.db.Path()
}

func (b boltDB) DBput(bucketName string, key string, value string) {

	err := b.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			panic("Bucket : " + bucketName + "existe pas")
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

	var resultat []string

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			value := fmt.Sprintf("%s", v)
			resultat = append(resultat, value)

		}

		return nil
	})

	if err != nil {
		log.Fatal(err)

	}

	return resultat
}

func (b boltDB) DBdelete(bucketName string, key string) error {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {

			panic("Bucket : " + bucketName + " non trouvé.")

		}
		bucket.Delete([]byte(key))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (b boltDB) DBcreateBucket(bucketName string) {

	err := b.db.Update(func(tx *bolt.Tx) error {
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

func (b boltDB) deleteBucket(bucketName string) {

	err := b.db.Update(func(tx *bolt.Tx) error {

		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Le bucket %s ne peut être surpprimé car il n'éxiste pas.\n", bucketName)
	}
}
