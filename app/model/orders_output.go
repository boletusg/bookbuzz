package model

import (
	"database/sql"
	"encoding/json"
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

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal(err)
	}
	session.Values["orderID"] = orderID
	err = session.Save(r, w)
	if err != nil {
		log.Fatal(err)
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

// RespondHandler Обработчик для записи отклика в базу данных
func RespondHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Fatal(err)
	}
	orderID, ok := session.Values["orderID"].(string)
	if !ok {
		http.Error(w, "Не найден идентификатор заказа", http.StatusBadRequest)
		return
	}

	userID, ok := session.Values["userID"].(int)
	if !ok {
		// Пользователь не аутентифицирован, перенаправление на страницу входа
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	log.Printf("ID user: %v", userID)
	log.Printf("ID order: %v", orderID)
	// Получение идентификатора заказа из параметров запроса или пути URL

	// Подключаемся к базе данных
	db, err := sql.Open("mssql", "server=boletusg;integrated security=SSPI;database=bookbuzz")
	if err != nil {
		http.Error(w, "Ошибка при подключении к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Проверяем, есть ли уже отклик от этого пользователя на этот заказ
	var count int
	query := "SELECT COUNT(*) FROM response2 WHERE fk_order = ? AND fk_user = ?"
	err = db.QueryRow(query, orderID, userID).Scan(&count)
	if err != nil {
		log.Printf("Ошибка при проверке наличия отклика: %v", err)
		http.Error(w, "Ошибка при обработке отклика", http.StatusInternalServerError)
		return
	}
	var response Response
	if count > 0 {
		// Отклик уже существует, выводим сообщение
		w.WriteHeader(http.StatusOK)
		response.Message = "Вы уже оставляли отклик на этот заказ."
		return
	} else {
		// Записываем отклик в таблицу response2
		query = "INSERT INTO response2 (fk_order, fk_user) VALUES (?, ?)"
		_, err = db.Exec(query, orderID, userID)
		if err != nil {
			http.Error(w, "Ошибка при записи отклика в базу данных", http.StatusInternalServerError)
			return
		}
		response.Success = true
	}

	// Преобразование данных в формат JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	response.Message = "Ваш отклик успешно записан."
	_, _ = w.Write(jsonResponse)
}
