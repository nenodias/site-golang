package main

import (
	"net/http"

	"github.com/nenodias/site-golang/routes"
)

func main() {
	routes.Register()
	http.ListenAndServe(":8080", nil)
}
