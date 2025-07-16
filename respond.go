package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/s-hammon/p"
)

func respondJSON(w http.ResponseWriter, code int, body map[string]string) {
	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			http.Error(w, p.Format("json.Marshal: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(j)))
		w.WriteHeader(code)
		if _, err := fmt.Fprintf(w, "%s", j); err != nil {
			http.Error(w, p.Format("fmt.Fprintf: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
