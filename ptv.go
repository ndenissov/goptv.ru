package ptv

import (
	"fmt"
	"github.com/ndenissov/goptv.ru/internal"
	"github.com/ndenissov/goptv.ru/m3u"
	"github.com/ndenissov/goptv.ru/pkg"
	"github.com/ndenissov/goptv.ru/source"
	"strings"
)

func searchFunc[T any](post func([]string) T) func(string) (T, error) {
	return func(query string) (T, error) {
		lines, err := internal.Search(query)
		return post(lines), err
	}
}

// Functions which no automatically insert prefix at the start of query
var (
	QueryProvider = searchFunc(source.NewProviders)
	QueryPlist    = searchFunc(source.NewPlists)
	QueryCh       = searchFunc(m3u.NewChannels)
	QueryPl       = searchFunc(source.NewPl)
)

const (
	plist    = "plist"
	ch       = "ch:"
	pl       = "pl:"
	provider = "provider"
)

// Functions which automatically insert prefix at the start of query

func Provider(name string) (source.Sources, error) { return QueryProvider(provider) }
func Plist() (source.Sources, error)               { return QueryPlist(plist) }

func Ch(name string) ([]*m3u.Channel, error) { return QueryCh(pkg.Concat(ch, name)) }
func Pl(name string) (*source.Pl, error)     { return QueryPl(pkg.Concat(pl, name)) }

func Search(q string) (fmt.Stringer, error) {
	if strings.EqualFold(q, plist) {
		return QueryPlist(q)
	} else if len(q) >= len(ch) && strings.EqualFold(q[:len(ch)], ch) {
		return QueryCh(q)
	} else if len(q) >= len(pl) && strings.EqualFold(q[:len(pl)], pl) {
		return QueryPl(q)
	}
	return QueryProvider(q)
}
