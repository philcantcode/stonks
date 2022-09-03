package htmlproxy

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/philcantcode/stonks/server/system"
)

func GetOptiononicsPage(ticker string) []byte {
	url := "https://www.optionistics.com/quotes/option-prices/" + ticker

	system.Log("Getting "+url, true)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return html
}

func GetOptiononicsPart(ticker string, target string) []byte {
	url := "https://www.optionistics.com/quotes/option-prices/" + ticker

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	results := ""

	doc.Find(target).Each(func(i int, s *goquery.Selection) {
		inside_html, _ := s.Html()
		results += inside_html
	})

	return []byte(results)
}
