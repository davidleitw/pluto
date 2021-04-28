package main

import (
	"log"

	"github.com/davidleitw/pluto/pkg/stock"
)

func main() {
	// stock.GetStockInformation("tse_2610.tw")
	info := stock.ShareHoldingQuery("tse_2610.tw")
	log.Println(info)

	a := stock.GenerateQueryString("2002", "2603", "1314", "3481")
	log.Println(a)
	info = stock.ShareHoldingQuery(a)
	log.Println(info)
}
