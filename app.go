package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/dfang/yuanxin/endpoints"
	_ "github.com/go-sql-driver/mysql"

	"os"

	// . "github.com/dfang/yuanxin/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, host, dbName string) {
	connectionString := fmt.Sprintf("%s:%s@%s/%s", user, password, host, dbName)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
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

	a.Router.HandleFunc("/users", endpoints.RegisterUserEndpoint(a.DB)).Methods("POST")

	a.Router.HandleFunc("/captcha/send", endpoints.SendSMSEndpoint).Methods("POST")
	a.Router.HandleFunc("/captcha/validate", endpoints.ValidateSMSEndpoint).Methods("POST")

	a.Router.HandleFunc("/registrations", endpoints.RegistrationEndpoint).Methods("POST")
	a.Router.HandleFunc("/sessions", endpoints.SessionEndpoint).Methods("POST")
	a.Router.HandleFunc("/passwords", endpoints.PasswordEndpoint).Methods("PUT")

	a.Router.HandleFunc("/exists", endpoints.ValidateSMSEndpoint).Methods("POST")

	//a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	//a.Router.HandleFunc("/user/{id:[0-9]+}", a.updateUser).Methods("PUT")
	//a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}
