package liquipedia

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"

	"github.com/myhro/feeds/errormap"
	"github.com/myhro/feeds/generator"
)

const Command = "liquipedia"
const FeedTitle = "Liquipedia - Player Transfers"

func Date(s *goquery.Selection) (time.Time, error) {
	date := s.Find(".Date").Text()

	created, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, fmt.Errorf("time.Parse: %w", err)
	}

	return created, nil
}

func CleanDescription(s *goquery.Selection) *goquery.Selection {
	s = s.Clone()

	s.Find(".flag").Remove()
	s.Find(".team-template-darkmode").Remove()
	s.Find("img").Each(func(i int, img *goquery.Selection) {
		alt, ok := img.Attr("alt")
		if ok {
			img.ReplaceWithHtml(alt)
		}
	})

	return s
}

func Description(s *goquery.Selection) (string, error) {
	s = CleanDescription(s)

	html, err := s.Html()
	if err != nil {
		return "", fmt.Errorf("Selection.Html: %w", err)
	}

	return html, nil
}

func DescriptionWithoutRef(s *goquery.Selection) (string, error) {
	s = CleanDescription(s)
	s.Find(".Ref").Remove()

	html, err := s.Html()
	if err != nil {
		return "", fmt.Errorf("Selection.Html: %w", err)
	}

	return html, nil
}

func ID(s *goquery.Selection) (string, error) {
	date, err := Date(s)
	if err != nil {
		return "", fmt.Errorf("Date: %w", err)
	}

	desc, err := DescriptionWithoutRef(s)
	if err != nil {
		return "", fmt.Errorf("DescriptionWithoutRef: %w", err)
	}

	sum := sha256.Sum256([]byte(desc))
	hash := base64.RawStdEncoding.EncodeToString(sum[:])
	id := fmt.Sprintf("tag:liquipedia.net,%v:%v", date.Format("2006-01-02"), hash)

	return id, nil
}

func Title(s *goquery.Selection) string {
	title := s.Find(".Name").Text()
	return strings.TrimSpace(title)
}

func XML() (string, error) {
	url := "https://liquipedia.net/dota2/Portal:Transfers"

	gen := generator.Generator{
		CSS:   ".divRow",
		Title: FeedTitle,
		URL:   url,
	}

	gen.Parse = func(i int, s *goquery.Selection) {
		if i >= 10 {
			return
		}

		created, err := Date(s)
		if err != nil {
			errormap.Store(Command, fmt.Errorf("date: %w", err))
		}

		description, err := Description(s)
		if err != nil {
			errormap.Store(Command, fmt.Errorf("description: %w", err))
		}

		id, err := ID(s)
		if err != nil {
			errormap.Store(Command, fmt.Errorf("id: %w", err))
		}

		item := &feeds.Item{
			Title:       Title(s),
			Created:     created,
			Id:          id,
			Link:        &feeds.Link{Href: url},
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
