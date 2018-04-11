package main

import (
	"database/sql"
	"fmt"
)

type Product struct {
	Name        			string
	Closing_price_yesterday string
	Current_price     		string
	Lowest_price_in_a_day   string
	Highest_price_in_a_day  string
	Daily_change    		string
	Price_in_hk  			string
	Lowest_price_in_a_week   string
	Highest_price_in_a_week      string
}

type NewsItem struct {
	ID  					int  	`json:"id"`
	Title        			string  `json:"title"`
	Link					string	`json:"link"`
	Type					string  `json:"type"`
	Description 			string  `json:"description"`
	Body 					string  `json:"body"`
	UpdatedAt  				string  `json:"updated_at"`
	Source     				string  `json:"source"`
}

func getNews(db *sql.DB, start, count int) ([]NewsItem, error) {
	statement := fmt.Sprintf("SELECT ID, Title, Description, COALESCE(Type, '') as Type, COALESCE(Link, '') as Link, COALESCE(Source, '') as Source, UpdatedAt FROM NewsItem")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	news := []NewsItem{}

	for rows.Next() {
		var u NewsItem
		if err := rows.Scan(&u.ID, &u.Title, &u.Description, &u.Type, &u.Link, &u.Source, &u.UpdatedAt); err != nil {
			return nil, err
		}
		news = append(news, u)
	}

	return news, nil

}

func (u *NewsItem) getNewsItem(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT Title, Link, Type, Description, Body, UpdatedAt, Source FROM News WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Title, &u.Link, &u.Type, &u.Description, &u.Body, &u.UpdatedAt)
}

func (item *NewsItem) insertNewsItem(db *sql.DB) (sql.Result, error)  {
	insertStatement, _ := db.Prepare("INSERT INTO NewsItem (Title, Description, Body, Type, UpdatedAt) VALUES (?, ?, ?, ?, ?)")

	//res, err := insertStatement.Exec(item.Title, item.Description, "", item.Type, item.UpdatedAt)
	//checkErr(err)
	//
	//id, err := res.LastInsertId()
	//checkErr(err)
	//
	//fmt.Printf("lastInsertId is %d\n", id)
	return insertStatement.Exec(item.Title, item.Description, "", item.Type, item.UpdatedAt)
}