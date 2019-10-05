package dnsmessage

import "fmt"

// ResponseCode for dns query responses see https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6
// should be max 4 bits
type ResponseCode uint16

// see https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-6
const (
	NoError   ResponseCode = 0
	FormError ResponseCode = 1
	ServFail  ResponseCode = 2
	NXDomain  ResponseCode = 3
	NotImp    ResponseCode = 4
	Refused   ResponseCode = 5
	YXDomain  ResponseCode = 6
	YXRRSet   ResponseCode = 7
	NXRRSet   ResponseCode = 8
	NotAuth   ResponseCode = 9
	NotZone   ResponseCode = 10
	DSOTYPENI ResponseCode = 11
	BADVERS   ResponseCode = 16
	BADSIG    ResponseCode = 16
	BADKEY    ResponseCode = 17
	BADTIME   ResponseCode = 18
	BADMODE   ResponseCode = 19
	BADNAME   ResponseCode = 20
	BADALG    ResponseCode = 21
	BADTRUNC  ResponseCode = 22
	BADCOOKIE ResponseCode = 23
)

func (rc ResponseCode) String() string {
	switch rc {
	case NoError:
		return "NoError"
	case FormError:
		return "FormError"
	case ServFail:
		return "ServFail"
	case NXDomain:
		return "NXDomain"
	case NotImp:
		return "NotImp"
	case Refused:
		return "Refused"
	case YXDomain:
		return "YXDomain"
	case YXRRSet:
		return "YXRRSet"
	case NXRRSet:
		return "NXRRSet"
	case NotAuth:
		return "NotAuth"
	case NotZone:
		return "NotZone"
	case DSOTYPENI:
		return "DSOTYPENI"
	case BADVERS:
		return "BADSIG"
	case BADKEY:
		return "BADKEY"
	case BADTIME:
		return "BADTIME"
	case BADMODE:
		return "BADMODE"
	case BADNAME:
		return "BADNAME"
	case BADALG:
		return "BADALG"
	case BADTRUNC:
		return "BADTRUNC"
	case BADCOOKIE:
		return "BADCOOKIE"

	}
	return fmt.Sprint(uint16(rc))
}
