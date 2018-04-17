package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/dfang/yuanxin/endpoints"
	"github.com/dfang/yuanxin/model"
	_ "github.com/go-sql-driver/mysql"

	"os"

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

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a App) Run(addr string) {
	http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/news", endpoints.ListNewsItemEndpoint(a.DB)).Methods("GET")
	a.Router.HandleFunc("/news/{id:[0-9]+}", endpoints.GetNewsItemEndpoint(a.DB)).Methods("GET")

	a.Router.HandleFunc("/captcha/send", endpoints.SendSMSEndpoint).Methods("POST")
	a.Router.HandleFunc("/captcha/validate", endpoints.ValidateSMSEndpoint).Methods("POST")

	a.Router.HandleFunc("/registrations", endpoints.RegistrationEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/sessions", endpoints.SessionEndpoint(a.DB)).Methods("POST")
	a.Router.HandleFunc("/passwords", endpoints.PasswordEndpoint).Methods("PUT")

	a.Router.HandleFunc("/exists", endpoints.ExistsEndpoint(a.DB)).Methods("POST")
}
