package htmlproxy

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/philcantcode/stonks/server/system"
)

func GetOpenInsiderPage(ticker string) []byte {
	url := "http://openinsider.com/search?q=" + ticker

	system.Log("Getting "+url, true)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return html
}

func GetOpenInsiderTargetID(ticker string, target string) []byte {
	url := "http://openinsider.com/search?q=" + ticker

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
