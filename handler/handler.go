package handler

import (
	"fmt"
	"io"
	mamasdb "mamastw/database"
	"net/http"
	"strconv"
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
	mamasdb.UpdateGasto(1, "2024-09-02", "Comida", 12.34111)
}

func DeleteGasto(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // id := r.URL.Query().Get("id")
	idInt, _ := strconv.Atoi(id)
	mama, err := mamasdb.DeleteGasto(idInt)
	if err != nil {
		fmt.Fprintf(w, "Error al eliminar el gasto: %s", err)
	}
	fmt.Fprintf(w, "DELETE method. ID: %s", id)
	fmt.Fprintf(w, "Gasto eliminado: %d", mama)
}

func GetIngresos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get method desde handler")
	FM, err := mamasdb.GetIngresos()
	if err != nil {
		fmt.Fprintf(w, "Error al obtener los ingresos: %s", err)
		return
	}
	fmt.Fprintf(w, "Ingresos: %s", FM)
}

func GetIngresoID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	FM, err := mamasdb.selectIngresoId(id)
	// FM, err := mamasdb.selectIngresoId(id)
	if err != nil {
		fmt.Fprintf(w, "Error al obtener el ingreso: %s", err)
	}
	fmt.Fprintf(w, "Ingreso: %s", FM)

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
