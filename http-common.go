package commons

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHeader(r *http.Request, headerKey string) (string, error) {
	header := r.Header.Get(headerKey)
	if header == "" {
		return "", fmt.Errorf("header [%s] was not found in the request", headerKey)
	}
	return header, nil
}

func RespondSuccess(w http.ResponseWriter, code int, message string) (int, error) {
	return RespondJSON(w, code, map[string]string{"message": message})
}

func RespondError(w http.ResponseWriter, code int, message string) (int, error) {
	return RespondJSON(w, code, map[string]string{"error": message})
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) (int, error) {
	h := map[string]string{"Content-Type": "application/json"}
	response, err := json.Marshal(payload)
	if err != nil {
		return Respond(w, http.StatusInternalServerError, []byte(err.Error()), h)
	}
	return Respond(w, status, response, h)
}

func Respond(w http.ResponseWriter, s int, r []byte, h map[string]string) (int, error) {
	for key, value := range h {
		w.Header().Set(key, value)
	}
	w.WriteHeader(s)
	return w.Write(r)
}
