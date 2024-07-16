package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func main() {
    port := flag.String("port", "8080", "Port to listen on")
    message := flag.String("message", "Hello, World!", "Message to respond with")
    showRequest := flag.Bool("m", false, "Show request details")
    flag.Parse()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        clientIP := r.RemoteAddr
        fmt.Fprintln(w, *message)
        if *showRequest {
            fmt.Printf("Client IP: %s\nRequest: %s\n", clientIP, r)
        } else {
            fmt.Printf("Client IP: %s\n", clientIP)
        }
    })

    fmt.Printf("Server is listening on port %s...\n", *port)
    log.Fatal(http.ListenAndServe(":"+*port, nil))
}
