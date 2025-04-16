package utilities

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func BuildErrorResponse(message string, err string, data interface{}) ResponseV2 {
	splittedError := strings.Split(err, "\n")
	res := ResponseV2{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}

	return res
}

func HandleError(w http.ResponseWriter, status int, message string, err error) {
	w.WriteHeader(status)
	res := BuildErrorResponse(message, err.Error(), EmptyObj{})
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error sending response: %v", err)
	}
}