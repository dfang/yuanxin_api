package endpoints

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/util"
)

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

type Adapter func(http.Handler) http.Handler

type MissParameterError struct {
	Parameter string
	Error     error
}

// func (p *MissParameterError) String() string {
// 	return fmt.Sprintf("%s 参数缺失", p.Parameter)
// }

type ParseError struct {
	Parameter string
	Error     error
}

// func (e *ParseError) String() string {
// 	return fmt.Sprintf("error parsing %q as int", e.Parameter)
// }
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

func RecoverEndpoint(w http.ResponseWriter) {
	err := recover()
	if err != nil {
		switch v := err.(type) {
		case *MissParameterError:
			util.RespondWithError(w, http.StatusBadRequest, 400, v.Error.Error())
		case *ParseError:
			util.RespondWithError(w, http.StatusBadRequest, 400, v.Error.Error())
		case *RecordNotFound:
			util.RespondWithError(w, http.StatusBadRequest, 400, v.Error.Error())
		case string:
			util.RespondWithError(w, http.StatusBadRequest, 400, v)
		case error:
			util.RespondWithError(w, http.StatusBadRequest, 400, v.Error())
		default:
			util.RespondWithError(w, http.StatusBadRequest, 400, "Internal Error")
		}
	}
	// if err != nil {
	// 	fmt.Println("Internal error:", err)
	// 	var err_str string
	// 	var ok bool
	// 	if err_str, ok = err.(string); !ok {
	// 		err_str = "We encountered an internal error"
	// 	}
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err_str))
	// }
}

func MyHandlerFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		h.ServeHTTP(w, r)
	})
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

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		defer log.Println("After")
		h.ServeHTTP(w, r)
	})
}

func AutoRecover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err != nil {
			fmt.Println("Internal error:", err)
			var err_str string
			var ok bool
			if err_str, ok = err.(string); !ok {
				err_str = "We encountered an internal error"
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err_str))
			return
		}
		h.ServeHTTP(w, r)
	})
}
