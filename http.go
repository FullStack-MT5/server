package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeError writes err to w in JSON format. If err is
// internal, obfuscates message and logs it the stdout.
func writeError(w http.ResponseWriter, err error) {
	code := errorCode(err)
	msg := errorMessage(err)

	if code >= http.StatusInternalServerError {
		// Obfuscate message.
		msg = http.StatusText(code)
	}
	// TODO verbose mode and/or logging to file.
	log.Println(err.Error())

	w.Header().Set("X-Content-Type-Options", "nosniff")

	writeJSON(w, httpError{Message: msg}, code)
}

// writeJSON writes data to w in JSON format and sets the response satus code.
func writeJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		writeError(w, &ErrInternal)
		return
	}
}
