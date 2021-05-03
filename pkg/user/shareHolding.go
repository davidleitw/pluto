package user

import (
	"log"
	"math"
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

/*
	StockNumber: 股票代號
	t: 交易時間
	logs: 買賣單位, 一張=1000 logs
	price: 買入金額
	whole: 是否交易整張
*/
func (p *Position) AddStock(stockNumber int, t time.Time, logs int, price float64, whole bool) (*transaction, error) {
	n := strconv.Itoa(stockNumber)

	stockName, err := stock.GetStockNameByNumber(n)
	if err != nil {
		return nil, err
	}
	// 整張交易 單位要乘以1000
	if whole {
		logs *= 1000
	}
	cost := math.Floor(float64(logs) * price)
	fee := math.Ceil(cost * 0.001425 * p.HandlingFee)

	// 整張最低手續費需要20元
	// 零股最低手續費需要1元
	if whole && fee < 20.0 {
		fee = 20.0
	} else if !whole && fee < 1.0 {
		fee = 1.0
	}

	log.Printf("購買%s %d 單位: 成本 %d元 手續費 %d元 總價格 %d元", stockName, logs, int(cost), int(fee), int(cost+fee))

	trans := &transaction{T: t, Total: int(cost + fee), Log: logs, Price: price, Fee: int(fee)}
	if p.ShareHolding[stockNumber] {
		for _, detail := range p.Details {
			if detail.StockNumber == stockNumber {
				detail.Items = append(detail.Items, trans)
				break
			}
		}
	} else {
		p.Details = append(p.Details, &Detail{StockName: stockName, StockNumber: stockNumber, Items: []*transaction{trans}})
	}
	return trans, nil
}
