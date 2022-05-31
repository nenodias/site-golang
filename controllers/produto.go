package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/nenodias/site-golang/db"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var conn *sql.DB

func Init() {
	db.Init()
	var err error
	conn, err = db.GetConnection()
	if err != nil {
		fmt.Println("Erro na conexao")
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := db.SelectAll(conn)
	fmt.Println(produtos)
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
