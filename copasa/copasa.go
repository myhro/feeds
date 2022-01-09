package copasa

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"

	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/generator"
)

const Command = "copasa"
const FeedTitle = "Copasa - Em Racionamento - Montes Claros"

func clean(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.Join(strings.Fields(s), " ")

	return s
}

func XML() (string, error) {
	baseURL := "http://www.copasa.com.br/wps/portal/internet/imprensa/noticias/plano-de-racionamento/filter"

	gen := generator.Generator{
		CSS:   ".item",
		Title: FeedTitle,
		URL:   baseURL + "?area=/site-copasa-conteudos/internet/perfil/imprensa/noticias/plano-de-racionamento/racionamento-encerrado/co-montes-claros", //nolint:lll
	}

	gen.Parse = func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		date := s.Find(".data").Text()
		title := s.Find(".titulo-imprensa").Text()
		link, _ := s.Find(".titulo-imprensa").Find("a").Attr("href")
		description := s.Find(".texto-imprensa").Text()

		created, err := time.Parse("02 Jan 2006", clean(date))
		if err != nil {
			errormap.Store(Command, fmt.Errorf("time.Parse: %w", err))
		}

		item := &feeds.Item{
			Title:       clean(title),
			Created:     created,
			Link:        &feeds.Link{Href: baseURL + link},
			Description: clean(description),
		}
		gen.Feed.Items = append(gen.Feed.Items, item)
	}

	atom, err := gen.Generate()
	if err != nil {
		return "", fmt.Errorf("gen.Generate: %w", err)
	}

	return atom, nil
}
