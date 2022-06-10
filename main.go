package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var ()

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("www.ebay.com"),
	)

	c.UserAgent = "Solutions Exercices"

	c.OnRequest(func(r *colly.Request) {

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}

		fmt.Println(r.Method)
	})

	fmt.Println("ARTICULOS--")
	c.OnHTML("li .s-item__wrapper", func(e *colly.HTMLElement) {
		//product := Product{}
		fmt.Println(e.ChildText(".s-item__title"))
		fmt.Println(e.ChildText(".s-item__price"))
		fmt.Println(e.ChildText(".SECONDARY_INFO"))
		fmt.Println(e.ChildText(".s-item__link"))

	})

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("-----------------------------")

		fmt.Println(r.StatusCode)

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	})

	c.Visit("https://www.ebay.com/sch/i.html?_nkw=garlandcomputer")

}
