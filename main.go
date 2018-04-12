package main

import (
	"github.com/robfig/cron"
	"fmt"
)

func main() {
	a := App{}

	// You need to set your Username and Password here
	a.Initialize("root", "OC#oc2018", "tcp(127.0.0.1:3306)", "news")

	a.insertNewsItem()

	c := cron.New()
	c.AddFunc("@hourly",      func() {
		fmt.Println("Every hour")
		a.insertNewsItem()
	})

	c.Start()


	a.Run(":9090")
}
