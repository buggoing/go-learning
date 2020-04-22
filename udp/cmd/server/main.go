package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
	"udp"
)

func handleClient(buf []byte, addr *net.UDPAddr) {
	msg := new(udp.Message)
	tmpBuf := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(tmpBuf)
	if err := decoder.Decode(msg); err != nil {
		fmt.Printf("failed to decode msg err: %v", err)
	}
	fmt.Printf("msg: %+v", *msg)
}

func handleClientJSON(buf []byte, addr *net.UDPAddr) {

	msg := new(udp.Message)
	if err := json.Unmarshal(buf, msg); err != nil {
		fmt.Printf("failed to unmarshall msg err: %v\n", err)
	}
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
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Printf("failed to read from udp err: %v", err)
			continue
		}
		fmt.Printf("received: %v from %+v \n", string(buf[:]), *addr)
		reply := []byte(fmt.Sprintf("server %v\n", time.Now()))
		_, err = conn.WriteTo(reply, addr)
		if err != nil {
			fmt.Printf("failed to send err: %v", err)
		}
		// go handleClientJSON(buf[:n], addr)
	}
}
