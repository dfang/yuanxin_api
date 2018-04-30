package util

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, httpStatusCode int, statusCode int, message string) {
	// RespondWithJSON(w, code, map[string]string{"error": message})
	RespondWithJSON(w, httpStatusCode, struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"msg"`
	}{
		StatusCode: statusCode,
		Message:    message,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
	// return
}

// func RespondWhenSuccess(staus_code int, payload interface{}) {
// 	// 	response, _ := json.Marshal(payload)
// 	// w.Header().Set("Content-Type", "application/json; charset=utf-8")

// 	RespondWithJSON(w, http.StatusOK, struct {
// 		StatusCode string      `json:"status_code"`
// 		Message    string      `json:"msg"`
// 		User       *model.User `json:"user"`
// 	}{
// 		StatusCode: "200",
// 		Message:    "更新成功",
// 		User:       user,
// 	})
// }
