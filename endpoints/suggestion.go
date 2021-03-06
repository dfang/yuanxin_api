package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
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

// SuggestionEndpoint 投诉建议
func SuggestionEndpoint(db *sql.DB) http.HandlerFunc {
	// return appHandler
	// return Chain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// }), Logging(), Auth())
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "content")

		userID := GetUIDFromContext(r)
		suggestion := model.Suggestion{
			UserID:  null.IntFrom(int64(userID)),
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
