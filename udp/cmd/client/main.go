package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"udp"
)

func main() {
	addr := "127.0.0.1:5000"
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatalf("failed to resolve upd addr err: %v", err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("failed to dial udp err: %v", err)
	}
	msg := udp.Message{
		DeviceID: "kasdjfalksdf",
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("jsonMsg err: %v", err)
	}
	fmt.Printf("jsonMsg: %v", string(jsonMsg))
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(msg)
	n, err := conn.Write(buf.Bytes())
	if err != nil {
		log.Fatalf("failed to write msg err: %v", err)
	}
	fmt.Printf("send: %v", n)
	ch := make(chan int)
	<-ch
}
