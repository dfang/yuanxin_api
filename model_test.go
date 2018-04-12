package main

import (
	"testing"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"database/sql"
)

func TestNewsItem_InsertNewsItem(t *testing.T) {

	item := NewsItem{
		Title:       "东芝存储芯片出售难产 债权银行施压其尽快完成交易",
		Link:        "http://www.chinaflashmarket.com/Producer/TOSHIBA/News/160860",
		Type:        "厂商动态",
		Description: "据《金融时报》北京时间4月9日报道，银行业人士消息称，尽管有维权股东认为东芝存储芯片业务实际价值是出售价的两倍以上，但是东芝主要债权银行正督促东芝推进这笔2万亿日元(约合187亿美元)的出售交易。",
		Body:        "",
		UpdatedAt:   "2018/4/9 11:30:55",
		Source:      "来源：凤凰科技",
	}

	username := "root"
	password := "OC#oc2018"
	host := "tcp(localhost:3306)"
	dbName := "news"
	connectionString := fmt.Sprintf("%s:%s@%s/%s", username, password, host, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//sqlStr := "INSERT INTO news_item (title, description, body, type, updated_at) VALUES (?, ?, ?, ?, ?)"
	//insStmt, _ := db.Prepare(sqlStr)
	//_, err = insStmt.Exec(item.Title, item.Description, item.Body, item.Type, time.Now())
	//if err != nil {
	//	log.Fatal(err)
	//}

	item.InsertNewsItem(db)
}