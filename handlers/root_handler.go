package handlers

import (
    "net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`
      {
        "serviceName": "babystep-api-gateway",
        "version": "v0.0.1"
      }
    `))
}
