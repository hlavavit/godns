package dnsmessage

import (
	"encoding/binary"
)

// Question is structure for dns query https://tools.ietf.org/html/rfc1035#section-4.1.2
type Question struct {
	QName  Name
	QType  ResourceRecordType
	QClass Class
}

// Serialize returns bytes ready to be send
func (q Question) Serialize() []byte {
	question := q.QName.Serialize()

	//save length for slicing later
	dataLenght := len(question)
	question = append(question, make([]byte, 4)...)

	binary.BigEndian.PutUint16(question[dataLenght:dataLenght+2], uint16(q.QType))
	binary.BigEndian.PutUint16(question[dataLenght+2:], uint16(q.QClass))

	return question
}

// DeserializeQuestion reads bytes and attemts to construct Question from starting index, returns question and total bytes read
func DeserializeQuestion(bytes []byte, startIndex uint32) (Question, uint32, error) {
	var q Question
	var bytesRead uint32
	var err error
	q.QName, bytesRead, err = DeserializeName(bytes, startIndex)
	if err != nil {
		return q, bytesRead, err
	}

	q.QType = ResourceRecordType(binary.BigEndian.Uint16(bytes[startIndex+bytesRead : startIndex+bytesRead+2]))
	bytesRead += 2
	q.QClass = Class(binary.BigEndian.Uint16(bytes[startIndex+bytesRead : startIndex+bytesRead+2]))
	bytesRead += 2

	return q, bytesRead, err
}

func (q Question) String() string {
	var text string
	text += formatLabelValue("QName", q.QName) + "\n"
	text += formatLabelValue("QType", q.QType) + "\n"
	text += formatLabelValue("QClass", q.QClass)
	return text
}
