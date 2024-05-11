package main

import (
	"bookbuzz/app/controller"
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
	r.GET("/login", controller.LoginPage)
	r.POST("/login", model.AuthHandler) // Изменили метод на POST для обработки отправки данных формы
	r.HandleMethodNotAllowed = false    // Отключаем автоматическую обработку методов, чтобы использовать кастомную обработку OPTIONS
	r.OPTIONS("/login", handleOptions)  // Добавляем обработчик OPTIONS для маршрута /login
	fs := http.FileServer(http.Dir("public/html"))
	r.Handler("GET", "/public/*filepath", http.StripPrefix("/public/", fs))
}
func handleOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Allow", "POST, GET, OPTIONS") // Устанавливаем заголовок Allow, указывающий разрешенные методы
	w.WriteHeader(http.StatusNoContent)
}
func main() {
	// Обработчик для статических файлов
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Обработчик для маршрута /login
	http.HandleFunc("/login", loginHandler)

	// Запуск сервера на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение страницы login.html
		http.ServeFile(w, r, "public/html/login.html")
	} else if r.Method == "POST" {
		// Обработка отправленных данных
		// Извлечение данных из формы
		username := r.FormValue("login")
		password := r.FormValue("password")

		// Дальнейшая обработка данных
		// ...

		// Отправка ответа клиенту
		fmt.Fprintf(w, "Данные получены: username=%s, password=%s", username, password)
	}
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
