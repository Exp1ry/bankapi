package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

func CrawlerUAE(pages int, categories []string) ([]FurnitureStore, error) {
	furnitureSlice := []FurnitureStore{}
	c := colly.NewCollector(
		colly.AllowedDomains("yellowpages-uae.com", "www.yellowpages-uae.com", "https://www.yellowpages-uae.com"),
	)

	c.OnHTML("div.right-section", func(r *colly.HTMLElement) {

		dom := r.DOM

		name := dom.Find("h2[itemprop]").Text()
		types := dom.Parent().Parent().Parent().Find("div.result-title").Find("h1").Text()
		location := dom.Find("span.location").Find("span[itemprop]").Text()
		city := dom.Find("span.locationCity").Find("strong[itemprop]").Text()
		phone := dom.Find("span.phonespn").Find("span.phone").Text()
		productAndService := dom.Find("div.categories").Find("a[href]").Text()

		furniture := FurnitureStore{
			Name:                name,
			Type:                types,
			Location:            location,
			City:                city,
			Phone:               phone,
			ProductsAndServices: productAndService,
		}
		fmt.Println(types)
		furnitureSlice = append(furnitureSlice, furniture)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(e *colly.Response, err error) {
		log.Println("Something wrong", err)
	})

	for i := 0; i < pages; i++ {
		for _, v := range categories {

			if i == 0 {
				c.Visit("https://www.yellowpages-uae.com/uae/" + v + "")
			}
			c.Visit("https://www.yellowpages-uae.com/uae/" + v + "-" + strconv.Itoa(i) + ".html")
		}

	}

	return furnitureSlice, nil
}
