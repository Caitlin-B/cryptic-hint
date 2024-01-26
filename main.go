package main

import (
    "fmt"
    "github.com/Caitlin-B/cryptic-hint/internal/server"
    "log"
)

const portNum string = ":8080"

func main() {
    srv := server.NewHTTPServer(portNum)
    fmt.Printf("starting server on port %s", portNum)
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
