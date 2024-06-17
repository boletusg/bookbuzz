package model

import (
	"database/sql"
	"encoding/json"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/sessions"
	_ "github.com/gorilla/sessions"
	"log"
	"net/http"
)

type User struct {
	Id           int    `json:"id_user" db:"id_user"`
	Login        string `json:"login_user" db:"login_user"`
	UserPassword string `json:"password_user" db:"password_user"`
	Name         string `json:"name_user" db:"name_user"`
	Nickname     string `json:"nickname" db:"nickname"`
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var store = sessions.NewCookieStore([]byte("my-secret-key"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение страницы login.html
		http.ServeFile(w, r, "public/html/login.html")
	} else if r.Method == "POST" {
		// Получение логина и пароля из формы входа
		login := r.FormValue("login")
		password := r.FormValue("password")

		// Подключение к базе данных
		db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
		if err != nil {
			log.Fatal(err)
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
			}
		}(db)

		// Подготовка SQL-запроса для проверки логина и пароля в таблице users2
		query := "SELECT COUNT(*) FROM users2 WHERE login_user = ? AND password_user = ?"
		var count int
		err = db.QueryRow(query, login, password).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		var response Response

		if count > 0 {
			// Логин и пароль найдены
			response.Success = true
			response.Message = "Успешная аутентификация!"
			var userID int
			query := "SELECT id_user FROM users2 WHERE login_user = ? AND password_user = ?"
			err = db.QueryRow(query, login, password).Scan(&userID)
			if err != nil {
				log.Fatal(err)
			}
			// Создание новой сессии
			session, err := store.New(r, "session-name")
			if err != nil {
				log.Fatal(err)
			}
			// Сохранение идентификатора пользователя в сессии
			session.Values["userID"] = userID
			err = session.Save(r, w)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// Логин и пароль не найдены в базе данных
			response.Message = "Неверный логин или пароль!"
		}
		// Преобразование данных в формат JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		// Установка заголовков для ответа
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Отправка ответа в формате JSON
		_, _ = w.Write(jsonResponse)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение страницы регистрации
		http.ServeFile(w, r, "public/html/registration_page.html")
	} else if r.Method == "POST" {
		var response Response

		// Получение данных пользователя из формы регистрации
		nickname := r.FormValue("nickname")
		name := r.FormValue("name")
		phone := r.FormValue("phonenumber")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")

		// Подключение к базе данных
		db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
		if err != nil {
			http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
			}
		}(db)
		// Проверка наличия пользователя с заданным логином
		exists, err := checkUserExists(db, nickname)
		// Проверка паролей на соответствие
		if password != password2 {
			response.Message = "Пароли не совпадают"
		} else if err != nil {
			http.Error(w, "Ошибка при выполнении SQL-запроса", http.StatusInternalServerError)
			log.Fatal(err)
			return
		} else if exists {
			response.Message = "Пользователь с таким логином уже существует"
		} else {

			// Подготовка SQL-запроса для вставки нового пользователя в таблицу users2
			query := "INSERT INTO users2 (login_user, name_user, phone_user, password_user) VALUES (?, ?, ?, ?)"
			_, err = db.Exec(query, nickname, name, phone, password)
			if err != nil {
				http.Error(w, "Ошибка при выполнении SQL-запроса", http.StatusInternalServerError)
				log.Fatal(err)
				return
			}
			// Отправка ответа об успешной регистрации
			response.Success = true
			response.Message = "Регистрация прошла успешно!"
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Отправка ответа в формате JSON
		_, _ = w.Write(jsonResponse)
	}
}
func checkUserExists(db *sql.DB, login string) (bool, error) {
	query := "SELECT COUNT(*) FROM users2 WHERE login_user = ?"
	var count int
	err := db.QueryRow(query, login).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
