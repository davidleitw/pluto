package stock

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

/******
來源: 政府開放資料平台/上市公司基本資料
url: http://mopsfin.twse.com.tw/opendata/t187ap03_L.csv
******/

var LISTED_COMPANY_DATA_URL string = "http://mopsfin.twse.com.tw/opendata/t187ap03_L.csv"

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	counter := &WriteCounter{}
	// Write the body to file
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	return err
}

func GetStockNameByNumber(number string) (string, error) {
	file, err := os.Open("listed_company.csv")
	if err != nil {
		return "", err
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if record[1] == number {
			return record[3], nil
		}
	}
	return "", nil
}
