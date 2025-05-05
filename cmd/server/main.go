package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println("Received a datadog alert:")
	fmt.Printf("Body:\n%s\n", string(body))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Alert received")
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	port := "8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
