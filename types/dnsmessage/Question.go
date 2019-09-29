package dnsmessage

import (
	"encoding/binary"
	"strings"
)

// Question is structure for dns query https://tools.ietf.org/html/rfc1035#section-4.1.2
type Question struct {
	QName  string
	QType  ResourceRecordType
	QClass Class
}

func (q Question) getLabels() []string {
	return strings.Split(q.QName, ".")
}

func encodeLabel(label string) []byte {
	buff := make([]byte, len(label)+1)
	buff[0] = byte(len(label))
	copy(buff[1:], []byte(label))
	return buff
}

// Serialize returns bytes ready to be send
func (q Question) Serialize() []byte {
	question := make([]byte, 0)
	for _, label := range q.getLabels() {
		question = append(question, encodeLabel(label)...)
	}
	//terminate labels
	question = append(question, byte(0))

	//save length for slicing later
	dataLenght := len(question)
	question = append(question, make([]byte, 4)...)

	binary.BigEndian.PutUint16(question[dataLenght:dataLenght+2], uint16(q.QType))
	binary.BigEndian.PutUint16(question[dataLenght+2:], uint16(q.QClass))

	return question
}
