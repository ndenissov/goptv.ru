package cmd

import (
	"flag"
	"fmt"
	"github.com/ndenissov/goptv.ru/m3u"
	"log"
	"os"
)

var Output, Author string

func init() {
	flag.StringVar(&Output, "o", "", "Output filename (empty for stdout)")
	flag.StringVar(&Author, "a", "", "Metadata author name")
}

func WriteAuthor(f *os.File) {
	if Author != "" {
		if _, err := fmt.Fprintln(f, m3u.EXTM3UHeader(Author)); err != nil {
			log.Println(err)
		}
	}
}

func Prologue() (closable bool, f *os.File) {
	var err error
	if Output == "" {
		f = os.Stdout
	} else {
		if f, err = os.OpenFile(Output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err == nil {
			closable = true
		} else {
			log.Fatalln(err)
		}
	}
	WriteAuthor(f)
	return
}
