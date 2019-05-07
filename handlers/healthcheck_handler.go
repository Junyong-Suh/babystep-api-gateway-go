package handlers

import (
    "net/http"
    "fmt"
    "io"
    "io/ioutil"
    "strings"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
    // A very simple health check.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    version, _ := ioutil.ReadFile("./VERSION")
    v := strings.TrimSpace(string(version))

    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, fmt.Sprintf("{\"version\": \"%s\", \"service\": \"babystep-api-gateway\"}", v))
}
