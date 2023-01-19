package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	CommonCodeNames = []string{"animals", "apple", "birds", "bugs", "dogs", "fishes",
		"flowers", "metal", "microsoft", "snakes", "units", "unusual_animals", "us_cities",
		"us_counties", "us_state_capitals", "us_states", "world_capitals", "world_countries"}
)

func main() {
	s := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(s)

	r1 := CommonCodeNames[r.Intn(len(CommonCodeNames))]
	r2 := CommonCodeNames[r.Intn(len(CommonCodeNames))]
	r3 := CommonCodeNames[r.Intn(len(CommonCodeNames))]

	resp, err := http.PostForm("https://www.codenamegenerator.com/",
		url.Values{"prefix": {r1}, "dictionary": {r2}, "suffix": {r3}})

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

	// str := strings.Join(words, " ")

	fmt.Println(words[0])
}
