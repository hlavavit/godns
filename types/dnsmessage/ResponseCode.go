package dnsmessage

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
