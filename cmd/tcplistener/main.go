package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

const port = ":42069"

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)

		currLineContent := ""
		for {
			b := make([]byte, 8, 8)
			n, err := f.Read(b)
			if err != nil {
				
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
			continue
		}

		fmt.Println("connection accepted")

		go func(c net.Conn) {
			defer fmt.Println("connection closed")

			for line := range getLinesChannel(c) {
				fmt.Printf("read: %s\n", line)
			}
		}(conn)
	}
}
