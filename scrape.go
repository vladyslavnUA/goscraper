package main

import (
	"fmt"

	"github.com/gocolly/colly"
)


// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/

type bookItem struct {
	Title string `json:"title"`
	Rating string `json:"rating"`
	Price string `json:"price"`
	Image string `json:"image"`
	Stock string `json:"stock"`
}



func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	books := []bookItem{}
	// On every a element which has href attribute call callback
	c.OnHTML(".row ol > li", func(e *colly.HTMLElement) {

		// link := e.Attr("href")
		book := bookItem{}
		book.Title = e.ChildText("h3")
		book.Rating = e.ChildText(".star-rating")
		book.Price = e.ChildText(".price_color")
		book.Image = e.ChildText(".thumbnail")
		book.Stock = e.ChildText(".availability")
		books = append(books, book)

		// Print link
		fmt.Printf("Title: %q\n, Rating: %q\n, Price: %q\n, Image: %q\n, Stock: %q\n,", book.Title, book.Rating, book.Price, book.Image, book.Stock)
                // fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// c.OnHTML("ol > li", func(e *colly.HTMLElement) {
	// 	link := e.Attr("a")
	// 	fmt.Printf("title: %t", e.Text, link)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://books.toscrape.com/")
}