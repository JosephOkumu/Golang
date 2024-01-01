// This program is for maintaining the catalog of items in a book store.

package main

import "fmt"

var publisher, writer, artist, title string
var year, pageNumber int
var grade float32
var cost float32

func Cost() {
	if pageNumber <= 50 && year < 2000 {
		cost = 50
	}
	if pageNumber > 50 && year > 2000 {
		cost = 100
	}
}

func main() {

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	Cost()

	fmt.Println(title, "written by", writer, "drawn by", artist, "rate", grade, "published by", publisher, "year of publishing", year, "page number", pageNumber, "price", cost)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	Cost()

	fmt.Println(title, "written by", writer, "drawn by", artist, "rate", grade, "published by", publisher, "year of publishing", year, "page number", pageNumber, "price", cost)
}
