package internal

import (
	"github.com/ndenissov/goptv.ru/pkg"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const SearchURL = "https://proxytv.ru/iptv/php/srch.php"

var DefaultSearchHeaders = http.Header{
	"Referer": {"https://proxytv.ru/"}, "User-Agent": {"Mozilla/5.0"},
	"Content-Type": {"application/x-www-form-urlencoded"}}

func Search(query string) (lines []string, err error) {
	request, err := http.NewRequest("POST", SearchURL, strings.NewReader(pkg.Concat("udpxyaddr=", url.QueryEscape(query))))
	if err != nil {
		return
	}
	request.Header = DefaultSearchHeaders
	do, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	var b []byte
	if do.ContentLength == -1 {
		b, err = io.ReadAll(do.Body)
	} else {
		b = make([]byte, do.ContentLength)
		_, err = do.Body.Read(b)
	}
	if err != nil {
		return
	}
	div := strings.SplitN(string(b), "<hr>", 3)[1]
	if hr := strings.LastIndex(div, "<br>"); hr != -1 {
		div = div[:hr]
	}
	return strings.Split(div, "<br>"), nil
}
