package main

import (
	"fmt"
)

func main() {

	var publisher, writer, artist, title string
	var year, pageNumber int16

	title = "The Count of Mounte Cristo"
	writer = "Alexandre Dumas"
	artist = "Auguste Maquet"
	publisher = "Wordsworth Editions"
	year = 1846
	pageNumber = 928

	fmt.Println(title, "written by", writer, "drawn by", artist, ".", "Published by", publisher, "in", year, "with", pageNumber, "pages.")

	title = "Les Misérables"
	writer = "Victor Hugo"
	artist = "Emile Bayard"
	publisher = "	A. Lacroix, Verboeckhoven & Cie."
	year = 1862
	pageNumber = 1264

	fmt.Println(title, "written by", writer, "drawn by", artist, ".", "Published by", publisher, "in", year, "with", pageNumber, "pages.")
}
