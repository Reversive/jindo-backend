package httpio

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorDetail struct {
	Path    string `json:"path"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RespondWithError(
	w http.ResponseWriter,
	req *http.Request,
	code int,
	responseError error,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	detail := ErrorDetail{
		Path:    req.RequestURI,
		Code:    code,
		Message: responseError.Error(),
	}

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(detail); err != nil {
		log.Printf("CRITICAL: Failed to encode JSON error response for client: %v. Original error was Code: %d, Message: %s", err, code, responseError.Error())
	}

}

func RespondWithJSON(
	w http.ResponseWriter,
	req *http.Request,
	code int,
	data any,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("CRITICAL: Failed to encode JSON data response for client: %v. Original data was Code: %d", err, code)
	}

}
