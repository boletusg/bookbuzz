package main

import (
	"bookbuzz/app/model"
	_ "bookbuzz/app/model"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func routes(r *httprouter.Router) {
	fs := http.FileServer(http.Dir("public/html"))
	r.Handler("GET", "/public/*filepath", http.StripPrefix("/public/", fs))
}

func main() {
	r := httprouter.New()
	routes(r)
	// Обработчик для статических файлов
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Обработчики для маршрутов страниц
	http.HandleFunc("/login", model.LoginHandler)
	http.HandleFunc("/registration_page", model.RegisterHandler)
	http.HandleFunc("/regform", model.RegisterHandler)
	http.HandleFunc("/home", model.HomeHandler)
	http.HandleFunc("/order_page", model.OrderHandler)
	http.HandleFunc("/new_order", model.NewOrderHandler)
	http.HandleFunc("/account_page", model.AccountHandler)
	http.HandleFunc("/account_edit", model.EditHandler)
	http.HandleFunc("/new_ad", model.NewAdHandler)
	http.HandleFunc("/respond", model.RespondHandler)
	http.HandleFunc("/my_orders", model.RespondsHandler)
	http.HandleFunc("/messenger_page", model.MessengerHandler)
	http.HandleFunc("/chat", model.HandleWebSocket)

	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
