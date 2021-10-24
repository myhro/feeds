package liquipedia

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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

func Title(s *goquery.Selection) string {
	title := s.Find(".Name").Text()
	return strings.TrimSpace(title)
}
