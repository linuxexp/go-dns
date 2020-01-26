package main

import (
	"fmt"
	"golang.org/x/net/dns/dnsmessage"
	"log"
	"net"
)

const port = 1023
const packetSize = 512

func main() {

	fmt.Println("Listening at port ", port)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: port})

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		buf := make([]byte, packetSize)
		_, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		var m dnsmessage.Message
		err = m.Unpack(buf)

		if err != nil {
			log.Println(err)
			continue
		}

		if len(m.Questions) > 0 {
			question := m.Questions[0]

			log.Println("Name: ", question.Name)
			log.Println("Type: ", question.Type)

			answer := dnsmessage.Resource{
				Header: dnsmessage.ResourceHeader{
					Name:   question.Name,
					Type:   question.Type,
					Class:  question.Class,
					TTL:    0,
					Length: 0,
				},
				Body: &dnsmessage.AResource{A: [4]byte{127, 0, 0, 1}},
			}

			m.Answers = []dnsmessage.Resource{answer}
			m.Response = true
			packed, _ := m.Pack()
			go conn.WriteTo(packed, addr)
		}

	}
}
