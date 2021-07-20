// Package response functions to build responses for HTTP requests
package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, message interface{}, statusCode int) {
	if message == nil {
		w.WriteHeader(statusCode)
		return
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}
