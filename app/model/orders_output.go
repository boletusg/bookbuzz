package model

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/denisenkom/go-mssqldb"
)

type Order struct {
	OrderID string `json:"id_order" db:"id_order"`
	Title   string `json:"title_order" db:"title_order"`
	Login   string `json:"login_user" db:"login_user"`
	Text    string `json:"text_order" db:"text_order"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Подключение к базе данных
	db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выполнение SQL-запроса для получения всех заказов
	query := "SELECT id_order, title_order, login_user, text_order FROM orders2 INNER JOIN users2 ON orders2.fk_user = users2.id_user"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Создание среза для хранения заказов
	var orders []Order

	// Итерация по результатам запроса и добавление заказов в срез
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.OrderID, &order.Title, &order.Login, &order.Text)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Загрузка шаблона HTML
	tmpl, err := template.ParseFiles("public/html/home.html")
	if err != nil {
		log.Fatal(err)
	}

	// Передача данных заказов в шаблон и их вывод на страницу
	err = tmpl.Execute(w, orders)
	if err != nil {
		log.Fatal(err)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	// Получение идентификатора заказа из параметров запроса или пути URL
	orderID := r.URL.Query().Get("id")

	// Проверка, что идентификатор заказа был передан
	if orderID == "" {
		http.Error(w, "Не указан идентификатор заказа", http.StatusBadRequest)
		return
	}

	// Подключение к базе данных
	db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выполнение SQL-запроса для получения информации о конкретном заказе
	query := "SELECT title_order, login_user, text_order FROM orders2 INNER JOIN users2 ON orders2.fk_user = users2.id_user WHERE id_order = ?"
	row := db.QueryRow(query, orderID)

	// Создание переменных для хранения данных о заказе
	var title, login, text string
	err = row.Scan(&title, &login, &text)
	if err != nil {
		log.Fatal(err)
	}

	// Загрузка шаблона HTML для страницы с подробностями заказа
	tmpl, err := template.ParseFiles("public/html/order_page.html")
	if err != nil {
		log.Fatal(err)
	}

	// Передача данных о заказе в шаблон и их вывод на страницу
	data := struct {
		Title string
		Login string
		Text  string
	}{
		Title: title,
		Login: login,
		Text:  text,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
