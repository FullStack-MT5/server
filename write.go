package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeError writes err to w in JSON format. If err is
// internal, obfuscates message and logs it to stdout.
func writeError(w http.ResponseWriter, err error) {
	e := httpErrorOf(err)
	message := e.Message

	if e.Code >= http.StatusInternalServerError {
		// Obfuscate message.
		message = http.StatusText(e.Code)
	}
	// TODO verbose mode and/or logging to file.
	log.Println(err.Error())

	writeJSON(w, httpError{Message: message}, e.Code)
}

// writeJSON writes data to w in JSON format and sets the response satus code.
func writeJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		writeError(w, &ErrInternal)
		return
	}
}
