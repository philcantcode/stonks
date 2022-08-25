package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/philcantcode/stonks/credentials"
)

var client = http.Client{
	Timeout: time.Second * 30,
}

func QuerySymbol(symbol string) YFStock {

	url := "https://yahoo-finance97.p.rapidapi.com/stock-info"

	payload := strings.NewReader("symbol=" + symbol)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", credentials.YAHOO_API_KEY)
	req.Header.Add("X-RapidAPI-Host", "yahoo-finance97.p.rapidapi.com")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request, %s", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

	yfStockData := YFStock{}
	json.Unmarshal(body, &yfStockData)

	return yfStockData

}
