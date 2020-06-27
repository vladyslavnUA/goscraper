package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gocolly/colly"
)


// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/

type bookItem struct {
	Title string `json:"title"`
	Rating string `json:"rating"`
	Price string `json:"price"`
	Stock string `json:"stock"`
}

func dataOut(file []byte) {
	this := ioutil.WriteFile("data.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}

func serializeJSON(boo []bookItem) {
	fmt.Println("Serializing Data")
	allbooksjstoned, _ := json.Marshal(boo)
	dataOut(allbooksjstoned)
	fmt.Println("Serializing Complete")
	fmt.Println(string(allbooksjstoned))
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
		book.Stock = e.ChildText(".availability")
		books = append(books, book)

		// Print link
		fmt.Printf("Title: %q\n, Rating: %q\n, Price: %q\n, Stock: %q\n,", book.Title, book.Rating, book.Price, book.Stock)
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
	serializeJSON(books)
}