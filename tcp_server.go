package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
	if err != nil {
		return // e.g., client disconnected
	}

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello world")
	//!+
	for {
		conn, err := listener.Accept()
		fmt.Println("c", conn)

		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
