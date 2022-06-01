package routes

import (
	"net/http"

	"github.com/nenodias/site-golang/controllers"
)

func Register() {
	controllers.Init()
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/create-database", controllers.CreateDatabase)
}
