package main

import (
	"encoding/json"
	"log"
	"os"
)

const dbFile = "data.json"

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Reserved bool   `json:"reserved"`
}

type InMemoryDB struct {
	data []Book
}

func NewInMemoryDB() *InMemoryDB {
	db := &InMemoryDB{}
	db.data = make([]Book, 0)
	return db
}

func (db *InMemoryDB) Get(index int) Book {
	value := db.data[index]
	return value
}

func (db *InMemoryDB) Set(title, author string) {
	db.data = append(db.data, Book{Title: title, Author: author, Reserved: false})
}

func (db *InMemoryDB) LoadFromFile(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("The database file is not found, a new database will be created.")
			return nil
		}
		return err
	}
	defer func() {
		if errClose := file.Close(); errClose != nil {
			log.Println("Error when closing a file:", errClose)
		}
	}()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db.data)
	if err != nil {
		return err
	}
	return nil
}

func (db *InMemoryDB) SaveToFile(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		if errClose := file.Close(); errClose != nil {
			log.Println("Error when closing a file:", errClose)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(db.data)
	if err != nil {
		return err
	}
	return nil
}
