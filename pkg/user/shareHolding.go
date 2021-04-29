package user

import "time"

type Details struct {
	StockName   string // 股票名稱
	StockNumber int    // 股票編號
	Items       []*transaction
}

type transaction struct {
	t     time.Time
	log   int
	price float64
}

func NewDetail(name string, number int) *Details {
	return &Details{
		StockName:   name,
		StockNumber: number,
		Items:       make([]*transaction, 0),
	}
}
