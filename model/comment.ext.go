package model

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func GetComments(db *sql.DB, start, count int, commentable_type string, commentable_id int) ([]Comment, error) {
	statement := fmt.Sprintf("SELECT * FROM news.comments where commentable_type = '%s' AND commentable_id = '%d' LIMIT %d, %d", commentable_type, commentable_id, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []Comment{}

	for rows.Next() {
		var item Comment
		if err := rows.Scan(&item.ID, &item.UserID, &item.CommentableType, &item.CommentableID, &item.Content, &item.IsPicked, &item.Likes, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
