package search

import (
    "fmt"
    "io"
    "net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        fmt.Printf("could not read body: %s\n", err)
    }

    fmt.Fprintf(w, "Got input: %s", string(body))
}
