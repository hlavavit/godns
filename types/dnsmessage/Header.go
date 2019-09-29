package dnsmessage

import (
	"encoding/binary"
	"fmt"
)

// Header type according to https://tools.ietf.org/html/rfc1035#section-4.1.1
type Header struct {
	ID                            uint16
	Response                      bool
	OpCode                        OpCode
	AuthoritativeAnswer           bool
	TrunCation                    bool
	RecursionDesired              bool
	RecursionAvailable            bool
	ResponseCode                  ResponseCode
	QueryCount                    uint16
	AnswerCount                   uint16
	NameServerResouceRecordsCount uint16
	AdditionalRecordsCount        uint16
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
	if h.TrunCation {
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

// Serialize returns bytes ready to be send
func (h Header) Serialize() []byte {
	var bytes [12]byte
	binary.BigEndian.PutUint16(bytes[0:2], h.ID)
	binary.BigEndian.PutUint16(bytes[2:4], h.getFlag())
	binary.BigEndian.PutUint16(bytes[4:6], h.QueryCount)
	binary.BigEndian.PutUint16(bytes[6:8], h.AnswerCount)
	binary.BigEndian.PutUint16(bytes[8:10], h.NameServerResouceRecordsCount)
	binary.BigEndian.PutUint16(bytes[10:12], h.AdditionalRecordsCount)
	return bytes[:]
}

func (h Header) String() string {
	return fmt.Sprintf("Formated flag %016b\n", h.getFlag())
}
