package main

import (
	"codewars/bookseller"
	"log"
)
import "codewars/roman"

func main()  {
	// Roman Decoder
	log.Print(roman.Decode("MDCLXVI"))
	// StockList
	var b = []string{} //"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"
	var c = []string{"A", "B", "C", "D"}
	log.Print(bookseller.StockList(b,c))
}
