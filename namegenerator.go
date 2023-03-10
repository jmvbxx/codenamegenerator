package codenamegenerator

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gobuffalo/flect"
)

func NameGenerate() string {
	doc, err := goquery.NewDocumentFromReader(getRawCodeName())

	if err != nil {
		log.Fatal(err)
	}

	var words []string

	doc.Find("b:first-child").Each(func(i int, s *goquery.Selection) {
		words = append(words, s.Text())
	})

	l := strings.ToLower(words[0])

	cn := flect.Dasherize(l)

	cn_array := strings.Split(cn, "-")
	cn_array = cn_array[:3]
	cn = strings.Join(cn_array, "-")

	return cn

}

func getRawCodeName() io.ReadCloser {
	s := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(s)
	u := "https://www.codenamegenerator.com"

	r1 := CommonCodeNames[r.Intn(len(CommonCodeNames))]
	r2 := CommonCodeNames[r.Intn(len(CommonCodeNames))]
	r3 := CommonCodeNames[r.Intn(len(CommonCodeNames))]

	resp, err := http.PostForm(u, url.Values{"prefix": {r1}, "dictionary": {r2}, "suffix": {r3}})

	if err != nil {
		log.Fatal(err)
	}

	// defer resp.Body.Close()

	return resp.Body
}
