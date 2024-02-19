// motd-api.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Message struct represents the JSON response format
type Message struct {
	Message string `json:"message"`
}

// Configuration struct holds the configuration variables
type Configuration struct {
	Message string
	Port    string
}

// Global variable for configuration
var config Configuration

// handler function to handle HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	response := Message{Message: config.Message}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	// Parse command line arguments
	flag.StringVar(&config.Message, "MESSAGE", "Default Message", "Message to be returned by the API")
	flag.StringVar(&config.Port, "APP_PORT", "8080", "Port to listen on")
	flag.Parse()

	// Override configuration with environment variables if set
	if envMessage := os.Getenv("MESSAGE"); envMessage != "" {
		config.Message = envMessage
	}
	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		config.Port = envPort
	}

	// Set up HTTP handler
	http.HandleFunc("/", handler)

	// Start the HTTP server
	fmt.Printf("starting app on port %s\n", config.Port)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
