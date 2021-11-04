package liquipedia

import (
	"crypto/sha256"
	"encoding/base64"
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

func ID(date time.Time, text string) string {
	sum := sha256.Sum256([]byte(text))
	hash := base64.RawStdEncoding.EncodeToString(sum[:])
	id := fmt.Sprintf("tag:liquipedia.net,%v:%v", date.Format("2006-01-02"), hash)
	return id
}

func Run(cmd *cobra.Command, args []string) {
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
			Id:          ID(created, description),
			Link:        &feeds.Link{Href: url},
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
