package dnsmessage

import "encoding/binary"

// ResourceRecord is structure for dns query response see https://tools.ietf.org/html/rfc1035#section-4.1.3
type ResourceRecord struct {
	RName       Name
	RType       ResourceRecordType
	RClass      Class
	RtimeToLive uint32 //time to live in seconds
	//RDataLenght uint16 // length of data - is in message but i dont need, because slices have length
	RData []byte
}

// Serialize returns bytes ready to be send
func (r ResourceRecord) Serialize() []byte {
	recordData := r.RName.Serialize()

	//save length for slicing later
	dataLenght := len(recordData)
	recordData = append(recordData, make([]byte, 10)...)

	binary.BigEndian.PutUint16(recordData[dataLenght:dataLenght+2], uint16(r.RType))
	binary.BigEndian.PutUint16(recordData[dataLenght+2:dataLenght+4], uint16(r.RClass))
	binary.BigEndian.PutUint32(recordData[dataLenght+4:dataLenght+8], r.RtimeToLive)
	binary.BigEndian.PutUint16(recordData[dataLenght+8:], uint16(len(r.RData)))

	recordData = append(recordData, r.RData...)

	return recordData
}

// ResourceRecord reads bytes and attemts to construct ResourceRecord from starting index, returns ResourceRecord and total bytes read
func DeserializeResourceRecord(bytes []byte, startIndex uint32) (ResourceRecord, uint32, error) {
	var rr ResourceRecord
	var bytesRead uint32
	var err error
	rr.RName, bytesRead, err = DeserializeName(bytes, startIndex)
	if err != nil {
		return rr, bytesRead, err
	}

	rr.RType = ResourceRecordType(binary.BigEndian.Uint16(bytes[startIndex+bytesRead : startIndex+bytesRead+2]))
	bytesRead += 2
	rr.RClass = Class(binary.BigEndian.Uint16(bytes[startIndex+bytesRead : startIndex+bytesRead+2]))
	bytesRead += 2
	rr.RtimeToLive = binary.BigEndian.Uint32(bytes[startIndex+bytesRead : startIndex+bytesRead+4])
	bytesRead += 4
	RDataLenght := uint32(binary.BigEndian.Uint16(bytes[startIndex+bytesRead : startIndex+bytesRead+2]))
	bytesRead += 2
	rr.RData = bytes[startIndex+bytesRead : startIndex+bytesRead+RDataLenght]
	bytesRead += RDataLenght

	return rr, bytesRead, err
}

func (rr ResourceRecord) String() string {
	var text string
	text += formatLabelValue("RName", rr.RName) + "\n"
	text += formatLabelValue("RType", rr.RType) + "\n"
	text += formatLabelValue("RClass", rr.RClass) + "\n"
	text += formatLabelValue("RtimeToLive", rr.RtimeToLive) + "\n"
	text += formatLabelValue("RData", rr.RData)
	return text
}
