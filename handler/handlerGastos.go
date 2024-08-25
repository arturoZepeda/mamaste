package handler

import (
	"fmt"
	"net/http"
)

func getGastos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET method.")
}

func getGastoID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "GET method ID:", id)
}

func postGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST method.")
}

func putGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT method.")
}

func deleteGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE method.")
}
