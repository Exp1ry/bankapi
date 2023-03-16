package main

import (
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()

	res, err := CrawlerUAE(10, []string{"furniture", "landscaping", "building"})
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res {
		if err := store.CreateCompany(&v); err != nil {
			log.Fatal(err)
		}
	}
}
