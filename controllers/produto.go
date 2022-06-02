package controllers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/nenodias/site-golang/db"
	"github.com/nenodias/site-golang/models"
	"github.com/nenodias/site-golang/utils"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var conn *sql.DB

func Init() {
	db.Init()
	var err error
	conn, err = db.GetConnection()
	if err != nil {
		log.Println("Erro na conexao")
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := db.SelectAll(conn)
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := utils.ToFloat64(r.FormValue("preco"))
		if err != nil {
			log.Println("Erro na conversão do preço: ", err)
		}
		quantidade, err := utils.ToInt(r.FormValue("quantidade"))
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

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Converter id ")
	id, err := utils.ToInt64(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Erro ao converter id: ", err)
	} else {
		log.Println("Deletando registro: ", id)
		err = db.Delete(conn, id)
		if err != nil {
			log.Println("Erro ao deletar produto: ", err)
		}
	}
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ToInt64(r.URL.Query().Get("id"))
	var produto *models.Produto = nil
	if err != nil {
		log.Println("Erro ao converter id: ", err)
	} else {
		log.Println("Editando registro: ", id)
		produto, err = db.SelectById(conn, id)
		if err != nil {
			log.Println("Erro ao deletar produto: ", err)
		}
	}
	if produto != nil {
		temp.ExecuteTemplate(w, "Edit", &produto)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := utils.ToInt64(r.FormValue("id"))
		if err == nil {
			nome := r.FormValue("nome")
			descricao := r.FormValue("descricao")
			preco, err := utils.ToFloat64(r.FormValue("preco"))
			if err != nil {
				log.Println("Erro na conversão do preço: ", err)
			}
			quantidade, err := utils.ToInt(r.FormValue("quantidade"))
			if err != nil {
				log.Println("Erro na conversão da quantidade: ", err)
			}
			p := models.Produto{Id: id, Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade}
			err = db.Update(conn, &p)
			if err != nil {
				log.Println("Erro ao salvar novo produto: ", err)
			}
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
