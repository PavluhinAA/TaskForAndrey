package main

import (
	"encoding/json"
	"log"
	"os"
)

const dbFile = "data.json"

// InMemoryDB - простая in-memory база данных с сохранением в JSON-файл
type book struct {
	Title    string
	Author   string
	reserved bool
}
type InMemoryDB struct {
	data []book
}

// NewInMemoryDB - конструктор для InMemoryDB
func NewInMemoryDB() *InMemoryDB {
	db := &InMemoryDB{}
	db.data = make([]book, 0)
	// mutex инициализируется автоматически нулевым значением, что подходит для sync.RWMutex
	return db
}

// Get - получает значение по ключу
func (db *InMemoryDB) Get(key int) book {
	value := db.data[key]
	return value
}

// Set - устанавливает значение по ключу
func (db *InMemoryDB) Set(title, author string) {
	db.data = append(db.data, book{Title: title, Author: author, reserved: true})
}

// LoadFromFile - загружает данные из JSON-файла
func (db *InMemoryDB) LoadFromFile(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не существует - ничего страшного, просто создадим пустую базу данных
			log.Println("Файл базы данных не найден, будет создана новая база данных")
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db.data)
	if err != nil {
		return err
	}

	return nil
}

// SaveToFile - сохраняет данные в JSON-файл
func (db *InMemoryDB) SaveToFile(filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(db.data)
	if err != nil {
		return err
	}
	return nil
}
