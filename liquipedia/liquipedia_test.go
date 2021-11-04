package liquipedia

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/suite"
)

type LiquipediaTestSuite struct {
	suite.Suite

	Brame                   *goquery.Selection
	Creepwave               *goquery.Selection
	ThePrimeAndArmyGeniuses *goquery.Selection
}

func TestLiquipediaTestSuite(t *testing.T) {
	suite.Run(t, new(LiquipediaTestSuite))
}

func (s *LiquipediaTestSuite) SetupTest() {
	file, err := os.Open("testdata/brame.html")
	s.Nil(err)
	doc, err := goquery.NewDocumentFromReader(file)
	s.Nil(err)
	s.Brame = doc.Children()

	file, err = os.Open("testdata/creepwave.html")
	s.Nil(err)
	doc, err = goquery.NewDocumentFromReader(file)
	s.Nil(err)
	s.Creepwave = doc.Children()

	file, err = os.Open("testdata/the-prime-and-army-geniuses.html")
	s.Nil(err)
	doc, err = goquery.NewDocumentFromReader(file)
	s.Nil(err)
	s.ThePrimeAndArmyGeniuses = doc.Children()
}

func (s *LiquipediaTestSuite) TearDownTest() {
	s.Brame = nil
	s.Creepwave = nil
	s.ThePrimeAndArmyGeniuses = nil
}

func (s *LiquipediaTestSuite) TestDescriptionLinks() {
	table := []struct {
		in   *goquery.Selection
		link string
	}{
		{
			in:   s.Brame,
			link: "<a href=\"/dota2/Brame\" title=\"Brame\">Brame</a>",
		},
		{
			in:   s.Creepwave,
			link: "<a href=\"/dota2/Creepwave\" title=\"Creepwave\">Creepwave</a>",
		},
		{
			in:   s.ThePrimeAndArmyGeniuses,
			link: "<a href=\"/dota2/The_Prime\" title=\"The Prime\">The Prime</a>",
		},
		{
			in:   s.ThePrimeAndArmyGeniuses,
			link: "<a href=\"/dota2/Army_Geniuses\" title=\"Army Geniuses\">Army Geniuses</a>",
		},
	}

	for _, tt := range table {
		desc, err := Description(tt.in)
		s.Nil(err)
		s.Contains(desc, tt.link)
		s.Equal(1, strings.Count(desc, tt.link))
	}
}

func (s *LiquipediaTestSuite) TestDescriptionRemovedFlags() {
	table := []struct {
		in    *goquery.Selection
		flags []string
	}{
		{
			in: s.Brame,
			flags: []string{
				"Ukraine",
			},
		},
		{
			in: s.Creepwave,
			flags: []string{
				"Belarus",
				"Bulgaria",
				"Jordan",
				"Netherlands",
				"Russia",
			},
		},
		{
			in: s.ThePrimeAndArmyGeniuses,
			flags: []string{
				"Indonesia",
			},
		},
	}

	for _, tt := range table {
		desc, err := Description(tt.in)
		s.Nil(err)
		for _, flag := range tt.flags {
			s.NotContains(desc, flag)
		}
	}
}

func (s *LiquipediaTestSuite) TestDescriptionRemovedImages() {
	list := []*goquery.Selection{
		s.Brame,
		s.Creepwave,
		s.ThePrimeAndArmyGeniuses,
	}

	for _, sel := range list {
		desc, err := Description(sel)
		s.Nil(err)
		s.NotEmpty(desc)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(desc))
		s.Nil(err)
		s.Zero(doc.Find("img").Length())
	}
}

func (s *LiquipediaTestSuite) TestDescriptionRemovedRef() {
	list := []*goquery.Selection{
		s.Brame,
		s.Creepwave,
		s.ThePrimeAndArmyGeniuses,
	}

	for _, sel := range list {
		desc, err := DescriptionWithoutRef(sel)
		s.Nil(err)
		s.NotEmpty(desc)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(desc))
		s.Nil(err)
		s.Zero(doc.Find(".Ref").Length())
	}
}

func (s *LiquipediaTestSuite) TestID() {
	date := time.Date(2021, time.October, 1, 12, 0, 0, 0, time.UTC)
	id := ID(date, "example text")
	s.Equal("tag:liquipedia.net,2021-10-01:DpSuNtpv8DmSpX/dvfRyi2CdDX/m6wGfqfG5tbVA2DU", id)
}

func (s *LiquipediaTestSuite) TestTitle() {
	table := []struct {
		in    *goquery.Selection
		title string
	}{
		{
			in:    s.Brame,
			title: "Nefrit",
		},
		{
			in:    s.Creepwave,
			title: "ATF Chu Crystallis Fishman hansha",
		},
		{
			in:    s.ThePrimeAndArmyGeniuses,
			title: "Azur4",
		},
	}

	for _, tt := range table {
		// Ensure 'Description()' has no side-effects
		Description(tt.in)
		s.Equal(tt.title, Title(tt.in))
	}
}
