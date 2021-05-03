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

func RespondSuccess(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"message": message})
}

func RespondSuccessWithSession(w http.ResponseWriter, code int, message, session string) {
	RespondJSONWithSession(w, code, map[string]string{"message": message}, session)
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

func RespondErrorWithSession(w http.ResponseWriter, code int, message, session string) {
	RespondJSONWithSession(w, code, map[string]string{"error": message}, session)
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	h := map[string]string{"Content-Type": "application/json"}
	response, err := json.Marshal(payload)
	if err != nil {
		Respond(w, http.StatusInternalServerError, []byte(err.Error()), h)
	}
	Respond(w, status, response, h)
}

func RespondJSONWithSession(w http.ResponseWriter, status int, payload interface{}, session string) {
	h := map[string]string{"Content-Type": "application/json", "x-router-session": session}
	response, err := json.Marshal(payload)
	if err != nil {
		Respond(w, http.StatusInternalServerError, []byte(err.Error()), h)
	}
	Respond(w, status, response, h)
}

func Respond(w http.ResponseWriter, s int, r []byte, h map[string]string) {
	for key, value := range h {
		w.Header().Set(key, value)
	}
	w.WriteHeader(s)
	_, _ = w.Write(r)
}
