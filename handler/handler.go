package handler

import (
	"fmt"
	"net/http"
)

func GetGastos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET method desde handler.go")
}

func GetGastoID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "GET method ID: %s", id)
}

func PostGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST method.")
}

func PutGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT method.")
}

func DeleteGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE method.")
}

func GetIngresos(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Get method desde handler")
}

func GetIngresoID(w http.ResponseWriter, r *http.Request){
  id := r.URL.Query().Get("id")
  fmt.Fprintf(w, "Get method con ID: %s",id)
}

func PostIngreso(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "POST method ")
}

func PutIngreso(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Put method")
}

func DeleteIngreso(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "DELETE method")
}
