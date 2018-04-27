package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/dfang/yuanxin/endpoints"
	"github.com/dfang/yuanxin/model"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, host, dbName string) {
	// why parseTime=true
	// error: "sql: Scan error on column index 7: null: cannot scan type []uint8 into null.Time: [50 48 49 56 45 48 52 45 49 52 32 49 51 58 52 56 58 48 52]"
	// https://github.com/xo/xo/issues/19
	connectionString := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, password, host, dbName)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	model.XOLog = func(s string, p ...interface{}) {
		fmt.Printf("> SQL: %s -- params: %v\n", s, p)
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a App) Run(addr string) {
	http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/news", endpoints.ListNewsItemEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/news/{id:[0-9]+}", endpoints.GetNewsItemEndpoint(a.DB)).Methods("GET")

	a.Router.HandleFunc("/users/{id:[0-9]+}", endpoints.GetUserEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/users", endpoints.ListUsersEndpoint(a.DB)).Methods("GET")

	a.Router.HandleFunc("/captcha/send", endpoints.SendSMSEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/captcha/validate", endpoints.ValidateSMSEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/registrations", endpoints.RegistrationEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/sessions", endpoints.SessionEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/passwords", endpoints.PasswordEndpoint(a.DB)).Methods("PUT")

	a.Router.HandleFunc("/exists", endpoints.ExistsEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/upload", endpoints.UploadEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/registrations", endpoints.UpdateRegistrationInfo(a.DB)).Methods("PUT")

	// a.Router.Handle("/registrations", endpoints.Logging(endpoints.RegistrationHandler)).Methods("PUT")

	a.Router.HandleFunc("/suggestions", endpoints.SuggestionEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/apply/seller", endpoints.ApplySellerEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/apply/expert", endpoints.ApplyExpertEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/invitations/check", endpoints.CheckInvitationCodeEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/help_requests", endpoints.PublishHelpRequestEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/buy_requests", endpoints.PublishBuyRequestEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/comments", endpoints.PublishCommentEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/news/{id:[0-9]+}/comments", endpoints.ListNewsCommentsEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/buy_requests/{id:[0-9]+}/comments", endpoints.ListBuyRequestCommentEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/help_requests/{id:[0-9]+}/comments", endpoints.ListHelpRequestCommentEndpoint(a.DB)).Methods("GET")
	// 查询news/buy_requests/help_requests的所有评论
	a.Router.HandleFunc("/comments", endpoints.ListCommentEndpoint(a.DB)).Methods("GET")

	a.Router.HandleFunc("/favorites", endpoints.ListFavoritesEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/favorites", endpoints.PubishFavoriteEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/favorites/{id:[0-9]+}", endpoints.DestroyFavoriteEndpoint(a.DB)).Methods("DELETE")

	a.Router.HandleFunc("/chips", endpoints.PublishChipEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/chips", endpoints.ListChipsEndpoint(a.DB)).Methods("GET")

	a.Router.HandleFunc("/help_requests", endpoints.ListHelpRequestEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/buy_requests", endpoints.ListBuyRequestEndpoint(a.DB)).Methods("GET")

	a.Router.HandleFunc("/invitations", endpoints.ListInvitationsEndpoint(a.DB)).Methods("GET")

}
