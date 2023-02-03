package codenamegenerator

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// NameGenerate takes the resp.Body from getRawCodeName and outputs the string
// in the desired format
func NameGenerate(io.ReadCloser) string {
	doc, err := goquery.NewDocumentFromReader(getRawCodeName())

	if err != nil {
		log.Fatal(err)
	}

	var words []string

	doc.Find("b:first-child").Each(func(i int, s *goquery.Selection) {
		words = append(words, s.Text())
	})

	l := strings.ToLower(words[0])

	split := strings.Split(l, "\n")

	sp0 := getSingleWord(split[0])
	sp1 := getSingleWord(split[1])
	sp2 := getSingleWord(split[2])

	codename := fmt.Sprintf("%s-%s-%s", sp0, sp1, sp2)

	return codename
}

// getRawCodeName queries the website for the raw resp.Body containing the codename
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
