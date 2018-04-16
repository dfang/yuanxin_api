package endpoints

import "net/http"

func SessionEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}
