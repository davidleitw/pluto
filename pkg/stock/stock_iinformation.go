package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var REALTIME_STOCK_INFORMATION string = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=%s&json=1&delay=0"

func GetStockInformation(stock string) {
	url := fmt.Sprintf(REALTIME_STOCK_INFORMATION, stock)
	log.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Println("http get stock url error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	jsonMap := make(map[string]interface{})
	jsonMsg := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &jsonMap)
	if err != nil {
		log.Println(err)
	}
	for k, v := range jsonMap {
		log.Println(k, v)
		if k == "msgArray" {
			// err = json.Unmarshal([]byte(v), &jsonMsg)
		}
	}
	for k, v := range jsonMsg {
		log.Println(k, v)
	}
}
