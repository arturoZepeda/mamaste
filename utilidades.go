package main

import (
	"fmt"
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
