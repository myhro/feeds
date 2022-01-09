package generator

import (
	"fmt"
	"log"

	"github.com/myhro/feeds/errormap"
)

func Print(cmd string, atom func() (string, error)) {
	feed, err := atom()
	if err != nil {
		log.Fatal("atom: ", err)
	}

	list := errormap.List(cmd)
	if len(list) > 0 {
		for _, err := range list {
			log.Print(err)
		}

		log.Fatal("Error while generating feed")
	}

	fmt.Println(feed)
}
