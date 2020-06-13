package main

import (
	"flag"
	"fmt"
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
			fmt.Println(err)
		}
	}()
	var err = startUDP(laddr, taddr)
	if err != nil {
		fmt.Println(err)
	}
}
