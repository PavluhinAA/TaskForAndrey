package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "это список книг доступных в данный момент")
	for i := 0; i < len
}
func handlerAllBooks(w http.ResponseWriter, r *http.Request) {

}
func handlerAuthors(w http.ResponseWriter, r *http.Request) {

}
func handlerReserved(w http.ResponseWriter, r *http.Request) {

}
func handlerBooksNew(w http.ResponseWriter, r *http.Request) {

}

var db *InMemoryDB // Глобальная переменная для хранения экземпляра базы данных

func main() {
	db = NewInMemoryDB() // Инициализируем базу данных

	// Загружаем данные из файла при запуске
	err := db.LoadFromFile(dbFile) // Загружаем данные из JSON-файла
	if err != nil {                // Проверяем, произошла ли ошибка при загрузке из файла
		log.Println("Ошибка при загрузке из файла:", err) // Записываем ошибку в лог
		// Не завершаем программу, а продолжаем работу с пустой базой данных
	}
	http.HandleFunc("/books/all", handlerAllBooks)
	http.HandleFunc("/books", handlerBooks) // Регистрируем функцию-обработчик для корневого пути ("/")
	http.HandleFunc("/author", handlerAuthors)
	http.HandleFunc("/reserved", handlerReserved)
	http.HandleFunc("/books/new", handlerBooksNew)
	fmt.Println("Запускаем сервер на порту 8080") // Выводим сообщение в консоль
	err = http.ListenAndServe(":8080", nil)       // Запускаем веб-сервер на порту 8080
	if err != nil {                               // Проверяем, произошла ли ошибка при запуске сервера
		fmt.Println("Ошибка запуска сервера:", err) // Выводим сообщение об ошибке в консоль
	}
}
