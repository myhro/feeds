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
	req, err := http.NewRequest("GET", g.URL, nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add("Accept-Language", "en-us")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Client.Do: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("goquery.NewDocumentFromReader: %w", err)
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
