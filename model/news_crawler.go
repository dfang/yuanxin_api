package model

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/guregu/null.v3"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	// "github.com/metakeule/fmtdate"
)

func (item *NewsItem) CollectBody(collector *colly.Collector) *NewsItem {
	//&collector.
	collector.Visit(item.Link.String)
	collector.OnHTML(".article-box", func(e *colly.HTMLElement) {
		//log.Println(e)
		//log.Println(e.DOM.Text())
		item.Body.String = e.Text
	})
	return item
}

func CrawNews() []NewsItem {
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
			// t, err := time.Parse("2018/4/17 16:45:50", el.ChildText(".top-title .add-time"))
			// s := strings.Replace(el.ChildText(".top-title .add-time"), "/", "-", -1)
			// t, err := fmtdate.Parse("YYYY-MM-DD hh:mm:ss", s)

			// null.TimeFrom()
			// fmt.Println(t)

			// checkErr(err)

			item := NewsItem{
				Title: null.StringFrom(el.ChildText(".top-title h3")),
				Type:  null.StringFrom(el.ChildText(".top-title .info a")),
				// UpdatedAt:   null.TimeFrom(t),
				Source:      null.StringFrom(el.ChildText(".top-title .source")),
				Description: null.StringFrom(el.ChildText("p")),
				Link:        null.StringFrom(e.Request.AbsoluteURL(el.ChildAttr(".top-title h3 a", "href"))),
			}

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
		"http://www.chinaflashmarket.com/News/Page-2",
		"http://www.chinaflashmarket.com/News/Page-3",
		"http://www.chinaflashmarket.com/News/Page-4",
		"http://www.chinaflashmarket.com/News/Page-5",

		"http://www.chinaflashmarket.com/Industry",
		"http://www.chinaflashmarket.com/Industry/Page-2",
		"http://www.chinaflashmarket.com/Industry/Page-3",
		"http://www.chinaflashmarket.com/Industry/Page-4",
		"http://www.chinaflashmarket.com/Industry/Page-5",
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

	return items
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
