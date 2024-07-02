package main

import (
	"fmt"
	"keyvaluestore/server"
	"net/http"
)

func main() {
	http.HandleFunc("/pair", server.CreatePairHandler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
