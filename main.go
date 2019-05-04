package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    h "./handlers"
)

func main() {
    r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", h.RootHandler)
    r.HandleFunc("/health", h.HealthcheckHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
