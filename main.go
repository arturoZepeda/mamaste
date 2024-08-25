package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "GET method.")
		case "POST":
			fmt.Fprintf(w, "POST method.")
		case "PUT":
			fmt.Fprintf(w, "PUT method.")
		case "DELETE":
			fmt.Fprintf(w, "DELETE method.")
		}
	})

	http.HandleFunc("/gastos", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "GET method.")
		case "POST":
			fmt.Fprintf(w, "POST method.")
		case "PUT":
			fmt.Fprintf(w, "PUT method.")
		case "DELETE":
			fmt.Fprintf(w, "DELETE method.")
		}
	})

	http.HandleFunc("/ingresos", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "GET method.")
		case "POST":
			fmt.Fprintf(w, "POST method.")
		case "PUT":
			fmt.Fprintf(w, "PUT method.")
		case "DELETE":
			fmt.Fprintf(w, "DELETE method.")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	} else {
		fmt.Println("Server started on port 8080")
	}
}