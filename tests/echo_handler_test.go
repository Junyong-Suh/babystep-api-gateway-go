package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/thedevsaddam/govalidator"

    h "../handlers"
)

type echo_response struct {
    method  string      `json:"method"`
    path    string      `json:"path"`
    args    string      `json:"args"`
    body    string      `json:"body"`
    headers []string    `json:"headers"`
    uuid    string      `json:"uuid"`
}

func TestEchoHandler(t *testing.T) {
    // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
    // pass 'nil' as the third parameter.
    req, err := http.NewRequest("GET", "/echo", nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(h.EchoHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    // expected := `Hello`
    // if rr.Body.String() != expected {
    //     t.Errorf("handler returned unexpected body: got %v want %v",
    //         rr.Body.String(), expected)
    // }

    rules := govalidator.MapData{
        // "username": []string{"required", "between:3,8"},
        // "email":    []string{"required", "min:4", "max:20", "email"},
        // "web":      []string{"url"},
        // "phone":    []string{"digits:11"},
        // "agree":    []string{"bool"},
        // "dob":      []string{"date"},
        "method":   []string{"required"},
        "path":     []string{"required"},
        "args":     []string{"required"},
        "body":     []string{"required"},
        "headers":  []string{"required"},
        "uuid":     []string{"required"},
	}

	opts := govalidator.Options{
        Request:         req,      // request object
        Rules:           rules,    // rules map
        RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
    fmt.Print(e)
}
