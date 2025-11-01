package main

import (
    "log"
    "net/http"
)

func main() {
    // Initialize the application
    // Here you can load configuration and set up dependencies

    // Start the server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}