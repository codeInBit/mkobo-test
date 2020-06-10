package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//JSON - This format JSON responses
func JSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.WriteHeader(statusCode)

	format := struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}{
		Data:    data,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(format)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

//ERROR - This format error messages
func ERROR(w http.ResponseWriter, statusCode int, err error, message string) {
	if err != nil {
		message := err.Error() + ": " + message
		JSON(w, statusCode, nil, message)
		return
	}
	JSON(w, http.StatusBadRequest, nil, message)
}
