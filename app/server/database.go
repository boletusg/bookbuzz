package server

import (
	_ "database/sql"
	_ "fmt"
	_ "github.com/denisenkom/go-mssqldb" // Импортируйте ваш драйвер базы данных
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "log"
	_ "net/http"
)

var Db *sqlx.DB

func InitDB() (err error) {
	//строка, содержащая данные для подключения к БД в следующем формате:
	//login:password@tcp(host:port)/dbname
	var dataSourceName = "server=boletusg;integrated security=SSPI;database=bookbuzz"
	//подключаемся к БД, используя нужный драйвер и данные для подключения
	Db, err = sqlx.Connect("mssql", dataSourceName)
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}
