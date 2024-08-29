package main

import (
	"fmt"
	mamasdb "mamastw/database"
	"net/http"
)

func logRequest(r *http.Request) {
	fmt.Println("==============================")
	fmt.Println("Nuevo request:")
	fmt.Println("Request received from", r.RemoteAddr)
	fmt.Println("Request method:", r.Method)
	fmt.Println("Request URI:", r.RequestURI)
	fmt.Println("Request headers:", r.Header)
	fmt.Println("Request body:", r.Body)
	fmt.Println("Request form:", r.Form)
	fmt.Println("Request post form:", r.PostForm)
	fmt.Println("Request multipart form:", r.MultipartForm)
	fmt.Println("Request cookies:", r.Cookies())
	fmt.Println("Request user agent:", r.UserAgent())
	fmt.Println("Request referer:", r.Referer())
	fmt.Println("==============================")
}

func inicializa() {
	mamasdb.CreateDB()
	/*
	   http.HandleFunc("/gastos", handler.GetGastos)

	   	http.HandleFunc("/gastos/", handler.GetGastoID)
	   	http.HandleFunc("/gastos", handler.PostGasto)
	   	http.HandleFunc("/gastos/", handler.PutGasto)
	   	http.HandleFunc("/gastos/", handler.DeleteGasto)
	   	http.HandleFunc("/ingresos", handler.GetIngresos)
	   	http.HandleFunc("/ingresos/", handler.GetIngresoID)
	   	http.HandleFunc("/ingresos", handler.PostIngreso)
	   	http.HandleFunc("/ingresos/", handler.PutIngreso)
	   	http.HandleFunc("/ingresos/", handler.DeleteIngreso)
	*/
}
