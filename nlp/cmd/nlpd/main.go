package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"nlp"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read from request", http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		http.Error(w, "missing data", http.StatusBadRequest)
		return
	}

	text := string(data)

	tokens := nlp.Tokenize(text)

	resp := map[string]any{
		"tokens": tokens,
	}
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't marshal response", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
