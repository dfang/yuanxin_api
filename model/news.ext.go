package model

import (
	"database/sql"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func GetNews(db *sql.DB, start, count int, t NewsItemType) ([]NewsItem, error) {
	statement := ""

	fmt.Println(t)
	fmt.Println(int(t))

	if t == Zero {
		statement = fmt.Sprintf("SELECT ID, Title, Description, COALESCE(Body, '') as Body, COALESCE(Image, '') as Image, COALESCE(Type, '') as Type, COALESCE(Link, '') as Link, COALESCE(Source, '') as Source, Updated_At FROM news.news_items LIMIT %d, %d", start, count)
	} else {
		statement = fmt.Sprintf("SELECT ID, Title, Description, COALESCE(Body, '') as Body, COALESCE(Image, '') as Image, COALESCE(Type, '') as Type, COALESCE(Link, '') as Link, COALESCE(Source, '') as Source, Updated_At FROM news.news_items where type = '%s' LIMIT %d, %d", t, start, count)
	}

	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	news := []NewsItem{}

	for rows.Next() {
		var item NewsItem
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Body, &item.Image, &item.Type, &item.Link, &item.Source, &item.UpdatedAt); err != nil {
			return nil, err
		}
		news = append(news, item)
	}

	return news, nil
}

func (item *NewsItem) InsertNewsItem(db *sql.DB) (sql.Result, error) {
	insertSQL := "INSERT INTO news.news_items (title, description, image, body, type, source, link, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	selectSQL := "select count(*) from news.news_items where title = ?"
	var count int
	err := db.QueryRow(selectSQL, item.Title).Scan(&count)

	switch {
	case err == sql.ErrNoRows:
		log.Fatalln("No news with that title. can insert")
		return nil, err
	case err != nil:
		log.Fatal(err)
		return nil, err
	default:
		if count == 0 {
			return db.Exec(insertSQL, item.Title, item.Description, item.Image, item.Body, item.Type, item.Source, item.Link, time.Now())
		}
		return nil, err
	}
}
