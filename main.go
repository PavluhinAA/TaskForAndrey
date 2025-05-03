package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := chi.NewRouter()

	go shutdown()

	db = NewInMemoryDB()
	err := db.LoadFromFile(dbFile)

	if err != nil {
		log.Println("Error when downloading from a file:", err)
	}

	r.Use(middleware.Logger)

	r.Get("/", Start)

	r.Route("/books", func(r chi.Router) {

		r.Get("/", Books)
		r.Get("/all", BooksAll)
		r.Method(http.MethodGet, "/new/{title}/{author}", http.HandlerFunc(BooksNew))
		r.Method(http.MethodPost, "/new/{title}/{author}", http.HandlerFunc(BooksNew))
		r.Patch("/reserved/{title}", Reserved)
		r.Delete("/delete/{title}", BooksDel)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func shutdown() {

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
