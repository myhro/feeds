package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
)

func main() {
	url := "https://liquipedia.net/dota2/Portal:Transfers"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http.Get: ", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal("goquery.NewDocumentFromResponse: ", err)
	}

	feed := &feeds.Feed{
		Title:   "Liquipedia - Player Transfers: Latest",
		Link:    &feeds.Link{Href: url},
		Created: time.Now().UTC(),
	}

	doc.Find(".divRow").Each(func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		date := s.Find(".Date").Text()
		created, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Fatal("time.Parse: ", err)
		}

		title := s.Find(".Name").Text()
		title = strings.TrimSpace(title)
		link, _ := s.Find(".Ref").Find("a").Attr("href")
		description, _ := s.Html()

		item := &feeds.Item{
			Title:       title,
			Created:     created,
			Link:        &feeds.Link{Href: link},
			Description: description,
		}
		feed.Items = append(feed.Items, item)
	})

	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal("feed.ToAtom: ", err)
	}
	fmt.Println(atom)
}
