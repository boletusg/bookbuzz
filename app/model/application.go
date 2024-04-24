package model

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type app struct {
	ctx context.Context
}

func (a app) Routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", a.StartPage)
}

func (a app) StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lp := filepath.Join("public", "html", "home.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		http.Error(
			rw,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	err = tmpl.ExecuteTemplate(rw, "motivation", nil)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}

func NewApp(ctx context.Context) *app {
	return &app{ctx}
}

func (a app) LoginPage(rw http.ResponseWriter, message string) {
	lp := filepath.Join("public", "html", "login.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	type answer struct {
		Message string
	}
	data := answer{message}
	err = tmpl.ExecuteTemplate(rw, "login", data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
