package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    fmt.Println("Starting Go Sync Service with SQLite...")
    db, err := sql.Open("sqlite3", "./data/keepingtab.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Example: read tab count
    row := db.QueryRow("SELECT COUNT(*) FROM tabs")
    var count int
    if err := row.Scan(&count); err == nil {
        log.Println("Tabs in DB:", count)
    }

    // TODO: implement sync logic
    select {} // block forever
}