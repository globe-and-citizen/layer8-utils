package utilities

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type EmptyObj struct{}

type ResponseV2 struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   []string    `json:"error"`
	Data    interface{} `json:"data"`
}

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