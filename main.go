package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dfang/yuanxin/model"
	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/robfig/cron"
)

var a App

func main() {
	a = App{}

	log.Printf("%s:%s@%s/%s\n", os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))

	// You need to set your Username and Password here
	a.Initialize(os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))

	var flagVal bool
	flag.BoolVar(&flagVal, "craw", false, "craw data ....")
	flag.Parse()

	if flagVal {
		fmt.Println("crawling data ....")
		items := model.CrawNews()
		fmt.Printf("%s", items[0].Body.String)

		for _, item := range items {
			item.InsertNewsItem(a.DB)
		}
	} else {
		fmt.Println("Server listening on 0.0.0.0:9090")
		a.Run(":9090")
	}

	// add a job, craw news every hour
	c := cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
		// a.insertNewsItem()
	})
	c.Start()
}
