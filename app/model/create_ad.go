package model

import (
	"database/sql"
	"log"
	"net/http"
)

func NewAdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение страницы регистрации
		http.ServeFile(w, r, "public/html/new_ad.html")
	} else if r.Method == "POST" {
		// Проверка аутентификации пользователя
		session, err := store.Get(r, "session-name")
		if err != nil {
			log.Fatal(err)
		}
		userID, ok := session.Values["userID"].(int)
		if !ok {
			// Пользователь не аутентифицирован, перенаправление на страницу входа
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		// Получение данных из формы
		title := r.FormValue("title_ad")
		description := r.FormValue("description")
		// Подключение к базе данных
		db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		// Выполнение SQL-запроса для сохранения заказа в базе данных
		query := "INSERT INTO ad2 (title_ad, text_ad, fk_user) VALUES (?, ?, ?)"
		_, err = db.Exec(query, title, description, userID)
		if err != nil {
			log.Fatal(err)
		}
		// Перенаправление пользователя на другую страницу после успешного сохранения заказа
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}
