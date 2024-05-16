package model

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
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
	if r.Method == "GET" {
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

		// Подключение к базе данных
		db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// Выполнение SQL-запроса для получения информации о пользователе по его ID
		query := "SELECT login_user, status_user, biography_user, avatar_user, datereg_user FROM users2 WHERE id_user = ?"
		row := db.QueryRow(query, userID)

		var accountData AccountData
		var login, status, bio, avatarURL, dateReg sql.NullString
		err = row.Scan(&login, &status, &bio, &avatarURL, &dateReg)
		if err != nil {
			// Обработка ошибки выполнения запроса
			if err == sql.ErrNoRows {
				log.Println("Нет данных о пользователе")
				http.Error(w, "User data not found", http.StatusNotFound)
				return
			}
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Присвоение значений из sql.NullString в указатели в структуре AccountData
		if login.Valid {
			accountData.Login = &login.String
		}
		if status.Valid {
			accountData.Status = &status.String
		}
		if bio.Valid {
			accountData.Bio = &bio.String
		}
		if avatarURL.Valid {
			accountData.AvatarURL = &avatarURL.String
		}
		if dateReg.Valid {
			dateParts := strings.Split(dateReg.String, "T")
			accountData.DateReg = &dateParts[0]
		}
		// Выполнение SQL-запроса для получения заказов пользователя
		query = "SELECT id_order, title_order, login_user, text_order FROM orders2 INNER JOIN users2 ON orders2.fk_user = users2.id_user WHERE fk_user = ?"
		rows, err := db.Query(query, userID)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Загрузка шаблона страницы аккаунта
		tmpl, err := template.ParseFiles("public/html/account_edit.html")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Вставка данных пользователя в шаблон и отправка его клиенту
		err = tmpl.Execute(w, accountData)
		if err != nil {
			// Обработка ошибки выполнения шаблона
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		var response Response
		login := r.FormValue("login")
		status := r.FormValue("status")
		bio := r.FormValue("biography")

		// Подключение к базе данных
		db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Выполнение SQL-запроса для обновления записи
		query := "UPDATE users2 SET login_user = ?, status_user = ?, biography_user = ? WHERE id_user = ?"
		_, err = db.Exec(query, login, status, bio, userID)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		response.Success = true
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Отправка ответа в формате JSON
		w.Write(jsonResponse)
	}
}
