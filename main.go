package main

import (
	"fmt"

	"log"
	"os"

	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/robfig/cron"
)

func main() {
	a := App{}

	log.Printf("%s:%s@%s/%s\n", os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))

	// You need to set your Username and Password here
	a.Initialize(os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))

	//craw news for the first time
	a.insertNewsItem()

	// add a job, craw news every hour
	c := cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
		a.insertNewsItem()
	})
	c.Start()

	a.Run(":9090")
}
