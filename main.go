package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    fmt.Println("Starting Go Sync Service...")
    redisURL := os.Getenv("REDIS_URL")
    dbURL := os.Getenv("DATABASE_URL")
    log.Println("Redis URL:", redisURL)
    log.Println("Postgres URL:", dbURL)

    // TODO: implement Redis queue consumer and tab merge logic
    select {} // keep running
}