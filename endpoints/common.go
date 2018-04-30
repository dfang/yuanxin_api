package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

// type Adapter func(http.Handler) http.Handler

// MissParameterError Missing parameter
type MissParameterError struct {
	Parameter string
	Error     error
}

// ParseError parsing error
type ParseError struct {
	Parameter string
	Error     error
}

// RecordNotFound record not found in db
type RecordNotFound struct {
	Error error
}

// CheckRequiredParameter check required post form parameter
func CheckRequiredParameter(r *http.Request, s string) {
	if r.PostFormValue(s) == "" {
		panic(&MissParameterError{
			Parameter: s,
			Error:     fmt.Errorf("%s 参数缺失", s),
		})
	}
}

// CheckRequiredParameters check required post form parameters
func CheckRequiredParameters(r *http.Request, params ...string) {
	for _, num := range params {
		CheckRequiredParameter(r, num)
	}
}

// CheckRequiredQueryString check required querystring
func CheckRequiredQueryString(r *http.Request, s string) {
	qs := r.URL.Query()
	if qs.Get(s) == "" {
		panic(&MissParameterError{
			Parameter: s,
			Error:     fmt.Errorf("querystring %s 缺失", s),
		})
	}
}

// CheckRequiredQueryStrings check required querystrings
func CheckRequiredQueryStrings(r *http.Request, qs ...string) {
	for _, num := range qs {
		CheckRequiredQueryString(r, num)
	}
}

// ParseParameterToInt parse parameter to int
func ParseParameterToInt(r *http.Request, s string) int {
	i, err := strconv.Atoi(r.PostFormValue(s))
	if err != nil {
		panic(&ParseError{
			Parameter: s,
			Error:     err,
		})
	}
	return i
}

// func MustParams(h http.Handler, params ...string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		q := r.PostForm
// 		for _, param := range params {
// 			if len(q.Get(param)) == 0 {
// 				http.Error(w, "missing "+param, http.StatusBadRequest)
// 				return // exit early
// 			}
// 		}
// 		h.ServeHTTP(w, r) // all params present, proceed
// 	})
// }

// PanicIfNotNil panic if err not nil
func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

// PayLoadFrom payload
type PayLoadFrom struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

// New PayLoadFrom.New
func (p PayLoadFrom) New(status int, msg string) PayLoadFrom {
	return PayLoadFrom{
		StatusCode: status,
		Message:    msg,
	}
}

// GetUIDFromContext get uid from decoded jwt in context(set by jwt middleware)
func GetUIDFromContext(r *http.Request) int {
	claims := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
	var u float64
	u = claims["uid"].(float64)
	return int(u)
}
