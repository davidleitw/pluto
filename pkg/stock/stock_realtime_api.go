package stock

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var REALTIME_STOCK_INFORMATION_URL string = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=%s&json=1&delay=0"

type StockApiJson struct {
	Msgarray []struct {
		Tv string `json:"tv"` // Temporal Volume 成交量
		A  string `json:"a"`  // 最佳五檔賣出價格
		B  string `json:"b"`  // 最佳五檔買入價格
		C  string `json:"c"`  // 股票代號
		D  string `json:"d"`  // 最近交易日期(YYYYMMDD)
		Ch string `json:"ch"` // Channel
		F  string `json:"f"`  // 最價五檔賣出數量
		G  string `json:"g"`  // 最佳五檔買入數量
		H  string `json:"h"`  // 最高
		L  string `json:"l"`  // 最低
		N  string `json:"n"`  // 公司簡稱
		O  string `json:"o"`  // 開盤價格
		Ex string `json:"ex"` // 上市 Or 上櫃
		T  string `json:"t"`  // 最近成交時刻(HH:MI:SS)
		U  string `json:"u"`  // 漲停價
		V  string `json:"v"`  // 當日累積成交量
		W  string `json:"w"`  // 跌停價
		Nf string `json:"nf"` // 公司全名
		Y  string `json:"y"`  // 昨收
		Z  string `json:"z"`  // 最近成交價
	} `json:"msgArray"`
}

func ShareHoldingQuery(query string) *StockApiJson {
	realtimeInfo := &StockApiJson{}
	queryURL := fmt.Sprintf(REALTIME_STOCK_INFORMATION_URL, query)
	res, err := http.Get(queryURL)
	if err != nil {
		log.Println("Error with ShareHoldingQuery function: ", err)
		return nil
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(realtimeInfo)
	if err != nil {
		log.Println("Error with json Decode: ", err)
		return nil
	}
	return realtimeInfo
}

func GenerateQueryString(stockNo ...string) string {
	queryString := ""
	for _, stock := range stockNo {
		queryString += ("tse_" + stock + ".tw|")
	}
	return queryString[:len(queryString)-1]
}
