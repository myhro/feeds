package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"

	"github.com/myhro/feeds/liquipedia"
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
		Title:   "Liquipedia - Player Transfers",
		Link:    &feeds.Link{Href: url},
		Created: time.Now().UTC(),
	}

	doc.Find(".divRow").Each(func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		created, err := liquipedia.Created(s)
		if err != nil {
			log.Fatal("liquipedia.Created: ", err)
		}
		description, err := liquipedia.Description(s)
		if err != nil {
			log.Fatal("liquipedia.Description: ", err)
		}

		item := &feeds.Item{
			Title:       liquipedia.Title(s),
			Created:     created,
			Link:        &feeds.Link{Href: liquipedia.Link(s)},
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
