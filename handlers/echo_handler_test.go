package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type HttpHeaders struct {
	HTTP_VERSION           string `json:"HTTP_VERSION"`
	HTTP_HOST              string `json:"HTTP_HOST"`
	HTTP_ACCEPT_ENCODING   string `json:"HTTP_ACCEPT_ENCODING"`
	HTTP_USER_AGENT        string `json:"HTTP_USER_AGENT"`
	HTTP_X_FORWARDED_FOR   string `json:"HTTP_X_FORWARDED_FOR"`
	HTTP_X_FORWARDED_HOST  string `json:"HTTP_X_FORWARDED_HOST"`
	HTTP_X_FORWARDED_PORT  string `json:"HTTP_X_FORWARDED_PORT"`
	HTTP_X_FORWARDED_PROTO string `json:"HTTP_X_FORWARDED_PROTO"`
	HTTP_FORWARDED         string `json:"HTTP_FORWARDED"`
}

type EchoResponse struct {
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Args    string      `json:"args"`
	Body    string      `json:"body"`
	Headers HttpHeaders `json:"headers"`
	Uuid    string      `json:"uuid"`
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
	handler := http.HandlerFunc(EchoHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var echoResponse EchoResponse
	err = json.Unmarshal([]byte(rr.Body.String()), &echoResponse)
	if err != nil {
		t.Fatal(err)
	}

	if echoResponse.Method != "GET" ||
		echoResponse.Path != "/" ||
		echoResponse.Args != "" ||
		echoResponse.Body != "" ||
		reflect.TypeOf(echoResponse.Uuid).String() != "string" ||
		reflect.TypeOf(echoResponse.Headers).String() != "handlers.HttpHeaders" {
		t.Errorf("handler returned unexpected body: %v", rr.Body.String())
	}
}
