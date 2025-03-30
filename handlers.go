package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var db *InMemoryDB

func handlerStart(w http.ResponseWriter, _ *http.Request) {
	errorFpr(fmt.Fprintf(w, "Library\n To enter the website, use the search bar to the specified format"))
}

func handlerBooks(w http.ResponseWriter, _ *http.Request) {
	errorFpr(fmt.Fprintf(w, "These books are now available\n"))
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Reserved == false {
			stringBook, _ := json.Marshal(db.data[i])
			formatStr := formatOutput(string(stringBook))
			errorFpr(fmt.Fprintf(w, "%s\n", formatStr))
		}
	}
}

func handlerBooksAll(w http.ResponseWriter, _ *http.Request) {
	errorFpr(fmt.Fprintf(w, "These books are reserved\n"))
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Reserved == true {
			stringBook, _ := json.Marshal(db.data[i])
			formatStr := formatOutput(string(stringBook))
			errorFpr(fmt.Fprintf(w, "%s\n", formatStr))
		}
	}
	errorFpr(fmt.Fprintf(w, "These books are available\n"))
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Reserved == false {
			stringBook, _ := json.Marshal(db.data[i])
			formatStr := formatOutput(string(stringBook))
			errorFpr(fmt.Fprintf(w, "%s\n", formatStr))
		}
	}
}

func handlerBooksNew(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")
	errorFpr(fmt.Fprintf(w, "Enter the title and author of the book\nFormat:\"http://localhost:8080/books/new?title=......&author=......\"\n"))
	if title != "" && author != "" {
		for i := 0; i < len(db.data); i++ {
			if db.data[i].Title == title {
				errorFpr(fmt.Fprintf(w, "This book is already in the library."))
				return
			}
		}
		db.Set(title, author)
		err := db.SaveToFile(dbFile)

		if err != nil {
			log.Println("Error when saving to a file:", err)
			return
		}
		errorFpr(fmt.Fprintf(w, "The book has been successfully added to the library"))
	}
}

func handlerReserved(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	reserved := r.URL.Query().Get("reserved")
	errorFpr(fmt.Fprintf(w, "specify the book and its condition\nFormat:\"http://localhost:8080/books/reserved?title=......&reserved=true/false\"\n"))
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Title == title {
			if reserved == "true" {
				db.data[i].Reserved = true
				err := db.SaveToFile(dbFile)
				if err != nil {
					log.Println("Error when saving to a file:", err)
					return
				}
				errorFpr(fmt.Fprintf(w, "The book is reserved"))
				return
			}
			if reserved == "false" {
				db.data[i].Reserved = false
				err := db.SaveToFile(dbFile)
				if err != nil {
					log.Println("Error when saving to a file:", err)
					return
				}
				errorFpr(fmt.Fprintf(w, "the book is no longer preserved"))
				return

			}
		}
	}
}

func handlerBooksDel(w http.ResponseWriter, r *http.Request) {
	errorFpr(fmt.Fprintf(w, "Enter the name of the book you want to delete from the library\nFormat:\"http://localhost:8080/books/delete?title=......\"\n"))
	title := r.URL.Query().Get("title")
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Title == title {
			db.data = append(db.data[:i], db.data[i+1:]...)
			err := db.SaveToFile(dbFile)
			if err != nil {
				log.Println("Error when saving to a file:", err)
			}
			errorFpr(fmt.Fprintf(w, "The book was deleted"))
		}
	}
}

func errorFpr(_ int, err error) {
	if err != nil {
		log.Println("error in displaying information to the user", err)
	}
}

func formatOutput(str string) string {
	stringBook := strings.Trim(str, "{")
	stringBook = strings.Trim(stringBook, "}")
	stringBook = strings.ReplaceAll(stringBook, "\"", "")
	stringBook = strings.ReplaceAll(stringBook, ",reserved:false", "")
	stringBook = strings.ReplaceAll(stringBook, ",reserved:true", "")
	return stringBook
}
