package handler

import (
	"fmt"
	"net/http"
)

func getGasto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET method.")
}
