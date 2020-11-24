package main

import (
	"io"
	"log"
	"net"
	"time"
)

func startTCP(laddr string, taddr string) error {
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		return err
	}
	log.Printf("tcp tunnel %s <----> %s", laddr, taddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handleStream(conn, taddr)
	}
}

func handleStream(conn net.Conn, taddr string) {
	defer conn.Close()
	rc, err := net.Dial("tcp", taddr)
	if err != nil {
		log.Printf("dial to %s error: %v", taddr, err)
		return
	}
	defer rc.Close()
	log.Printf("tcp %s <----> %s <----> %s", conn.RemoteAddr().String(), conn.LocalAddr().String(), rc.RemoteAddr().String())
	err = relayStream(conn, rc)
	if err != nil {
		log.Printf("relay stream error: %v", err)
	}
}

// relayStream copy stream from left to right and right to left.
func relayStream(left, right net.Conn) error {
	ch := make(chan error, 2)

	go func() {
		_, err := io.Copy(right, left)
		ch <- err
		right.SetReadDeadline(time.Now()) // unblock read on right
	}()

	_, err := io.Copy(left, right)
	ch <- err
	left.SetReadDeadline(time.Now()) // unblock read on left

	// the first err is relay error reason.
	err = <-ch
	// wait goroutine done
	<-ch
	return err
}
