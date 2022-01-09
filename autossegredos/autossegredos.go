package autossegredos

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"

	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/generator"
)

const Command = "autossegredos"
const FeedTitle = "Autos Segredos - Arquivos Segredos"

func XML() (string, error) {
	gen := generator.Generator{
		CSS:   ".tdb_module_loop",
		Title: FeedTitle,
		URL:   "https://www.autossegredos.com.br/category/segredos/",
	}

	gen.Parse = func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		date, _ := s.Find(".td-post-date").Find("time").Attr("datetime")
		title := s.Find(".td-module-title").Text()
		link, _ := s.Find(".td-module-title").Find("a").Attr("href")
		description := s.Find(".td-excerpt").Text()

		created, err := time.Parse(time.RFC3339, date)
		if err != nil {
			errormap.Store(Command, fmt.Errorf("time.Parse: %w", err))
		}

		item := &feeds.Item{
			Title:       title,
			Created:     created,
			Link:        &feeds.Link{Href: link},
			Description: strings.TrimSpace(description),
		}
		gen.Feed.Items = append(gen.Feed.Items, item)
	}

	atom, err := gen.Generate()
	if err != nil {
		return "", fmt.Errorf("gen.Generate: %w", err)
	}

	return atom, nil
}
