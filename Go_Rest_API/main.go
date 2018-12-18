package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

func main(){
	//init router
	r := mux.NewRouter()

	//Route Handlers /Endpoints
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")

	//to run the server
	log.Fatal(http.ListenAndServe(":8000",r))


}