package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/nenodias/site-golang/db"
	"github.com/nenodias/site-golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var conn *sql.DB

func Init() {
	db.Init()
	var err error
	conn, err = db.GetConnection()
	err = db.Create(conn)
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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade: ", err)
		}
		p := models.Produto{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}
		err = db.Insert(conn, &p)
		if err != nil {
			log.Println("Erro ao salvar novo produto: ", err)
		}
	}
	http.Redirect(w, r, "/", 301)
}

func CreateDatabase(w http.ResponseWriter, r *http.Request) {
	err := db.Create(conn)
	if err != nil {
		log.Println("Erro na conexao")
		panic(err)
	}
	http.Redirect(w, r, "/", 301)
}
