package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Определение флагов командной строки
	port := flag.String("port", "8080", "Port to listen on")
	message := flag.String("message", "Hello, World!", "Message to respond with")
	monitor := flag.Bool("m", false, "Monitor and log client IP and request details")
	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("Usage: simplegowebserver -port <port> -message <message> [-m]")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		fmt.Fprintln(w, *message)
		if *monitor {
			fmt.Printf("Client IP: %s\nRequest: %s\n", clientIP, r)
		} else {
			fmt.Printf("Client IP: %s\n", clientIP)
		}
	})

	fmt.Printf("Server is listening on port %s...\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
