package oldnewthing

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"

	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/generator"
)

const Command = "oldnewthing"
const FeedTitle = "The Old New Thing"

func CleanDescription(s *goquery.Selection) string {
	s = s.Clone()
	s.Find(".entry-header").Remove()
	s.Find(".entry-footer").Remove()

	return strings.TrimSpace(s.Text())
}

func XML() (string, error) {
	gen := generator.Generator{
		CSS:   ".entry-area",
		Title: FeedTitle,
		URL:   "https://devblogs.microsoft.com/oldnewthing/",
	}

	gen.Parse = func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		date := s.Find(".entry-post-date").Text()
		date = strings.TrimSpace(date)
		title := s.Find(".entry-title").Text()
		link, _ := s.Find(".entry-title").Find("a").Attr("href")
		description := CleanDescription(s.Find(".entry-content"))

		created, err := time.Parse("January 2, 2006", date)
		if err != nil {
			errormap.Store(Command, fmt.Errorf("time.Parse: %w", err))
		}

		item := &feeds.Item{
			Title:       title,
			Created:     created,
			Link:        &feeds.Link{Href: link},
			Description: description,
		}
		gen.Feed.Items = append(gen.Feed.Items, item)
	}

	atom, err := gen.Generate()
	if err != nil {
		return "", fmt.Errorf("gen.Generate: %w", err)
	}

	return atom, nil
}
