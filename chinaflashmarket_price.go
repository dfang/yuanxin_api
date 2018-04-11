package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func (a *App) craw_price() {
	c := colly.NewCollector()

	products := make([]Product, 0, 200)

	c.OnHTML(".price-box .table-main", func(e *colly.HTMLElement) {
		e.ForEach("table > tbody > tr", func(_ int, el *colly.HTMLElement) {
			p := Product{
				Name: el.ChildText("td:nth-child(1)"),
				Closing_price_yesterday: el.ChildText("td:nth-child(2)"),
				Current_price: el.ChildText("td:nth-child(3)"),
				Lowest_price_in_a_day: el.ChildText("td:nth-child(4)"),
				Highest_price_in_a_day: el.ChildText("td:nth-child(5)"),
				Daily_change: el.ChildText("td:nth-child(6)"),
				Price_in_hk: el.ChildText("td:nth-child(7)"),
				Lowest_price_in_a_week: el.ChildText("td:nth-child(7)"),
				Highest_price_in_a_week: el.ChildText("td:nth-child(8)"),
			}
			products = append(products, p)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	urls := []string{
		"http://www.chinaflashmarket.com/pricecenter/flashcard",
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
	enc.Encode(products)
}
