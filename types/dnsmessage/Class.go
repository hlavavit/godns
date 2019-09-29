package dnsmessage

// Class - query class for dns message
type Class uint16

//see https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-2
const (
	ClassInternet Class = 1
	ClassChaos    Class = 3
	ClassHesiod   Class = 4
)
