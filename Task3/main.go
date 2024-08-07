package main

import (
    "log"
    "net/http"
)

func main() {
    // Set up the HTTP handler
    http.HandleFunc("/", handleRequestAndRedirect)
    log.Println("Starting forward proxy server on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
