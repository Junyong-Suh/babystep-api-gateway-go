package handlers

import (
    "net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
    // Add Unit Tests
    // https://github.com/Junyong-Suh/babystep-api-gateway-go/issues/16
    w.WriteHeader(resp.StatusCode)
}
