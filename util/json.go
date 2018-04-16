package util

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}

type NullString sql.NullString

func (ns NullString) MarshalJSON() ([]byte, error) {
	// if !ns.Valid {
	// 	return []byte("null"), nil
	// }
	// return json.Marshal(ns.String)
	return json.Marshal(ns.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.String = str
	s.Valid = str != ""
	return nil
}
