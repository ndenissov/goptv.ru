package main

import (
	"flag"
	"fmt"
	"github.com/ndenissov/goptv.ru"
	"github.com/ndenissov/goptv.ru/cmd"
	"log"
)

func main() {
	query := flag.String("q", "", "Query")
	flag.Parse()
	cls, f := cmd.Prologue()
	if cls {
		defer f.Close()
	}
	search, err := ptv.Search(*query)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = fmt.Fprintln(f, search.String()); err != nil {
		log.Println(err)
	}
}
