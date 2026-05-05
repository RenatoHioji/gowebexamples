package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	number := vars["number"]

	fmt.Fprintf(w, "Olá requisição de %s e número %s ", name, number)
}

func handleExample(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Example route")
}

/*
Mux adds specific handlers for HTTP methods such as GET, POST, PUT, PATCH and DELETE. It also can make a endpoint host specific
So in a multi tenancy situation, the endpoint just is hit by the client which aims for a specific host.
It also can be added a Schemes which defines the type of request as http and https.
To end, you can add a path prefix to a router if wanted as in the last example
*/
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/person/{name}/department/{number}", handle).Methods("GET").Host("localhost").Schemes("http")

	exampleRouter := r.PathPrefix("/example").Subrouter()

	exampleRouter.HandleFunc("/", handleExample)

	http.ListenAndServe(":80", r)
}
