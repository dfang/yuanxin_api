package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gocolly/colly"
	//"github.com/gocolly/colly/debug"
	//"github.com/PuerkitoBio/goquery"

	"database/sql"
	 _ "github.com/go-sql-driver/mysql"

	"fmt"
)

func (item *NewsItem) CollectBody(collector *colly.Collector) *NewsItem  {
	//&collector.
	collector.Visit(item.Link)
	collector.OnHTML(".article-box", func(e *colly.HTMLElement) {
		//log.Println(e)
		//log.Println(e.DOM.Text())
		item.Body = e.Text
	})
	return item
}

func (a *App) craw_news() {
	c := colly.NewCollector(
		// Turn on asynchronous requests
		//	colly.Async(true),
		// Attach a debugger to the collector
		//	colly.Debugger(&debug.LogDebugger{}),
	)

	//detailCollector := c.Clone()

	items := make([]NewsItem, 0, 200)

	c.OnHTML(".list-box.article-list", func(e *colly.HTMLElement) {
		e.ForEach(".newsli", func(_ int, el *colly.HTMLElement) {
			item := NewsItem{
				Title: el.ChildText(".top-title h3"),
				Type: el.ChildText(".top-title .info a"),
				UpdatedAt: el.ChildText(".top-title .add-time"),
				Source: el.ChildText(".top-title .source"),
				Description: el.ChildText("p"),
				Link: e.Request.AbsoluteURL(el.ChildAttr(".top-title h3 a", "href")),
			}

			db, err := sql.Open("mysql", "root:OC#oc2018@/News")
			checkErr(err)

			defer db.Close()

			insertStatement, _ := db.Prepare("INSERT INTO NewsItem (Title, Description, Body, Type, UpdatedAt) VALUES (?, ?, ?, ?, ?)")

			res, err := insertStatement.Exec(item.Title, item.Description, "", item.Type, item.UpdatedAt)
			checkErr(err)

			id, err := res.LastInsertId()
			checkErr(err)

			fmt.Printf("lastInsertId is %d\n", id)

			//item.CollectBody(detailCollector)
			//detailCollector.Visit(e.Request.AbsoluteURL(item.Link))
			items = append(items, item)

		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	urls := []string{
		"http://www.chinaflashmarket.com/News",
		//"http://www.chinaflashmarket.com/pricecenter/nandflash",
		//"http://www.chinaflashmarket.com/pricecenter/ddr",
		//"http://www.chinaflashmarket.com/pricecenter/lpddr",
		//"http://www.chinaflashmarket.com/pricecenter/emmc",
		//"http://www.chinaflashmarket.com/pricecenter/emcp",
		//"http://www.chinaflashmarket.com/pricecenter/ssd",
		//"http://www.chinaflashmarket.com/pricecenter/usbmodule",
		//"http://www.chinaflashmarket.com/pricecenter/usb3",
	}

	for _, url := range urls {
		c.Visit(url)
	}

	// Start scraping
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(items)
}



func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
