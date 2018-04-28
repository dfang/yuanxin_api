package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/thoas/stats"
	"github.com/urfave/negroni"
	"github.com/zbindenren/negroni-prometheus"

	"github.com/dfang/yuanxin/endpoints"
	"github.com/dfang/yuanxin/model"
	_ "github.com/go-sql-driver/mysql"

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

	jmiddle := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	a.initializeRoutes(jmiddle)
}

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	fmt.Fprintf(w, "This is an authenticated request")
	fmt.Fprintf(w, "Claim content:\n")
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		fmt.Fprintf(w, "%s :\t%#v\n", k, v)
	}
})

func (a App) Run(addr string) {
	// http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, a.Router))

	// n := negroni.Classic() // Includes some default middlewares
	n := negroni.New()
	// n := negroni.New(negroni.HandlerFunc(
	// 	jwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(myHandler),
	// )

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

	n.Use(negroni.HandlerFunc(endpoints.Recovery))

	n.UseHandler(a.Router)
	http.ListenAndServe(":9090", n)
}

// func reportToSentry(info *negroni.PanicInformation) {
// }

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		start := time.Now()
		defer func() { fmt.Println("timing: ", r.URL.Path, time.Since(start)) }()
		next.ServeHTTP(w, r)
	})
}

func (a *App) initializeRoutes(jwtmiddleware *jwtmiddleware.JWTMiddleware) {
	r := a.Router

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		panic(errors.New("panic"))
	})

	r.Handle("/metrics", prometheus.Handler())

	r.HandleFunc("/news", endpoints.ListNewsItemEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/news/{id:[0-9]+}", endpoints.GetNewsItemEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/users/{id:[0-9]+}", endpoints.GetUserEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/users", endpoints.ListUsersEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/captcha/send", endpoints.SendSMSEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/captcha/validate", endpoints.ValidateSMSEndpoint(a.DB)).Methods("POST")

	r.HandleFunc("/registrations", endpoints.RegistrationEndpoint(a.DB)).Methods("POST")

	r.HandleFunc("/sessions", endpoints.SessionEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/passwords", endpoints.PasswordEndpoint(a.DB)).Methods("PUT")

	r.HandleFunc("/exists", endpoints.ExistsEndpoint(a.DB)).Methods("POST")

	r.Handle("/upload",
		negroni.New(
			negroni.HandlerFunc(jwtmiddleware.HandlerWithNext),
			negroni.Wrap(http.HandlerFunc(endpoints.UploadEndpoint(a.DB))),
		)).Methods("POST")

	r.HandleFunc("/registrations", endpoints.UpdateRegistrationInfo(a.DB)).Methods("PUT")

	// r.Handle("/registrations", endpoints.Logging(endpoints.RegistrationHandler)).Methods("PUT")

	// r.HandleFunc("/suggestions", endpoints.SuggestionEndpoint(a.DB)).Methods("POST")
	r.Handle("/suggestions", loggingMiddleware(endpoints.SuggestionEndpoint(a.DB))).Methods("POST")

	r.Handle("/apply/seller",
		negroni.New(
			negroni.HandlerFunc(jwtmiddleware.HandlerWithNext),
			negroni.Wrap(http.HandlerFunc(endpoints.ApplySellerEndpoint(a.DB))),
		)).Methods("POST")

	r.Handle("/apply/expert",
		negroni.New(
			negroni.HandlerFunc(jwtmiddleware.HandlerWithNext),
			negroni.Wrap(http.HandlerFunc(endpoints.ApplyExpertEndpoint(a.DB))),
		)).Methods("POST")

	r.HandleFunc("/invitations/check", endpoints.CheckInvitationCodeEndpoint(a.DB)).Methods("POST")

	r.Handle("/help_requests",
		negroni.New(
			negroni.HandlerFunc(jwtmiddleware.HandlerWithNext),
			negroni.Wrap(http.HandlerFunc(endpoints.PublishHelpRequestEndpoint(a.DB))),
		)).Methods("POST")

	r.Handle("/buy_requests",
		negroni.New(
			negroni.HandlerFunc(jwtmiddleware.HandlerWithNext),
			negroni.Wrap(http.HandlerFunc(endpoints.PublishBuyRequestEndpoint(a.DB))),
		)).Methods("POST")
	// r.HandleFunc("/help_requests", endpoints.PublishHelpRequestEndpoint(a.DB)).Methods("POST")
	// r.HandleFunc("/buy_requests", endpoints.PublishBuyRequestEndpoint(a.DB)).Methods("POST")

	r.HandleFunc("/comments", endpoints.PublishCommentEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/news/{id:[0-9]+}/comments", endpoints.ListNewsCommentsEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/buy_requests/{id:[0-9]+}/comments", endpoints.ListBuyRequestCommentEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/help_requests/{id:[0-9]+}/comments", endpoints.ListHelpRequestCommentEndpoint(a.DB)).Methods("GET")
	// 查询news/buy_requests/help_requests的所有评论
	r.HandleFunc("/comments", endpoints.ListCommentsEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/favorites", endpoints.ListFavoritesEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/favorites", endpoints.PubishFavoriteEndpoint(a.DB)).Methods("POST")
	r.HandleFunc("/favorites/{id:[0-9]+}", endpoints.DestroyFavoriteEndpoint(a.DB)).Methods("DELETE")

	r.HandleFunc("/chips", endpoints.PublishChipEndpoint(a.DB)).Methods("POST")

	r.HandleFunc("/chips", endpoints.ListChipsEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/help_requests", endpoints.ListHelpRequestEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/buy_requests", endpoints.ListBuyRequestEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/chips/{id:[0-9]+}", endpoints.GetChipEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/help_requests/{id:[0-9]+}", endpoints.GetHelpRequestEndpoint(a.DB)).Methods("GET")
	r.HandleFunc("/buy_requests/{id:[0-9]+}", endpoints.GetBuyRequestEndpoint(a.DB)).Methods("GET")

	r.HandleFunc("/invitations", endpoints.ListInvitationsEndpoint(a.DB)).Methods("GET")
}
