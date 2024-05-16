package model

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type AccountData struct {
	Id        int         `json:"id_user" db:"id_user"`
	Login     *string     `json:"login_user" db:"login_user"`
	Status    *string     `json:"status_user" db:"status_user"`
	Bio       *string     `json:"biography_user" db:"biography_user"`
	AvatarURL *string     `json:"avatar_user" db:"avatar_user"`
	DateReg   *string     `json:"datereg_user" db:"datereg_user"`
	Orders    []OrderData `json:"orders"`
}
type OrderData struct {
	OrderID int
	Title   string
	Login   string
	Text    string
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Проверка аутентификации пользователя
		session, err := store.Get(r, "session-name")
		if err != nil {
			log.Fatal(err)
		}

		userID, ok := session.Values["userID"].(int)
		if !ok {
			// Пользователь не аутентифицирован, выполните необходимые действия, например, перенаправление на страницу входа
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

		var orders []OrderData

		// Обработка результатов запроса
		for rows.Next() {
			var order OrderData
			err := rows.Scan(&order.OrderID, &order.Title, &order.Login, &order.Text)
			if err != nil {
				log.Fatal(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			orders = append(orders, order)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Присоединение заказов к данным аккаунта
		accountData.Orders = orders

		// Загрузка шаблона страницы аккаунта
		tmpl, err := template.ParseFiles("public/html/account_page.html")
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
	}

}
