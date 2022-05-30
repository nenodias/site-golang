package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/nenodias/site-golang/db"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id         int64
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	conn, err := db.GetConnection()
	if err != nil {
		fmt.Println("Erro na conexao")
		panic(err)
	}
	err = db.Create(conn)
	if err != nil {
		fmt.Println("Erro no create")
	}
	res, err := conn.Query("SELECT * FROM produto")
	if err != nil {
		fmt.Println("Erro no select")
		panic(err)
	}
	fmt.Println(res.Columns())
	fmt.Println(res.Next())
	http.HandleFunc("/", index)
	fmt.Println(temp)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
