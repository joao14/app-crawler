package main

import (
	"app-crawler/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	c.OnHTML("li .s-item__wrapper", func(e *colly.HTMLElement) {
		product := model.Product{}
		product.Title = e.ChildText(".s-item__title")
		product.Price = e.ChildText(".s-item__price")
		product.Product_url = e.ChildText(".SECONDARY_INFO")
		product.Condition = e.ChildAttr(".s-item__link", "href")

		url := strings.Split(e.ChildAttr(".s-item__link", "href"), "/itm/")

		f, err := os.Create("data/" + strings.Split(url[1], "?")[0] + ".json")

		if err != nil {
			panic(err)
		}

		f.Close()

		file, _ := json.MarshalIndent(product, "", "")

		_ = ioutil.WriteFile("data/"+strings.Split(url[1], "?")[0]+".json", file, 0644)

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
