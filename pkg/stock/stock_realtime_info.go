package stock

import (
	"time"
)

// twse api json 格式參考自 https://github.com/Asoul/tsrtc

type unixMapData map[int64]StockData

type StockSummary struct {
	StockNumber string
	Date        time.Time
	UnixMapData unixMapData // 時間資料暫存
	StockKind   string      // TSE, OTC
}

type StockInformation struct {
	StockKind       string // TSE, OTC
	FullCompanyName string
	StockName       string
	StockNumber     string
	StockTicker     string
	StockCategory   string
}

type StockData struct {
	BestAskPrice   []float64              // 最佳五檔賣出價資訊
	BestBidPrice   []float64              // 最佳五檔買進價資訊
	BestAskVolume  []int64                // 最佳五檔賣出量資訊
	BestBidVolume  []int64                // 最佳五檔買進量資訊
	OpenedPrice    float64                // 開盤價格
	HeighestPrice  float64                // 最高價
	LowestPrice    float64                // 最低價
	NowPrice       float64                // 該盤成交價格
	LimitUp        float64                // 漲停價
	LimitDown      float64                // 跌停價
	Volume         float64                // 該盤成交量
	VolumeAcc      float64                // 累計成交量
	YesterdayPrice float64                // 昨日收盤價格
	TradeTime      time.Time              // 交易時間
	Info           StockInformation       // 相關資訊
	SysInfo        map[string]interface{} // 系統回傳資訊
}
