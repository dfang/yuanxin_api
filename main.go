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
		items := model.CrawNews()

		fmt.Println("vvvvvv")
		fmt.Println(items)

		for _, item := range items {
			item.InsertNewsItem(a.DB)
		}
		fmt.Println("crawling data ....")
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
