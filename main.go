package main

import (
	"log"
	"net/http"

	"github.com/SidhnatPandey/go-bookstore/pkg/config"
	"github.com/SidhnatPandey/go-bookstore/pkg/models"
	"github.com/SidhnatPandey/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	var books models.Book
	err := config.DB.Table("sid.book").AutoMigrate(&books)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
