package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dfang/yuanxin_api/model"
	_ "github.com/jpfuentes2/go-env/autoload"
	"github.com/robfig/cron"
)

func main() {
	a := App{}

	log.Printf("APP_DB_USER: %s, APP_DB_PASSWORD: %s, APP_DB_HOST: %s, APP_DB_NAME: %s", os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))

	a.Initialize(os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_HOST"), os.Getenv("APP_DB_NAME"))

	var Craw bool
	flag.BoolVar(&Craw, "craw", false, "craw data ....")
	flag.Parse()

	if Craw {
		fmt.Println("crawling data ....")
		items := model.CrawNews()
		for _, item := range items {
			item.InsertNewsItem(a.DB)
		}
	} else {
		fmt.Println("Server listening on 0.0.0.0:9090")
		a.Run(":9090")
	}

	// Make an enqueuer with a particular namespace
	// var enqueuer = work.NewEnqueuer("work", redisPool)
	// _, err := enqueuer.Enqueue("register_user_to_netease_im", work.Q{"user_id": 4})
	// if err != nil {
	// 	panic(err)
	// }
	// pool := work.NewWorkerPool(nil, 10, "work", redisPool)
	// log.Println("enqueue a cron job ....")
	// pool.PeriodicallyEnqueue("* * * * *", "cron_job_craw_news") // This will enqueue a "cron_job_craw_news" job every hour

	// pool.PeriodicallyEnqueue("* * * * *", "register_user_to_netease_im") // This will enqueue a "cron_job_craw_news" job every hour
	// pool.Job("register_user_to_netease_im", (*yuanxin_worker.Context).RegisterAccid) // Still need to register a handler for this job separately

	// add a job, craw news every hour
	c := cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
		// a.insertNewsItem()
	})
	c.Start()
}
