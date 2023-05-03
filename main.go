package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Greeting struct {
    Message string `json:"message"`
}

func main() {
    // Define a handler function for the root endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Create a new Greeting instance
        greeting := Greeting{Message: "Hello, World!"}
        // Convert the Greeting instance to JSON
        json, err := json.Marshal(greeting)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // Set the Content-Type header to indicate that the response is JSON
        w.Header().Set("Content-Type", "application/json")
        // Write the JSON response to the response writer
        w.Write(json)
    })

    // Start the server on port 8000
    log.Fatal(http.ListenAndServe(":8000", nil))
}
