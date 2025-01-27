package actions

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

const (
	ADD       = "add"
	DO        = "do"
	LIST      = "list"
	REMOVE    = "rm"
	COMPLETED = "completed"
)

func Actions(action string, arg ...string) {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = initBucketIfNotExists(db, "todo")
	err = initBucketIfNotExists(db, "complete")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	switch action {
	case ADD:
		addTask, _ := addTODO(db, arg[0])
		fmt.Printf("Added \"%s\" to your task list.\n", addTask)
	case DO:
		err := doTask(db, arg[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	case LIST:
		tasks, _ := listTODO(db)
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	case REMOVE:
		rmTask, err := removeTask(db, arg[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("You have deleted the \"%s\" task.\n", rmTask)
	case COMPLETED:
		completedTask, _ := listCompleted(db)
		fmt.Println("You have finished the following tasks today:")
		for _, complete := range completedTask {
			fmt.Printf("- %s\n", complete)
		}
	}
}

// list today completed task
func listCompleted(db *bolt.DB) ([]string, error) {
	completedTask := make([]string, 0)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("complete"))
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			now := time.Now()
			if isSameDate(string(k), now.Format(time.DateTime)) {
				completedTask = append(completedTask, string(v))
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return completedTask, nil
}

// remove removeNo. task from bucket
func removeTask(db *bolt.DB, removeNo string) (string, error) {
	var removedTask string
	err := db.Update(func(tx *bolt.Tx) error {
		removeNumber, _ := strconv.Atoi(removeNo)

		bucket := tx.Bucket([]byte("todo"))
		c := bucket.Cursor()

		i := 1
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if i == removeNumber {
				removedTask = string(v)
				return bucket.Delete(k)
			}
			i += 1
		}

		return errors.New(fmt.Sprintf("No such serial number %s", removeNo))
	})

	if err != nil {
		return "", err
	}

	return removedTask, nil
}

// mark the finished task to complete bucket
func doTask(db *bolt.DB, doneNo string) error {
	// remove task from todo bucket
	finishedTask, err := removeTask(db, doneNo)
	if err != nil {
		return err
	}

	// add finishedTask to complete bucket
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("complete"))
		if err != nil {
			return err
		}
		t, task := time.Now().Format(time.DateTime), finishedTask
		fmt.Printf("At %s, you finished \"%s\"\n", t, task)
		return bucket.Put([]byte(t), []byte(task))
	})
}

// show todo tasks list
func listTODO(db *bolt.DB) ([]string, error) {
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

	if err != nil {
		return nil, err
	}

	return todoList, nil
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

// add todo task to todo bucket
func addTODO(db *bolt.DB, task string) (string, error) {
	var addedTask string
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return err
		}
		id, _ := bucket.NextSequence()
		addedTask = string(task)
		return bucket.Put(iotb(int(id)), []byte(task))
	})

	if err != nil {
		return "", err
	}

	return addedTask, nil
}
