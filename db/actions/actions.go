package actions

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

const (
	ADD  = "add"
	DO   = "do"
	LIST = "list"
)

func Actions(action string, arg ...string) {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = initBucketIfNotExists(db, "todo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	switch action {
	case ADD:
		addTODO(db, arg[0])
	case DO:
		doneTODO(db, arg[0])
	case LIST:
		listTODO(db)
	}
}

func doneTODO(db *bolt.DB, doneNo string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		doneNumber, _ := strconv.Atoi(doneNo)

		bucket := tx.Bucket([]byte("todo"))
		c := bucket.Cursor()

		i := 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if i == doneNumber {
				fmt.Printf("You have completed the \"%s\" task.\n", v)
				return bucket.Delete(k)
			}
			i += 1
		}

		return errors.New("No such serial number.")
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func listTODO(db *bolt.DB) error {
	todoList := make([]string, 0)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("todo"))
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			// fmt.Printf("%d. %s\n", i, v)
			todoList = append(todoList, string(v))
		}

		return nil
	})

	fmt.Println("You have the following tasks:")
	for i, todo := range todoList {
		fmt.Printf("%d. %s\n", i+1, todo)
	}

	if err != nil {
		return err
	}

	return nil
}

func initBucketIfNotExists(db *bolt.DB, bucketName string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
}

// todo table:
//         key: auto-increment integer
//         value: todo task
//
// completed table:
//         key: time.Now()
//         value: completed task

// add todo task to todo bucket
func addTODO(db *bolt.DB, task string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return err
		}
		id, _ := bucket.NextSequence()
		fmt.Printf("Added \"%s\" to your task list.\n", task)
		return bucket.Put(iotb(int(id)), []byte(task))
	})

	if err != nil {
		return err
	}

	return nil
}

// itob returns an 8-byte big endian representation of v.
func iotb(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
