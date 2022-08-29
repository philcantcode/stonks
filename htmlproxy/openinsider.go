package htmlproxy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"github.com/philcantcode/stonks/system"
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

func GetOpenInsiderResults(ticker string) []byte {
	doc, err := html.Parse(strings.NewReader(string(GetOpenInsiderPage(ticker))))
	if err != nil {
		panic("Fail to parse!")
	}

	r1 := getElementById(doc, "results")
	var b bytes.Buffer
	html.Render(&b, r1)

	fmt.Println(b.String())

	return []byte(b.String())
}
