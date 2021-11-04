package liquipedia

import (
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/suite"
)

type LiquipediaTestSuite struct {
	suite.Suite

	InputName map[*goquery.Selection]string

	Brame                   *goquery.Selection
	BrameDiffRef            *goquery.Selection
	Creepwave               *goquery.Selection
	ThePrimeAndArmyGeniuses *goquery.Selection
}

func TestLiquipediaTestSuite(t *testing.T) {
	suite.Run(t, new(LiquipediaTestSuite))
}

func (s *LiquipediaTestSuite) SetupTest() {
	s.InputName = make(map[*goquery.Selection]string)

	s.Brame = s.LoadFixture("testdata/brame.html")
	s.InputName[s.Brame] = "Brame"

	s.BrameDiffRef = s.LoadFixture("testdata/brame-diff-ref.html")
	s.InputName[s.BrameDiffRef] = "BrameDiffRef"

	s.Creepwave = s.LoadFixture("testdata/creepwave.html")
	s.InputName[s.Creepwave] = "Creepwave"

	s.ThePrimeAndArmyGeniuses = s.LoadFixture("testdata/the-prime-and-army-geniuses.html")
	s.InputName[s.ThePrimeAndArmyGeniuses] = "ThePrimeAndArmyGeniuses"
}

func (s *LiquipediaTestSuite) TearDownTest() {
	s.Brame = nil
	s.Creepwave = nil
	s.ThePrimeAndArmyGeniuses = nil
}

func (s *LiquipediaTestSuite) LoadFixture(fixture string) *goquery.Selection {
	file, err := os.Open(fixture)
	s.Nil(err)
	doc, err := goquery.NewDocumentFromReader(file)
	s.Nil(err)
	return doc.Children()
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
	brameID := "tag:liquipedia.net,2021-10-21:Z0vW1sdL0qyi7i4s/gzMqfefZC6XOrUlK3XfR7eD4wI"
	table := []struct {
		in  *goquery.Selection
		out string
	}{
		{
			in:  s.Brame,
			out: brameID,
		},
		{
			in:  s.BrameDiffRef,
			out: brameID,
		},
		{
			in:  s.Creepwave,
			out: "tag:liquipedia.net,2021-10-22:CxJ7qmuZME1yW1XZbBWK3tNVXS/n9MdyHlY6vW3o6n8",
		},
		{
			in:  s.ThePrimeAndArmyGeniuses,
			out: "tag:liquipedia.net,2021-10-22:ZSn2YVaXK+alPQgptA3LwU5MfddMQxvOqD1mw/iJ23k",
		},
	}

	for _, tt := range table {
		id, err := ID(tt.in)
		s.Nil(err)
		s.Equal(tt.out, id, s.InputName[tt.in])
	}
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
