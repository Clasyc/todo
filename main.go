package main

import (
	"github.com/Clasyc/todo/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

const ConnectionString = "user=postgres password=admin dbname=todo host=localhost port=5432 sslmode=disable"

func main() {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	model.InitService(db)

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/create", createHandler)
	r.HandleFunc("/delete/{id:[0-9]+}", deleteHandler)
	r.HandleFunc("/toggle/{id:[0-9]+}", toggleHandler)
	log.Fatal(http.ListenAndServe(":8090", r))
}