package main

import (
	"fmt"

	"github.com/jpfuentes2/go-env"
	"github.com/robfig/cron"
)

func main() {
	env.ReadEnv(".env")

	a := App{}

	// You need to set your Username and Password here
	a.Initialize("root", "OC#oc2018", "tcp(db:3306)", "news")

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
