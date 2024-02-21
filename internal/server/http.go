package server

import (
    "database/sql"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

type server struct {
    // this probably be sql connection
    DB         *sql.DB
    Indicators *Indicators
}

func (s *server) handlePost(w http.ResponseWriter, r *http.Request) {
    var req SearchRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // access psql connection and search DB
    i := s.Search(req.Query)
    json.NewEncoder(w).Encode(i)
}

type SearchRequest struct {
    Query string `json:"query"`
}

func NewHTTPServer(addr string, db *sql.DB) *http.Server {
    server := &server{Indicators: &Indicators{}, DB: db}
    r := mux.NewRouter()
    r.HandleFunc("/", server.handlePost).Methods("POST")
    //r.HandleFunc("/", server.handleGet).Methods("GET") // todo add get all indicators? or all clue types?

    return &http.Server{
        Addr:    addr,
        Handler: r,
    }
}
