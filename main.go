package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type EncryptionRequest struct {
	Value string `json:"value"`
}

type EncryptionResponse struct {
	EncryptedValue string `json:"encrypted_value"`
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var req EncryptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perform encryption logic here

	resp := EncryptionResponse{
		EncryptedValue: "encrypted_value_placeholder",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var req EncryptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perform decryption logic here

	resp := EncryptionResponse{
		EncryptedValue: "decrypted_value_placeholder",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")
	r.HandleFunc("/api/decrypt", DecryptHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
