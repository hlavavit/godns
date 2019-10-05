package dnsmessage

import "fmt"

// Class - query class for dns message
type Class uint16

//see https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-2
const (
	ClassInternet Class = 1
	ClassChaos    Class = 3
	ClassHesiod   Class = 4
)

func (c Class) String() string {
	switch c {
	case ClassInternet:
		return "Internet"
	case ClassChaos:
		return "Chaos"
	case ClassHesiod:
		return "Hesiod"
	}
	return fmt.Sprint(uint16(c))
}
