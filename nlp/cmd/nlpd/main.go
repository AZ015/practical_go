package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"nlp"
)

type Server struct {
	logger *log.Logger
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.logger.Printf("error: method not allowed")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Printf("error: can't read from request - %s", err)
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
		s.logger.Printf("error: can't marshal response - %s", err)
		http.Error(w, "can't marshal response", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	logger := log.New(log.Writer(), "[nlp] ", log.LstdFlags|log.Lshortfile)
	s := Server{
		logger: logger,
	}
	//http.HandleFunc("/health", healthHandler)
	//http.HandleFunc("/tokenize", tokenizeHandler)

	r := mux.NewRouter()
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	addr := ":8080"
	s.logger.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
