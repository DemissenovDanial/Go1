package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req JsonRequest
	err := decoder.Decode(&req)
	if req.Data == "" {
		http.Error(w, "Invalid JSON message: Empty 'data' field", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Println("Received data:", req.Data)

	response := struct {
		Result string `json:"result"`
	}{
		Result: "Data successfully received",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
