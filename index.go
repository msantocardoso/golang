package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/msantocardoso/golang/routers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/go/quicksort/{tamanhoLista}", routers.QuicksortRouter).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}