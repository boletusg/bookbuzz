package model

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

type Respond struct {
	ID     int    `db:"id_response"`
	Title  string `json:"title_order" db:"title_order"`
	User   string `json:"login_user" db:"login_user"`
	UserID int    `json:"id_user" db:"id_user"`
}

func RespondsHandler(w http.ResponseWriter, r *http.Request) {
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

	// Выполнение запроса для входящих откликов
	incomingRespondsQuery := `
    SELECT o.title_order, u.login_user, u.id_user
    FROM orders2 o
    JOIN response2 r ON o.id_order = r.fk_order
    JOIN users2 u ON r.fk_user = u.id_user
    WHERE o.fk_user = ?
    `
	incomingRespondsRows, err := db.Query(incomingRespondsQuery, userID)
	if err != nil {
		log.Fatal(err)
	}
	defer incomingRespondsRows.Close()

	// Обработка результатов для входящих откликов
	var incomingResponds []Respond
	for incomingRespondsRows.Next() {
		var respond Respond
		if err := incomingRespondsRows.Scan(&respond.Title, &respond.User, &respond.UserID); err != nil {
			log.Fatal(err)
		}
		incomingResponds = append(incomingResponds, respond)
	}

	// Выполнение запроса для исходящих откликов
	outgoingRespondsQuery := `
    SELECT r.id_response, o.title_order, u.login_user, u.id_user
    FROM response2 r
    JOIN orders2 o ON r.fk_order = o.id_order
    JOIN users2 u ON o.fk_user = u.id_user
    WHERE r.fk_user = ?
    `
	outgoingRespondsRows, err := db.Query(outgoingRespondsQuery, userID)
	if err != nil {
		log.Fatal(err)
	}
	defer outgoingRespondsRows.Close()

	// Обработка результатов для исходящих откликов
	var outgoingResponds []Respond
	for outgoingRespondsRows.Next() {
		var respond Respond
		if err := outgoingRespondsRows.Scan(&respond.ID, &respond.Title, &respond.User, &respond.UserID); err != nil {
			log.Fatal(err)
		}
		outgoingResponds = append(outgoingResponds, respond)
	}

	// Загрузка шаблона HTML
	tmpl, err := template.ParseFiles("public/html/my_orders.html")
	if err != nil {
		log.Fatal(err)
	}

	// Передача данных заказов в шаблон и их вывод на страницу
	err = tmpl.Execute(w, map[string]interface{}{
		"OutgoingResponds": outgoingResponds,
		"IncomingResponds": incomingResponds,
	})
	if err != nil {
		log.Fatal(err)
	}
}
