package handlers

import (
    "net/http"
    "io"
    "io/ioutil"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
    // Stupid sinatra app that returns back info about the request
    // https://github.com/3scale/echo-api
    resp, err := http.Get("https://echo-api.3scale.net")
    if err != nil {
      // handle error
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      // handle error
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(resp.StatusCode)
    io.WriteString(w, string(body))
}
