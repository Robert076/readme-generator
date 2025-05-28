package api

import (
	"encoding/json"
	"net/http"

	"github.com/Robert076/readme-generator/internal/generator"
)

func methodNotAllowed(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return true
	}
	return false
}

func GetReadme(w http.ResponseWriter, r *http.Request) {
	if methodNotAllowed(w, r, http.MethodGet) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(FormatPOST); err != nil {
		http.Error(w, "Failed to encode the format for the POST endpoint", http.StatusInternalServerError)
		return
	}
}

func PostReadme(w http.ResponseWriter, r *http.Request) {
	if methodNotAllowed(w, r, http.MethodPost) {
		return
	}

	var dict map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&dict); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	markdown, err := generator.GenerateReadme(dict)
	if err != nil {
		http.Error(w, "Failed to generate README: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/markdown")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(markdown))
}
