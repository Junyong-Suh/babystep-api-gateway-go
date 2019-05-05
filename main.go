package main

import (
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "strings"

    "github.com/gorilla/mux"

    h "./handlers"
)

func main() {
    r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", h.RootHandler)
    r.HandleFunc("/health", h.HealthcheckHandler)

    // get version
    version, _ := ioutil.ReadFile("./VERSION")
    v := strings.TrimSpace(string(version))

    // Bind to a port and pass our router in
    log.Print(fmt.Sprintf("Starting babystep-api-gateway %s on port 8000", v))
    log.Fatal(http.ListenAndServe(":8000", r))
}
