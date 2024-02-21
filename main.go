package main

import (
    "database/sql"
    "fmt"
    "github.com/Caitlin-B/cryptic-hint/internal/server"
    cfg "github.com/spf13/viper"
    "log"
)

const (
    CfgPortNum    = "PORT"
    cfgDBHost     = "DB_HOST"
    cfgDBPort     = "DB_PORT"
    cfgDBUser     = "DB_USER"
    cfgDBPassword = "DB_PASSWORD"
    cfgDBName     = "DB_NAME"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        cfg.GetString(cfgDBHost), cfg.GetString(cfgDBPort), cfg.GetString(cfgDBUser),
        cfg.GetString(cfgDBPassword), cfg.GetString(cfgDBName))
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    srv := server.NewHTTPServer(cfg.GetString(CfgPortNum), db)
    fmt.Printf("starting server on port %s", cfg.GetString(CfgPortNum))
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
