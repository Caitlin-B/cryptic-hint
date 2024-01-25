package main

import (
    "fmt"
    "github.com/Caitlin-B/cryptic-hint/search"
    "net/http"
)

const portNum string = ":8080"

func main() {
    http.HandleFunc("/", search.Search)

    err := http.ListenAndServe(portNum, nil)
    if err != nil {
        fmt.Printf("unable to start server: %s", err.Error())
        return
    }
    fmt.Printf("listening on port %s", err.Error())
}
