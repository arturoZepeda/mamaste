package handler

import (
	"fmt"
	"io"
	mamasdb "mamastw/database"
	"net/http"
)

func GetGastos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET method desde handler")
	FM, err := mamasdb.GetGastos()
	if err != nil {
		fmt.Fprintf(w, "Error al obtener los gastos: %s", err)
		return
	}
	fmt.Fprintf(w, "Gastos: %s", FM)
}

func GetGastoID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "GET method ID: %s", id)
}

func PostGasto(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al leer el body: %s", err)
		return
	}
	// body := r.Body
	fmt.Fprintf(w, "POST method ")
	fmt.Fprintf(w, "Body: %s", body)
	mamasdb.InsertGasto("2021-09-01", "Comida", 12.34)
}

func PutGasto(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintf(w, "Error al leer el id")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al leer el body: %s", err)
		return
	}
	fmt.Fprintf(w, "PUT method.")
	fmt.Fprintf(w, "ID: %s", id)
	fmt.Fprintf(w, "Body: %s", body)
}

func DeleteGasto(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "DELETE method. ID: %s", id)
}

func GetIngresos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get method desde handler")
}

func GetIngresoID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "Get method con ID: %s", id)
}

func PostIngreso(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al leer el body: %s", err)
		return
	}
	fmt.Fprintf(w, "POST method ")
	fmt.Fprintf(w, "Body: %s", body)
}

func PutIngreso(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintf(w, "Error al leer el id")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error al leer el body: %s", err)
		return
	}
	fmt.Fprintf(w, "PUT method.")
	fmt.Fprintf(w, "ID: %s", id)
	fmt.Fprintf(w, "Body: %s", body)
}

func DeleteIngreso(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Fprintf(w, "DELETE method ID: %s", id)
}
