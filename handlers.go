package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strings"
)

var db *InMemoryDB

func Start(w http.ResponseWriter, _ *http.Request) {
	errorFpr(fmt.Fprintf(w, "Library\n To enter the website, use the search bar to the specified format"))
}

func Books(w http.ResponseWriter, _ *http.Request) {
	errorFpr(fmt.Fprintf(w, "These books are now available\n"))
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Reserved == false {
			stringBook, _ := json.Marshal(db.data[i])
			formatStr := formatOutput(string(stringBook))
			errorFpr(fmt.Fprintf(w, "%s\n", formatStr))
		}
	}
}

func BooksAll(w http.ResponseWriter, _ *http.Request) {
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

func BooksNew(_ http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	author := chi.URLParam(r, "author")

	if title != "" && author != "" {

		for i := 0; i < len(db.data); i++ {

			if db.data[i].Title == title {
				return
			}
		}
		db.Set(title, author)
		err := db.SaveToFile(dbFile)
		if err != nil {
			log.Println("Error when saving to a file:", err)
			return
		}
		log.Println("Saved to a file:", dbFile)
	}
}

func Reserved(_ http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Title == title {
			if db.data[i].Reserved == true {
				db.data[i].Reserved = false
				err := db.SaveToFile(dbFile)
				if err != nil {
					log.Println("Error when saving to a file:", err)
					return
				}
				return
			}
			if db.data[i].Reserved == false {
				db.data[i].Reserved = true
				err := db.SaveToFile(dbFile)
				if err != nil {
					log.Println("Error when saving to a file:", err)
					return
				}
				return
			}
		}
	}
}

func BooksDel(_ http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	for i := 0; i < len(db.data); i++ {
		if db.data[i].Title == title {
			db.data = append(db.data[:i], db.data[i+1:]...)
			err := db.SaveToFile(dbFile)
			if err != nil {
				log.Println("Error when saving to a file:", err)
			}
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
