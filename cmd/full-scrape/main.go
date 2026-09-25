package main

import (
	"flag"
	"fmt"
	ptv "github.com/ndenissov/goptv.ru"
	"github.com/ndenissov/goptv.ru/cmd"
	"log"
)

func main() {
	flag.Parse()
	cls, f := cmd.Prologue()
	if cls {
		defer f.Close()
	}
	plist, err := ptv.Plist()
	if err != nil {
		log.Fatalln(err)
	}
	for _, pl := range plist {
		if pl, err := ptv.Pl(pl.Key); err == nil {
			for _, chs := range pl.ExtInfos {
				if _, err = fmt.Fprintln(f, chs); err != nil {
					log.Println(err)
				}
			}
		} else {
			log.Println(err)
		}
	}
}
