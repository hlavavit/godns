package dnsmessage

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Header type according to https://tools.ietf.org/html/rfc1035#section-4.1.1
type Header struct {
	ID                    uint16
	Response              bool
	OpCode                OpCode
	AuthoritativeAnswer   bool
	Truncation            bool
	RecursionDesired      bool
	RecursionAvailable    bool
	ResponseCode          ResponseCode
	QuestionCount         uint16
	AnswerRecordCount     uint16
	AuthorityRecordCount  uint16
	AdditionalRecordCount uint16
}

func (h Header) getFlag() uint16 {
	var flag uint16

	if h.Response {
		flag |= 1 << 15
	}
	flag |= uint16(h.OpCode) << 11
	if h.AuthoritativeAnswer {
		flag |= 1 << 10
	}
	if h.Truncation {
		flag |= 1 << 9
	}
	if h.RecursionDesired {
		flag |= 1 << 8
	}
	if h.RecursionAvailable {
		flag |= 1 << 7
	}
	flag |= uint16(h.ResponseCode)
	return flag
}

func hasBit(n uint16, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// flag is two bytes of data
// byte shift position: |    15    | 14 13 12 11 | 10 | 09 | 08 | 07 |  06 05 04  |  03 02 01 00 |
// byte meaning:        | response |    opcode   | AA | TC | RD | RD | unassigned | responseCode |
func (h *Header) parseFlag(flag uint16) {
	h.Response = flag&(1<<15) > 0

	opCodeValue := flag & (0b1111 << 11) >> 11 // get just the four at elenth position
	h.OpCode = OpCode(opCodeValue)

	h.AuthoritativeAnswer = flag&(1<<10) > 0
	h.Truncation = flag&(1<<9) > 0
	h.RecursionDesired = flag&(1<<8) > 0
	h.RecursionAvailable = flag&(1<<7) > 0

	responseCodeValue := flag & (0b1111) // get just the four
	h.ResponseCode = ResponseCode(responseCodeValue)

}

// Serialize returns bytes ready to be send
func (h Header) Serialize() []byte {
	var bytes [12]byte
	binary.BigEndian.PutUint16(bytes[0:2], h.ID)
	binary.BigEndian.PutUint16(bytes[2:4], h.getFlag())
	binary.BigEndian.PutUint16(bytes[4:6], h.QuestionCount)
	binary.BigEndian.PutUint16(bytes[6:8], h.AnswerRecordCount)
	binary.BigEndian.PutUint16(bytes[8:10], h.AuthorityRecordCount)
	binary.BigEndian.PutUint16(bytes[10:12], h.AdditionalRecordCount)
	return bytes[:]
}

// DeserializeHeader reads bytes and attemts to construct header
func DeserializeHeader(bytes []byte) (Header, error) {
	var header Header
	if len(bytes) < 12 {
		return header, errors.New(fmt.Sprint("bytes too short to deserialize dnsmessage.Header, expected at least 12 bytes but got", len(bytes)))
	}
	header.ID = binary.BigEndian.Uint16(bytes[0:2])
	header.parseFlag(binary.BigEndian.Uint16(bytes[2:4]))
	header.QuestionCount = binary.BigEndian.Uint16(bytes[4:6])
	header.AnswerRecordCount = binary.BigEndian.Uint16(bytes[6:8])
	header.AuthorityRecordCount = binary.BigEndian.Uint16(bytes[8:10])
	header.AdditionalRecordCount = binary.BigEndian.Uint16(bytes[10:12])
	return header, nil
}

func (h Header) String() string {
	var text string
	text += formatLabelValue("ID", h.ID) + "\n"
	text += formatLabelValue("Response", h.Response) + "\n"
	text += formatLabelValue("OpCode", h.OpCode) + "\n"
	text += formatLabelValue("AuthoritativeAnswer", h.AuthoritativeAnswer) + "\n"
	text += formatLabelValue("Truncation", h.Truncation) + "\n"
	text += formatLabelValue("RecursionDesired", h.RecursionDesired) + "\n"
	text += formatLabelValue("RecursionAvailable", h.RecursionAvailable) + "\n"
	text += formatLabelValue("ResponseCode", h.ResponseCode) + "\n"
	text += formatLabelValue("QuestionCount", h.QuestionCount) + "\n"
	text += formatLabelValue("AnswerRecordCount", h.AnswerRecordCount) + "\n"
	text += formatLabelValue("AuthorityRecordCount", h.AuthorityRecordCount) + "\n"
	text += formatLabelValue("AdditionalRecordCount", h.AdditionalRecordCount)
	return text
}
