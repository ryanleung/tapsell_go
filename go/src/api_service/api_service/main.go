package main

import (
	//"fmt"
	"flag"
	"api_service/controllers"
	"api_service/base"
	"log"
)

var port = flag.Int("port", 8000, "Theport on which to listen for requests")
var databaseSSL = flag.Bool(
	"db_ssl",
	false,
	"Whether to attempt to use SSL")
var databaseUri = flag.String(
	"db_uri",
	"postgres",
	"URI of postgres database to connect with")

func main() {
	flag.Parse()

	log.Println("Creating Message Store...")
	ms := base.NewMessageStore()
	log.Println("Message Store Initialized.")

	jsonService := controllers.JsonService{ms}
	jsonService.Serve(*port)
}
