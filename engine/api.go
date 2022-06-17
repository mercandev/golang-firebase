package engine

import (
	"gofirestorre/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeApi() {

	r := mux.NewRouter()
	r.HandleFunc("/post", service.AddData).Methods("GET")
	r.HandleFunc("/get", service.GetData).Methods("GET")
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8001",
	}

	log.Fatal(srv.ListenAndServe())
}
