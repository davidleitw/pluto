package main

// github.com/davidleitw/pluto

import (
	"log"

	"github.com/davidleitw/pluto/pkg/stock"
)

func main() {
	s, err := stock.GetStockNameByNumber("2610")
	if err != nil {
		log.Println(err)
	}
	log.Println(s)
}
