package source

import (
	"fmt"
	"github.com/ndenissov/goptv.ru/pkg"
	"strconv"
)

type Source struct {
	Country, Name, City, Key string

	Channels int
}
type Sources []*Source

func (s *Source) String() string {
	return fmt.Sprintf("%s %s, г. %s, %s, источников: %d", s.Country, s.Name, s.City, s.Key, s.Channels)
}

func (s Sources) String() string { return pkg.NewlinesJoin(s) }

func newSource(match []string) (p *Source) {
	if len(match) != 0 {
		p = &Source{Country: match[1], Name: match[2], City: match[3], Key: match[4]}
		p.Channels, _ = strconv.Atoi(match[5])
	}
	return
}
