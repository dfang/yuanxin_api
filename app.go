package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/thoas/stats"
	"github.com/urfave/negroni"
	"github.com/zbindenren/negroni-prometheus"

	. "github.com/dfang/yuanxin/endpoints"
	"github.com/dfang/yuanxin/model"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// App Main App
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

var jmw = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	},
	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	SigningMethod: jwt.SigningMethodHS256,
})

// Initialize intialize database, routes
func (a *App) Initialize(user, password, host, dbName string) {
	// why parseTime=true
	// error: "sql: Scan error on column index 7: null: cannot scan type []uint8 into null.Time: [50 48 49 56 45 48 52 45 49 52 32 49 51 58 52 56 58 48 52]"
	// https://github.com/xo/xo/issues/19
	connectionString := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, password, host, dbName)
	fmt.Printf("connectionString is %s\n", connectionString)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = a.DB.Ping()
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
	a.initializeRoutes(jmw)
}

// Run run http server
func (a App) Run(addr string) {
	// http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, a.Router))

	// n := negroni.Classic() // Includes some default middlewares
	n := negroni.New()
	l := negroni.NewLogger()
	m := negroniprometheus.NewMiddleware("serviceName")
	s := stats.New()

	n.Use(l)
	n.Use(m)
	n.Use(s)

	a.Router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stats := s.Data()
		b, _ := json.Marshal(stats)
		w.Write(b)
	})
	// recovery := negroni.NewRecovery()
	// recovery.PanicHandlerFunc = reportToSentry
	// recovery.Formatter = &negroni.HTMLPanicFormatter{}

	n.Use(negroni.HandlerFunc(Recovery))

	n.UseHandler(a.Router)
	http.ListenAndServe(":9090", n)
}

func (a *App) initializeRoutes(jwtmiddleware *jwtmiddleware.JWTMiddleware) {
	r := a.Router

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		panic(errors.New("panic"))
	})

	// prometheus metrics
	r.Handle("/metrics", prometheus.Handler())

	r.HandleFunc("/news", ListNewsItemEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/news/{id:[0-9]+}", GetNewsItemEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/users/{id:[0-9]+}", GetUserEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/profile", GetUserProfileEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/captcha/send", SendSMSEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/captcha/validate", ValidateSMSEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/exists", ExistsEndpoint(a.DB)).Methods("POST")

	r.HandleFunc("/registrations", RegistrationEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/sessions", SessionEndpoint(a.DB)).Methods("POST")
	r.Handle("/passwords", PasswordEndpoint(a.DB)).Methods("PUT")

	r.Handle("/upload", Protected(UploadEndpoint(a.DB))).Methods("POST")
	r.Handle("/registrations", Protected(UpdateRegistrationInfo(a.DB))).Methods("PUT")
	r.Handle("/suggestions", Protected(SuggestionEndpoint(a.DB))).Methods("POST")

	r.Handle("/apply/seller", Protected(ApplySellerEndpoint(a.DB))).Methods("POST")
	r.Handle("/invitations/check", Protected(CheckInvitationCodeEndpoint(a.DB))).Methods("POST")
	r.Handle("/apply/expert", Protected(ApplyExpertEndpoint(a.DB))).Methods("POST")

	r.Handle("/help_requests", Protected(PublishHelpRequestEndpoint(a.DB))).Methods("POST")
	r.Handle("/buy_requests", Protected(PublishBuyRequestEndpoint(a.DB))).Methods(http.MethodPost)

	r.Handle("/comments", Protected(PublishCommentEndpoint(a.DB))).Methods("POST")

	// 查询news/buy_requests/help_requests的所有评论
	r.HandleFunc("/comments", ListCommentsEndpoint(a.DB)).Methods("GET")

	r.Handle("/favorable", Protected(FavorableEndpoint(a.DB))).Methods("PUT")
	r.Handle("/likable", Protected(LikableEndpoint(a.DB))).Methods("PUT")

	r.HandleFunc("/favorites", ListFavoritesEndpoint(a.DB)).Methods("GET")

	r.Handle("/chips", Protected(PublishChipEndpoint(a.DB))).Methods("POST")
	r.HandleFunc("/chips", ListChipsEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/help_requests", ListHelpRequestEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/buy_requests", ListBuyRequestEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/chips/{id:[0-9]+}", GetChipEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/help_requests/{id:[0-9]+}", GetHelpRequestEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/buy_requests/{id:[0-9]+}", GetBuyRequestEndpoint(a.DB)).Methods("GET")

	// helpers
	r.HandleFunc("/users", ListUsersEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/invitations", ListInvitationsEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/news/{id:[0-9]+}/comments", ListNewsCommentsEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/buy_requests/{id:[0-9]+}/comments", ListBuyRequestCommentEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/help_requests/{id:[0-9]+}/comments", ListHelpRequestCommentEndpoint(a.DB)).Methods("GET")

	// deprecated in favor of /favorable
	// r.HandleFunc("/favorites", PubishFavoriteEndpoint(a.DB)).Methods("POST")
	// r.HandleFunc("/favorites/{id:[0-9]+}", DestroyFavoriteEndpoint(a.DB)).Methods("DELETE")

}

// Protected  Just a authentication wrapper
func Protected(h http.HandlerFunc) http.Handler {
	return negroni.New(
		negroni.HandlerFunc(jmw.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(h)),
	)
}
