package main

import (
	"bookbuzz/app/controller"
	"bookbuzz/app/model"
	_ "bookbuzz/app/model"
	"bookbuzz/app/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/home", controller.HomePage)  // Добавленный маршрут для домашней страницы
	r.POST("/login", model.LoginHandler) // Изменили метод на POST для обработки отправки данных формы
}
func main() {
	//инициализируем подключение к базе данных
	err := server.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	//создаем и запускаем в работу роутер для обслуживания запросов
	r := httprouter.New()
	routes(r)
	//прикрепляемся хосту и порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err = http.ListenAndServe(":8080", r) // Замените "localhost:4444" на ":8080" или другой желаемый порт
	if err != nil {
		log.Fatal(err)
	}
}
