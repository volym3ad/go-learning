package main
import (
        "github.com/gorilla/mux"
        "log"
        "net/http"
      )

func YourHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Gorilla!\n"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("response"))
}

func main() {
  r := mux.NewRouter()
  // Routes consist of a path and a handler function.
  r.HandleFunc("/", YourHandler)
  r.HandleFunc("/hello", HelloHandler)

  // Bind to a port and pass our router in
  log.Fatal(http.ListenAndServe(":8080", r))
}
