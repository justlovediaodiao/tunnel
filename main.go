package main

import (
	"flag"
	"log"
)

func main() {
	var laddr, taddr string
	flag.StringVar(&laddr, "l", "", "listen address")
	flag.StringVar(&taddr, "t", "", "target address")
	flag.Parse()
	if laddr == "" || taddr == "" {
		flag.Usage()
		return
	}
	go func() {
		var err = startTCP(laddr, taddr)
		if err != nil {
			log.Print(err)
		}
	}()
	var err = startUDP(laddr, taddr)
	if err != nil {
		log.Print(err)
	}
}
