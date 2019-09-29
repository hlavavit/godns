package dnsmessage

// Message is type for sending and recveiving dns messages
// details https://tools.ietf.org/html/rfc1035#section-4
type Message struct {
	Header    Header
	Questions []Question
}

// Serialize returns bytes ready to be send
func (m Message) Serialize() []byte {
	buff := make([]byte, 0, 512)
	buff = append(buff, m.Header.Serialize()...)
	for _, question := range m.Questions {
		buff = append(buff, question.Serialize()...)
	}
	return buff
}
