package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"database/sql"

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
	connectionString := fmt.Sprintf("%s:%s@%s/%s", user, password, host, dbName)
	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	// defer a.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/news", a.getNews).Methods("GET")
	a.Router.HandleFunc("/news/{id:[0-9]+}", a.getNewsItem).Methods("GET")
	a.Router.HandleFunc("/sms_captcha", a.getSmsCaptcha).Methods("POST")
	//a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	//a.Router.HandleFunc("/user/{id:[0-9]+}", a.updateUser).Methods("PUT")
	//a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getNews(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	t, _ := strconv.Atoi(r.FormValue("type"))

	if count < 1 {
		count = 9
	}

	if start < 0 {
		start = 0
	}

	news, err := getNews(a.DB, start, count, NewsItemType(t))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, news)
}

func (a *App) getNewsItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := NewsItem{ID: id}
	if err := u.getNewsItem(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) insertNewsItem() {
	news := a.CrawNews()

	for _, item := range news {
		//fmt.Printf("%v", item)
		//item.InsertNewsItem(a.DB)
		result, err := item.InsertNewsItem(a.DB)
		fmt.Println(result)
		if err != nil {
			fmt.Println(err)
		}
		if result != nil {
			fmt.Println(result.RowsAffected())
		}
	}
}

func (a *App) getSmsCaptcha(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")

	data := SMSData{
		Phone: phone,
	}

	result, err := SendSms(&data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, *result)
}
