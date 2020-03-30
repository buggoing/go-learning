package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"udp"
)

func handleClient(buf []byte, addr *net.UDPAddr) {
	msg := new(udp.Message)
	tmpBuf := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(tmpBuf)
	if err := decoder.Decode(msg); err != nil {
		fmt.Printf("failed to decode mes err: %v", err)
	}
	fmt.Printf("msg: %+v", *msg)

}

func main() {
	addr := ":5000"
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatalf("failed to resolve upd addr err: %v", err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	var buf [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Printf("failed to read from udp err: %v", err)
			continue
		}
		go handleClient(buf[:n], addr) // handle connections concurrently
	}
}
