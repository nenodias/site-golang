package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/nenodias/site-golang/db"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var conn *sql.DB

func main() {
	db.Init()
	var err error
	conn, err = db.GetConnection()
	if err != nil {
		fmt.Println("Erro na conexao")
		panic(err)
	}

	http.HandleFunc("/", index)
	fmt.Println(temp)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := db.SelectAll(conn)
	fmt.Println(produtos)
	temp.ExecuteTemplate(w, "Index", produtos)
}
