package endpoints

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// type Adapter func(http.Handler) http.Handler

type MissParameterError struct {
	Parameter string
	Error     error
}

type ParseError struct {
	Parameter string
	Error     error
}

type RecordNotFound struct {
	Error error
}

func CheckRequiredParameter(r *http.Request, s string) {
	if r.PostFormValue(s) == "" {
		panic(&MissParameterError{
			Parameter: s,
			Error:     errors.New(fmt.Sprintf("%s 参数缺失", s)),
		})
	}
}

func CheckRequiredParameters(r *http.Request, params ...string) {
	for _, num := range params {
		CheckRequiredParameter(r, num)
	}
}

func CheckRequiredQueryString(r *http.Request, s string) {
	qs := r.URL.Query()
	if qs.Get(s) == "" {
		panic(&MissParameterError{
			Parameter: s,
			Error:     errors.New(fmt.Sprintf("querystring %s 缺失", s)),
		})
	}
}
func CheckRequiredQueryStrings(r *http.Request, qs ...string) {
	for _, num := range qs {
		CheckRequiredQueryString(r, num)
	}
}

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

func MustParams(h http.Handler, params ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.PostForm
		for _, param := range params {
			if len(q.Get(param)) == 0 {
				http.Error(w, "missing "+param, http.StatusBadRequest)
				return // exit early
			}
		}
		h.ServeHTTP(w, r) // all params present, proceed
	})
}

func PanicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

type PayLoadFrom struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func (p PayLoadFrom) New(status int, msg string) PayLoadFrom {
	return PayLoadFrom{
		StatusCode: status,
		Message:    msg,
	}
}
