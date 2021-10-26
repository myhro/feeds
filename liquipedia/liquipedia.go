package liquipedia

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
	"github.com/spf13/cobra"

	"github.com/myhro/feeds/generator"
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
	gen := generator.Generator{
		CSS:   ".divRow",
		Title: FeedTitle,
		URL:   "https://liquipedia.net/dota2/Portal:Transfers",
	}

	gen.Parse = func(i int, s *goquery.Selection) {
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
		gen.Feed.Items = append(gen.Feed.Items, item)
	}

	atom, err := gen.Generate()
	if err != nil {
		log.Fatal("Generator.Generate: ", err)
	}
	fmt.Println(atom)
}

func Title(s *goquery.Selection) string {
	title := s.Find(".Name").Text()
	return strings.TrimSpace(title)
}
