package dnsmessage

import (
	"fmt"
	"strings"
)

// Message is type for sending and recveiving dns messages
// details https://tools.ietf.org/html/rfc1035#section-4
type Message struct {
	Header            Header
	Questions         []Question
	AnswerRecords     []ResourceRecord
	AuthorityRecords  []ResourceRecord
	AdditionalRecords []ResourceRecord
}

// Serialize returns bytes ready to be send
func (m Message) Serialize() []byte {
	buff := make([]byte, 0, 512)

	// fix header counts before serialization
	m.Header.QuestionCount = uint16(len(m.Questions))
	m.Header.AnswerRecordCount = uint16(len(m.AnswerRecords))
	m.Header.AuthorityRecordCount = uint16(len(m.AuthorityRecords))
	m.Header.AdditionalRecordCount = uint16(len(m.AdditionalRecords))

	buff = append(buff, m.Header.Serialize()...)
	for _, question := range m.Questions {
		buff = append(buff, question.Serialize()...)
	}
	for _, rr := range m.AnswerRecords {
		buff = append(buff, rr.Serialize()...)
	}
	for _, rr := range m.AuthorityRecords {
		buff = append(buff, rr.Serialize()...)
	}
	for _, rr := range m.AdditionalRecords {
		buff = append(buff, rr.Serialize()...)
	}
	return buff
}

// DeserializeMessage reads bytes and attemts to construct message
func DeserializeMessage(data []byte) (Message, error) {
	var message Message
	var err error
	message.Header, err = DeserializeHeader(data[0:12])
	var read uint32 = 12
	for i := uint16(0); i < message.Header.QuestionCount; i++ {
		question, readBytes, err := DeserializeQuestion(data, read)
		if err != nil {
			return message, err
		}
		read += readBytes
		message.Questions = append(message.Questions, question)
	}
	for i := uint16(0); i < message.Header.AnswerRecordCount; i++ {
		rr, readBytes, err := DeserializeResourceRecord(data, read)
		if err != nil {
			return message, err
		}
		read += readBytes
		message.AnswerRecords = append(message.AnswerRecords, rr)
	}
	for i := uint16(0); i < message.Header.AuthorityRecordCount; i++ {
		rr, readBytes, err := DeserializeResourceRecord(data, read)
		if err != nil {
			return message, err
		}
		read += readBytes
		message.AuthorityRecords = append(message.AuthorityRecords, rr)
	}
	for i := uint16(0); i < message.Header.AdditionalRecordCount; i++ {
		rr, readBytes, err := DeserializeResourceRecord(data, read)
		if err != nil {
			return message, err
		}
		read += readBytes
		message.AdditionalRecords = append(message.AdditionalRecords, rr)
	}
	return message, err
}

func formatArray(slice []interface{}) string {
	var text string
	for i, record := range slice {
		recordLines := strings.Split(fmt.Sprint(record), "\n")
		for _, line := range recordLines {
			text += fmt.Sprintf("| %-50v |\n", line)
		}
		if i+1 != len(slice) {
			text += "|////////////////////////////////////////////////////|\n"
		}
	}
	return text
}

func formatResourceRecords(slice []ResourceRecord) string {
	interfaceSlice := make([]interface{}, len(slice))
	for i, rr := range slice {
		interfaceSlice[i] = rr
	}
	return formatArray(interfaceSlice)
}
func formatQuestions(slice []Question) string {
	interfaceSlice := make([]interface{}, len(slice))
	for i, rr := range slice {
		interfaceSlice[i] = rr
	}
	return formatArray(interfaceSlice)
}

func (m Message) String() string {
	text := "|//////////////////// DNS Header ////////////////////|\n"
	headerLines := strings.Split(m.Header.String(), "\n")
	for _, line := range headerLines {
		text += fmt.Sprintf("| %-50v |\n", line)

	}
	if len(m.Questions) > 0 {
		text += "|////////////////// DNS Questions ///////////////////|\n"
		text += formatQuestions(m.Questions)
	}
	if len(m.AnswerRecords) > 0 {
		text += "|//////////////// DNS Answer Records ////////////////|\n"
		text += formatResourceRecords(m.AnswerRecords)
	}
	if len(m.AuthorityRecords) > 0 {
		text += "|////////////// DNS Authority Records ///////////////|\n"
		text += formatResourceRecords(m.AuthorityRecords)
	}
	if len(m.AdditionalRecords) > 0 {
		text += "|////////////// DNS Additional Records //////////////|\n"
		text += formatResourceRecords(m.AdditionalRecords)
	}
	text += "|////////////////////////////////////////////////////|\n"

	return text
}
