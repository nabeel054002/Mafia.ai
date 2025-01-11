package utils

import (
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
	log.Println(message)
}
