package main

import (
    "context"
    "flag"
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "strings"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/mux"

    h "github.com/Junyong-Suh/babystep-api-gateway-go/handlers"
)

func main() {
    var wait time.Duration
    flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.Parse()

    r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", h.RootHandler)
    r.HandleFunc("/health", h.HealthcheckHandler)
    r.HandleFunc("/echo", h.EchoHandler)
    port := "8080"

    srv := &http.Server{
        Addr:         fmt.Sprintf("0.0.0.0:%s", port),
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }

    // get version
    version, _ := ioutil.ReadFile("./VERSION")
    v := strings.TrimSpace(string(version))

    // Run our server in a goroutine so that it doesn't block.
    go func() {
        // Bind to a port and pass our router in
        log.Print(fmt.Sprintf("Starting babystep-api-gateway %s on port %s", v, port))
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()

    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)

    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println(fmt.Sprintf("Shutting down babystep-api-gateway %s on port %s", v, port))
    os.Exit(0)
}
