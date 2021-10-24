package liquipedia

import (
	"os"
	"testing"

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
		s.Zero(sel.Find("img").Length())
	}
}
