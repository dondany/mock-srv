package main

import (
	"fmt"
	"log"
	"mocksrv/db"
	"mocksrv/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	database, err := db.LoadDatabase("db.json")
	if err != nil {
		panic(err)
	}

	for collectionKey, _ := range database.Data {
		handler := handlers.Handler{
			Collection: collectionKey,
			DB:         database,
		}
		mux.HandleFunc(fmt.Sprintf("GET /%v", collectionKey), handler.GetAll)
		mux.HandleFunc(fmt.Sprintf("GET /%v/{id}", collectionKey), handler.GetById)
		mux.HandleFunc(fmt.Sprintf("POST /%s", collectionKey), handler.Post)
		mux.HandleFunc(fmt.Sprintf("PUT /%s/{id}", collectionKey), handler.Put)
		mux.HandleFunc(fmt.Sprintf("DELETE /%s/{id}", collectionKey), handler.Delete)
	}

	log.Fatal(http.ListenAndServe(":8000", mux))
}
