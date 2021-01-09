package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteError(rw http.ResponseWriter, err error) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(500)
	e := json.NewEncoder(rw).Encode(struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("Reading tree failed: %s", err.Error()),
	})
	if e != nil {
		log.Printf("Error writing response: %s", e)
	}
}
