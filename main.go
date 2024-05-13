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
	//r.ServeFiles("/public/*filepath", http.Dir("public"))
	//r.GET("/", controller.StartPage)
	//r.GET("/home", controller.HomePage) // Добавленный маршрут для домашней страницы
	//r.GET("/login", controller.LoginPage)
	//r.POST("/login", model.AuthHandler) // Изменили метод на POST для обработки отправки данных формы
	//r.HandleMethodNotAllowed = false   // Отключаем автоматическую обработку методов, чтобы использовать кастомную обработку OPTIONS
	//r.OPTIONS("/login", handleOptions) // Добавляем обработчик OPTIONS для маршрута /login
	fs := http.FileServer(http.Dir("public/html"))
	r.Handler("GET", "/public/*filepath", http.StripPrefix("/public/", fs))
	//r.GET("/registration_page", model.RegistrationPage) // Добавляем маршрут для страницы регистрации
	//r.POST("/registration_page", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//	model.RegistrationPage(w, r)
	//})

}

func main() {
	r := httprouter.New()
	routes(r)
	// Обработчик для статических файлов
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Обработчик для маршрута /login
	http.HandleFunc("/login", model.LoginHandler)

	http.HandleFunc("/registration_page", model.RegisterHandler)
	http.HandleFunc("/regform", model.RegisterHandler)

	http.HandleFunc("/home", model.HomeHandler)
	http.HandleFunc("/order_page", model.OrderHandler)
	//	http.HandleFunc("/new_order", model.NewOrderHandler)

	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
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
*/
