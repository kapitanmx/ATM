package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
	"github.com/kapitanmx/ATM/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Routes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
