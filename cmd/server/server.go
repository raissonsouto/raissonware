package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", uploadFileHandler)

	_ = http.ListenAndServe(":3333", nil)
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		return
	}
}
