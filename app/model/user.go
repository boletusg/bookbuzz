package model

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb" // Драйвер MSSQL для Go
)

type User struct {
	Id           int    `json:"id_user" db:"id_user"`
	Login        string `json:"login_user" db:"login_user"`
	UserPassword string `json:"password_user" db:"password_user"`
	Name         string `json:"name_user" db:"name_user"`
	Nickname     string `json:"nickname" db:"nickname"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Установка соединения с базой данных MSSQL
	db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Получение введенных данных из формы
	login := r.FormValue("login")
	password := r.FormValue("password")

	// Проверка введенных данных с данными в базе данных
	query := "SELECT COUNT(*) FROM users2 WHERE login_user = ? AND password_user = ?"
	var count int
	err = db.QueryRow(query, login, password).Scan(&count)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		// Вход успешен, выполните необходимые действия (например, установка сессии пользователя)
		fmt.Fprintln(w, "Вход успешен")
	} else {
		// Неверные учетные данные
		fmt.Fprintln(w, "Неверные учетные данные")
	}
}

/*
// AuthenticateUser Функция для проверки введенных учетных данных пользователя
func AuthenticateUser(username, password string) (*User, error) {
	// Здесь вы можете реализовать логику проверки пользователя
	// Например, сравнение данных с данными в базе данных или другом источнике данных

	// Создайте подключение к базе данных
	db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Выполните запрос к базе данных для получения пользователя по логину
	row := db.QueryRow("SELECT id_user, login_user, password_user, name_user, nickname FROM users WHERE login_user = ?", username)

	// Инициализируйте переменные для хранения значений из базы данных
	var id int
	var login, userPassword, name, nickname string

	// Сканируйте значения из результата запроса в переменные
	err = row.Scan(&id, &login, &userPassword, &name, &nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			// Пользователь не найден, возвращаем ошибку
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	// Проверьте, совпадают ли введенные пароль и пароль пользователя из базы данных
	if password != userPassword {
		// Неверный пароль, возвращаем ошибку
		return nil, errors.New("неправильный логин или пароль")
	}

	// Создайте объект User с данными пользователя из базы данных
	user := &User{
		Id:           id,
		Login:        login,
		UserPassword: userPassword,
		Name:         name,
		Nickname:     nickname,
	}

	// Возвращаем пользовательскую запись и ошибку (если есть)
	return user, nil
}

// LoginHandler Функция для обработки запроса на вход пользователя
func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Получение данных пользователя из формы входа (например, имя пользователя и пароль)
	username := r.FormValue("login")
	password := r.FormValue("password")

	// Проверка пользователя на аутентификацию
	_, err := AuthenticateUser(username, password)
	if err != nil {
		// Неверные учетные данные пользователя, отображение ошибки или перенаправление на страницу ошибки
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	// Аутентификация прошла успешно, перенаправление пользователя на домашнюю страницу
	http.ServeFile(w, r, "home.html")
}
*/
