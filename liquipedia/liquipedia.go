package liquipedia

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
	"github.com/spf13/cobra"
)

const FeedTitle = "Liquipedia - Player Transfers"

func Created(s *goquery.Selection) (time.Time, error) {
	date := s.Find(".Date").Text()
	created, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, fmt.Errorf("time.Parse: %w", err)
	}
	return created, nil
}

func Description(s *goquery.Selection) (string, error) {
	s = s.Clone()
	s.Find(".flag").Remove()
	s.Find(".team-template-darkmode").Remove()
	s.Find("img").Each(func(i int, img *goquery.Selection) {
		alt, ok := img.Attr("alt")
		if ok {
			img.ReplaceWithHtml(alt)
		}
	})

	html, err := s.Html()
	if err != nil {
		return "", fmt.Errorf("Selection.Html: %w", err)
	}
	return html, nil
}

func Link(s *goquery.Selection) string {
	link, ok := s.Find(".Ref").Find("a").Attr("href")
	if !ok {
		return ""
	}
	return link
}

func Run(cmd *cobra.Command, args []string) {
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
		Title:   FeedTitle,
		Link:    &feeds.Link{Href: url},
		Created: time.Now().UTC(),
	}

	doc.Find(".divRow").Each(func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		created, err := Created(s)
		if err != nil {
			log.Fatal("Created: ", err)
		}
		description, err := Description(s)
		if err != nil {
			log.Fatal("Description: ", err)
		}

		item := &feeds.Item{
			Title:       Title(s),
			Created:     created,
			Link:        &feeds.Link{Href: Link(s)},
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

func Title(s *goquery.Selection) string {
	title := s.Find(".Name").Text()
	return strings.TrimSpace(title)
}
