package main

import (
	"fmt"
)

func main() {

	var publisher, writer, artist, title string
	var year, pageNumber int16

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14

	fmt.Println(title, "written by", writer, "drawn by", artist, ".", "Published by", publisher, "in", year, "with", pageNumber, "pages.")

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160

	fmt.Println(title, "written by", writer, "drawn by", artist, ".", "Published by", publisher, "in", year, "with", pageNumber, "pages.")
}
