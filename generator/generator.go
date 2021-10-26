package generator

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
)

type Generator struct {
	CSS   string
	Feed  *feeds.Feed
	Parse func(i int, s *goquery.Selection)
	Title string
	URL   string
}

func (g *Generator) Generate() (string, error) {
	resp, err := http.Get(g.URL)
	if err != nil {
		return "", fmt.Errorf("http.Get: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "", fmt.Errorf("goquery.NewDocumentFromResponse: %w", err)
	}

	g.Feed = &feeds.Feed{
		Title:   g.Title,
		Link:    &feeds.Link{Href: g.URL},
		Created: time.Now().UTC(),
	}

	doc.Find(g.CSS).Each(g.Parse)

	atom, err := g.Feed.ToAtom()
	if err != nil {
		return "", fmt.Errorf("Generator.Feed.ToAtom: %w", err)
	}
	return atom, nil
}