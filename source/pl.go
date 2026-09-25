package source

import (
	"github.com/ndenissov/goptv.ru/m3u"
	"strings"
)

// Pl One provider m3u
type Pl struct {
	*m3u.Header
	ExtInfos []*m3u.Extinf
}

func (pl *Pl) String() string {
	s := make([]string, 1, len(pl.ExtInfos)+1)
	s[0] = pl.Header.String()
	for _, info := range pl.ExtInfos {
		s = append(s, info.String())
	}
	return strings.Join(s, "\n")
}
func NewPl(s []string) *Pl {
	if len(s) != 0 {
		if h := m3u.NewHeader(s[0]); h != nil {
			return &Pl{Header: h, ExtInfos: m3u.NewExtInfos(s[1:])}
		}
	}
	return nil
}
