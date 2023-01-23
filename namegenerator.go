package codenamegenerator

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func NameGenerate() {
	s := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(s)
	u := "http://www.codenamegenerator.com"

	r1 := CommonCodeNames[r.Intn(len(CommonCodeNames))]
	r2 := CommonCodeNames[r.Intn(len(CommonCodeNames))]
	r3 := CommonCodeNames[r.Intn(len(CommonCodeNames))]

	resp, err := http.PostForm(u, url.Values{"prefix": {r1}, "dictionary": {r2}, "suffix": {r3}})

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var words []string

	doc.Find("b:first-child").Each(func(i int, s *goquery.Selection) {
		words = append(words, s.Text())
	})

	l := strings.ToLower(words[0])

	split := strings.Split(l, "\n")

	fmt.Printf("%s-%s-%s", split[0], split[1], split[2])
}
