package dnsmessage

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

// Name is extended string adding extra methods for serialization and deserialization
type Name string

func (n Name) getLabels() []string {
	return strings.Split(string(n), ".")
}

func encodeLabel(label string) []byte {
	buff := make([]byte, len(label)+1)
	buff[0] = byte(len(label))
	copy(buff[1:], []byte(label))
	return buff
}

// Serialize returns bytes ready to be send
func (n Name) Serialize() []byte {
	nameData := make([]byte, 0, len(n)+5)
	for _, label := range n.getLabels() {
		nameData = append(nameData, encodeLabel(label)...)
	}
	//terminate labels
	nameData = append(nameData, byte(0))

	return nameData
}

func test() {

}

// getPointerValue checks slice for pointer, returns -1 if value not a pointer, pointer value othervise. if slice is not big enough returns error
func getPointerValue(bytes []byte) (int, error) {
	pointer := -1
	if len(bytes) < 2 {
		return pointer, errors.New(fmt.Sprint("bytes too short to deserialize dnsmessage.Header, expected at least 12 bytes but got", len(bytes)))
	}

	var pointerMask byte = 0b11000000
	if bytes[0]&pointerMask > 0 {
		pointerBytes := make([]byte, 2)
		copy(pointerBytes, bytes)
		//Flip masked bites to not include them in pointer val
		pointerBytes[0] = pointerBytes[0] & (^pointerMask)
		// unsignes16 to signed int first to preseve the signed bit, then to int.
		pointer = int(binary.BigEndian.Uint16(pointerBytes))
	}

	return pointer, nil
}

func DeserializeName(bytes []byte, startIndex uint32, visitedIdexes ...uint32) (Name, uint32, error) {
	var name Name = ""
	var read uint32

	for _, testIndex := range visitedIdexes {
		if testIndex == startIndex {
			return name, read, errors.New("Trying to visit already visited index, continuing would cause endless cycle")
		}
	}

	pointer, err := getPointerValue(bytes[startIndex : startIndex+2])
	if err != nil {
		return name, read, err
	}

	if pointer >= 0 {
		namePart, _, err := DeserializeName(bytes, uint32(pointer), append(visitedIdexes, startIndex)...)
		if err != nil {
			return name, read, err
		}
		read = 2 // pointer is always two bytes
		name = namePart
		//Pointer is basicly an end
		return name, read, nil
	}

	//not a pointer so current bit must be lengt of next label or zero for end
	labelLength := uint32(bytes[startIndex])
	if labelLength == 0 {
		//0 byte - end of name
		read = 1 //read one byte for length
		return name, read, nil
	}

	name = Name(bytes[startIndex+1 : startIndex+1+labelLength])
	read = labelLength + 1 //length of string +1 for length bit

	nextLabel, nextRead, err := DeserializeName(bytes, startIndex+1+labelLength, append(visitedIdexes, startIndex)...)

	if err != nil {
		return name, read, err
	}

	if len(nextLabel) > 0 {
		name += "." + nextLabel
	}

	read += nextRead

	return name, read, nil
}
