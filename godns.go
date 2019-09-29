package main

import (
	"fmt"
	"godns/types/dnsmessage"
	"godns/utils"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Started godns with arguments=", os.Args[1:])

	// var header dnsmessage.Header
	// header.Response = true
	// header.OpCode = dnsmessage.OpCodeInverse
	// fmt.Println(header)

	// var question dnsmessage.Question
	// question.QName = "a"
	// question.QType = dnsmessage.ResourceRecordALL
	// question.QClass = dnsmessage.ClassInternet
	// utils.PrintByteArray(question.Serialize())

	sendUDP("8.8.8.8:53", prepareTestMessage())
}

func sendUDP(ipAddress string, message dnsmessage.Message) (err error) {
	fmt.Println("Sending message")
	utils.PrintByteArray(message.Serialize())

	conn, err := net.Dial("udp", ipAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(message.Serialize())
	if err != nil {
		return err
	}

	deadline := time.Now().Add(time.Duration(5) * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	readCount, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	result := buffer[0:readCount]

	fmt.Println("Received response")
	utils.PrintByteArray(result)

	return
}

func prepareTestMessage() dnsmessage.Message {
	var header = dnsmessage.Header{ID: uint16(rand.Uint32()), OpCode: dnsmessage.OpCodeQuery, QueryCount: 1, RecursionDesired: true}
	var question = dnsmessage.Question{QName: "equabank.cz", QType: dnsmessage.ResourceRecordA, QClass: dnsmessage.ClassInternet}
	return dnsmessage.Message{Header: header, Questions: []dnsmessage.Question{question}}
}
