package controller

import (
	"bookbuzz/app/model"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//указываем путь к файлу с шаблоном
	main := filepath.Join("public", "html", "usersDynamicPage.html")
	//создаем html-шаблон
	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//исполняем именованный шаблон "users", передавая туда массив со списком пользователей
	err = tmpl.ExecuteTemplate(rw, "users", users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}
func AddUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра name, переданного в форме запроса
	name := r.FormValue("name_user")
	//получаем значение из параметра surname, переданного в форме запроса
	login := r.FormValue("login_user")

	//проверяем на пустые значения
	if name == "" || login == "" {
		http.Error(rw, "Имя и фамилия не могут быть пустыми", 400)
		return
	}
	//создаем новый объект
	user := model.NewUser(name, login)
	//записываем нового пользователя в таблицу БД
	err := user.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//возвращаем текстовое подтверждение об успешном выполнении операции
	err = json.NewEncoder(rw).Encode("Пользователь успешно добавлен!")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
