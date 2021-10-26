package copasa

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

const FeedTitle = "Copasa - Em Racionamento - Montes Claros"

func clean(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func Run(cmd *cobra.Command, args []string) {
	baseURL := "http://www.copasa.com.br/wps/portal/internet/imprensa/noticias/plano-de-racionamento/filter"

	gen := generator.Generator{
		CSS:   ".item",
		Title: FeedTitle,
		URL:   baseURL + "?area=/site-copasa-conteudos/internet/perfil/imprensa/noticias/plano-de-racionamento/racionamento-encerrado/co-montes-claros",
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
			log.Fatal("time.Parse: ", err)
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
		log.Fatal("Generator.Generate: ", err)
	}
	fmt.Println(atom)
}
