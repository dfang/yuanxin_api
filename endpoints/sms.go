package endpoints

import "net/http"

func SendSMSEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}

func ValidateSMSEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}
