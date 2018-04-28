package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	null "gopkg.in/guregu/null.v3"
)

type appHandler func(http.ResponseWriter, *http.Request)

// type HandlerFunc func(ResponseWriter, *Request)

func (f appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type Middleware func(http.Handler) http.Handler

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {
	// Create a new Middleware
	return func(f http.Handler) http.Handler {
		// Define the http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do middleware things
			start := time.Now()
			defer func() { fmt.Println("timing: ", r.URL.Path, time.Since(start)) }()
			// Call the next middleware/handler in chain
			f.ServeHTTP(w, r)
		})
	}
}

func Auth() Middleware {
	return func(f http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// do authentication

			fmt.Println("authenticating ............")

			f.ServeHTTP(w, r)
		})
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func SuggestionEndpoint(db *sql.DB) http.Handler {
	// return appHandler
	// return Chain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// }), Logging(), Auth())

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "user_id", "content")

		user_id := ParseParameterToInt(r, "user_id")
		suggestion := model.Suggestion{
			UserID:  null.IntFrom(int64(user_id)),
			Content: null.StringFrom(r.PostFormValue("content")),
		}

		err := suggestion.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, "提交失败"})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "提交成功"})
	})
}
