package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handlerBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "это список книг доступных в данный момент") //
	for i := 0; i < len(db.data); i++ {
		if db.data[i].reserved == true {
			stringBook, _ := json.Marshal(db.data[i])
			fmt.Fprintf(w, string(stringBook))
		}
	}
}

func handlerReservedBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "данные книги зарезервированы")
	for i := 0; i < len(db.data); i++ {
		if db.data[i].reserved == false {
			stringBook, _ := json.Marshal(db.data[i])
			fmt.Fprintf(w, string(stringBook))
		}
	}
}

func handlerBooksNew(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")

	// Проверка, что оба поля заполнены
	if title != "" && author != "" {
		db.Set(title, author)        // Добавляем книгу в базу данных
		err := db.SaveToFile(dbFile) // Сохраняем базу данных в файл

		if err != nil {
			log.Println("Ошибка при сохранении в файл:", err)
			http.Error(w, "Ошибка при сохранении в файл", http.StatusInternalServerError) // Отправляем ошибку клиенту
			return
		}
		fmt.Fprintf(w, "Книга успешно внесена в библиотеку") // Отправляем сообщение об успехе
	} else {
		fmt.Fprintf(w, "Укажите название и автора книги в поисковую строку в формате:\n /?title=ваша книга&author=ваш автор")
	}
}

var db *InMemoryDB // Глобальная переменная для хранения экземпляра базы данных

func main() {
	var wg sync.WaitGroup
	go shutdown(&wg)
	db = NewInMemoryDB() // Инициализируем базу данных

	// Загружаем данные из файла при запуске
	err := db.LoadFromFile(dbFile) // Загружаем данные из JSON-файла
	if err != nil {                // Проверяем, произошла ли ошибка при загрузке из файла
		log.Println("Ошибка при загрузке из файла:", err) // Записываем ошибку в лог
		// Не завершаем программу, а продолжаем работу с пустой базой данных
	}

	http.HandleFunc("/books/all", handlerReservedBooks)
	http.HandleFunc("/books", handlerBooks)
	http.HandleFunc("/books/new", handlerBooksNew)
	fmt.Println("Запускаем сервер на порту 8080") // Выводим сообщение в консоль
	err = http.ListenAndServe(":8080", nil)       // Запускаем веб-сервер на порту 8080
	if err != nil {                               // Проверяем, произошла ли ошибка при запуске сервера
		fmt.Println("Ошибка запуска сервера:", err) // Выводим сообщение об ошибке в консоль
	}
	wg.Wait()
}

func shutdown(wg *sync.WaitGroup) {

	wg.Add(1)
	defer wg.Done()

	var stopSignal = make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case <-stopSignal:
		db.SaveToFile(dbFile)
		fmt.Println("the process is completed because the completion signal has been received")
		return
	}
}
