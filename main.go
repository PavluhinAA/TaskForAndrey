package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	go shutdown(&wg)

	db = NewInMemoryDB()

	err := db.LoadFromFile(dbFile)
	if err != nil {
		log.Println("Error when downloading from a file:", err)
	}

	http.HandleFunc("/", handlerStart)
	http.HandleFunc("/books/delete", handlerBooksDel)
	http.HandleFunc("/books/all", handlerBooksAll)
	http.HandleFunc("/books", handlerBooks)
	http.HandleFunc("/books/new", handlerBooksNew)
	http.HandleFunc("/books/reserved", handlerReserved)

	log.Println("Starting the server")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server startup error:", err)
	}

	wg.Wait()
}

func shutdown(wg *sync.WaitGroup) {

	wg.Add(1)
	defer wg.Done()

	var stopSignal = make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopSignal
	log.Println("saving data")
	err := db.SaveToFile(dbFile)
	if err != nil {
		log.Println(err)
	}
	log.Println("saving completed, program termination")
	os.Exit(0)
}
