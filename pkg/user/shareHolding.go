package user

import (
	"log"
	"strconv"
	"time"

	"github.com/davidleitw/pluto/pkg/stock"
)

// 紀錄持股相關的資料結構

type Position struct {
	HandlingFee  float64 // 手續費折數
	ShareHolding map[int]bool
	Details      []*Detail
}

type Detail struct {
	StockName   string // 股票名稱
	StockNumber int    // 股票編號
	Items       []*transaction
}

type transaction struct {
	T     time.Time // 買入時間
	Total int       // 總價格
	Log   int       // 股數
	Price float64   // 買入價格, 成本價
	Fee   int       // 手續費
}

func NewPosition(hf float64) *Position {
	return &Position{
		HandlingFee:  hf,
		ShareHolding: make(map[int]bool),
		Details:      make([]*Detail, 0),
	}
}

func NewDetail(name string, number int) *Detail {
	return &Detail{
		StockName:   name,
		StockNumber: number,
		Items:       make([]*transaction, 0),
	}
}

func (p *Position) AddStock(StockNumber int, t time.Time, logs int, price float64) error {
	n := strconv.Itoa(StockNumber)

	stockName, err := stock.GetStockNameByNumber(n)
	if err != nil {
		return err
	}

	cost := float64(logs) * price
	// fmt.Println(stockName, cost)
	log.Println(stockName, cost)
	return nil
}
