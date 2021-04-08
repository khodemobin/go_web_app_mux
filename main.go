package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)

	_ = http.ListenAndServe(":8000", nil)
}
