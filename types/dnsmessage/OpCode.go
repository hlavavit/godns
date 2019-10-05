package dnsmessage

import (
	"fmt"
)

// OpCode for dns message header
type OpCode uint16

//see https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-5
const (
	OpCodeQuery             OpCode = 0
	OpCodeInverse           OpCode = 1
	OpCodeStatus            OpCode = 2
	OpCodeNotify            OpCode = 4
	OpCodeUpdate            OpCode = 5
	OpCodeStatefulOperation OpCode = 6
)

func (op OpCode) String() string {
	switch op {
	case OpCodeQuery:
		return "QUERY"
	case OpCodeInverse:
		return "IQUERY"
	case OpCodeStatus:
		return "STATUS"
	case OpCodeNotify:
		return "NOTIFY"
	case OpCodeUpdate:
		return "UPDATE"
	case OpCodeStatefulOperation:
		return "STATEFUL_OP"
	}
	return fmt.Sprint(uint16(op))
}
