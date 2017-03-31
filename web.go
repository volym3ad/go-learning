package main

import (
	"io"
	"net/http"
	"log"
)

func SimpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "blablabla\n")
} 

func main() {
	http.HandleFunc("/", SimpleServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}