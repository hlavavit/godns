package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"godns/types/dnsmessage"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Started godns with arguments=", os.Args[1:])

	message, err := sendMessage("8.8.8.8:53", prepareTestMessage())
	//message, err := sendMessage("1.1.1.1:53", prepareTestMessage())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(message)
}

func sendMessage(ipAddress string, message dnsmessage.Message) (dnsmessage.Message, error) {
	var responseMessage dnsmessage.Message

	conn, err := net.Dial("udp", ipAddress)
	if err != nil {
		return responseMessage, err
	}
	defer conn.Close()

	_, err = conn.Write(message.Serialize())
	if err != nil {
		return responseMessage, err
	}

	deadline := time.Now().Add(time.Duration(5) * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return responseMessage, err
	}

	buffer := make([]byte, 1024)
	readCount, err := conn.Read(buffer)
	if err != nil {
		return responseMessage, err
	}
	result := buffer[0:readCount]

	responseMessage, err = dnsmessage.DeserializeMessage(result)
	if err != nil {
		return responseMessage, err
	}

	return responseMessage, nil
}

func prepareTestMessage() dnsmessage.Message {
	random := make([]byte, 2)
	rand.Read(random)
	var header = dnsmessage.Header{ID: binary.BigEndian.Uint16(random), OpCode: dnsmessage.OpCodeQuery, QuestionCount: 0, RecursionDesired: true}
	var question1 = dnsmessage.Question{QName: "google.com", QType: dnsmessage.ResourceRecordA, QClass: dnsmessage.ClassInternet}
	//var question2 = dnsmessage.Question{QName: "google.com", QType: dnsmessage.ResourceRecordA, QClass: dnsmessage.ClassInternet}
	return dnsmessage.Message{Header: header, Questions: []dnsmessage.Question{question1}}
}
