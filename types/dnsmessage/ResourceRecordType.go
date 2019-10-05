package dnsmessage

import "fmt"

// ResourceRecordType for dns
type ResourceRecordType uint16

//see https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4
const (
	ResourceRecordA          ResourceRecordType = 1  //a host address
	ResourceRecordNS         ResourceRecordType = 2  //an authoritative name server
	ResourceRecordMD         ResourceRecordType = 3  //a mail destination (OBSOLETE - use MX)
	ResourceRecordMF         ResourceRecordType = 4  //a mail forwarder (OBSOLETE - use MX)
	ResourceRecordCNAME      ResourceRecordType = 5  //the canonical name for an alias
	ResourceRecordSOA        ResourceRecordType = 6  //marks the start of a zone of authority
	ResourceRecordMB         ResourceRecordType = 7  //a mailbox domain name (EXPERIMENTAL)
	ResourceRecordMG         ResourceRecordType = 8  //a mail group member (EXPERIMENTAL)
	ResourceRecordMR         ResourceRecordType = 9  //a mail rename domain name (EXPERIMENTAL)
	ResourceRecordNULL       ResourceRecordType = 10 //a null RR (EXPERIMENTAL)
	ResourceRecordWKS        ResourceRecordType = 11 //a well known service description
	ResourceRecordPTR        ResourceRecordType = 12 //a domain name pointer
	ResourceRecordHINFO      ResourceRecordType = 13 //host information
	ResourceRecordMINFO      ResourceRecordType = 14 //mailbox or mail list information
	ResourceRecordMX         ResourceRecordType = 15 //mail exchange
	ResourceRecordTXT        ResourceRecordType = 16 //text strings
	ResourceRecordRP         ResourceRecordType = 17 //for Responsible Person
	ResourceRecordAFSDB      ResourceRecordType = 18 //for AFS Data Base location
	ResourceRecordX25        ResourceRecordType = 19 //for X.25 PSDN address
	ResourceRecordISDN       ResourceRecordType = 20 //for ISDN address
	ResourceRecordRT         ResourceRecordType = 21 //for Route Through
	ResourceRecordNSAP       ResourceRecordType = 22 //for NSAP address, NSAP style A record
	ResourceRecordNSAPPTR    ResourceRecordType = 23 //for domain name pointer, NSAP style
	ResourceRecordSIG        ResourceRecordType = 24 //for security signature
	ResourceRecordKEY        ResourceRecordType = 25 //for security key
	ResourceRecordPX         ResourceRecordType = 26 //X.400 mail mapping information
	ResourceRecordGPOS       ResourceRecordType = 27 //Geographical Position
	ResourceRecordAAAA       ResourceRecordType = 28 //IP6 Address
	ResourceRecordLOC        ResourceRecordType = 29 //Location Information
	ResourceRecordNXT        ResourceRecordType = 30 //Next Domain (OBSOLETE)
	ResourceRecordEID        ResourceRecordType = 31 //Endpoint Identifier
	ResourceRecordNIMLOC     ResourceRecordType = 32 //Nimrod Locator
	ResourceRecordSRV        ResourceRecordType = 33 //Server Selection
	ResourceRecordATMA       ResourceRecordType = 34 //ATM Address
	ResourceRecordNAPTR      ResourceRecordType = 35 //Naming Authority Pointer
	ResourceRecordKX         ResourceRecordType = 36 //Key Exchanger
	ResourceRecordCERT       ResourceRecordType = 37 //CERT
	ResourceRecordA6         ResourceRecordType = 38 //A6 (OBSOLETE - use AAAA)
	ResourceRecordDNAME      ResourceRecordType = 39 //DNAME
	ResourceRecordSINK       ResourceRecordType = 40 //SINK
	ResourceRecordOPT        ResourceRecordType = 41 //OPT
	ResourceRecordAPL        ResourceRecordType = 42 //APL
	ResourceRecordDS         ResourceRecordType = 43 //Delegation Signer
	ResourceRecordSSHFP      ResourceRecordType = 44 //SSH Key Fingerprint
	ResourceRecordIPSECKEY   ResourceRecordType = 45 //IPSECKEY
	ResourceRecordRRSIG      ResourceRecordType = 46 //RRSIG
	ResourceRecordNSEC       ResourceRecordType = 47 //NSEC
	ResourceRecordDNSKEY     ResourceRecordType = 48 //DNSKEY
	ResourceRecordDHCID      ResourceRecordType = 49 //DHCID
	ResourceRecordNSEC3      ResourceRecordType = 50 //NSEC3
	ResourceRecordNSEC3PARAM ResourceRecordType = 51 //NSEC3PARAM
	ResourceRecordTLSA       ResourceRecordType = 52 //TLSA
	ResourceRecordSMIMEA     ResourceRecordType = 53 //S/MIME cert association
	ResourceRecordHIP        ResourceRecordType = 55 //Host Identity Protocol
	ResourceRecordNINFO      ResourceRecordType = 56 //NINFO
	ResourceRecordRKEY       ResourceRecordType = 57 //RKEY
	ResourceRecordTALINK     ResourceRecordType = 58 //Trust Anchor LINK
	ResourceRecordCDS        ResourceRecordType = 59 //Child DS
	ResourceRecordCDNSKEY    ResourceRecordType = 60 //DNSKEY(s) the Child wants reflected in DS
	ResourceRecordOPENPGPKEY ResourceRecordType = 61 //OpenPGP Key
	ResourceRecordCSYNC      ResourceRecordType = 62 //Child-To-Parent Synchronization
	ResourceRecordZONEMD     ResourceRecordType = 63 //message digest for DNS zone
	ResourceRecordSPF        ResourceRecordType = 99
	ResourceRecordUINFO      ResourceRecordType = 100
	ResourceRecordUID        ResourceRecordType = 101
	ResourceRecordGID        ResourceRecordType = 102
	ResourceRecordUNSPEC     ResourceRecordType = 103
	ResourceRecordNID        ResourceRecordType = 104
	ResourceRecordL32        ResourceRecordType = 105
	ResourceRecordL64        ResourceRecordType = 106
	ResourceRecordLP         ResourceRecordType = 107
	ResourceRecordEUI48      ResourceRecordType = 108   //an EUI-48 address
	ResourceRecordEUI64      ResourceRecordType = 109   //an EUI-64 address
	ResourceRecordTKEY       ResourceRecordType = 249   //Transaction Key
	ResourceRecordTSIG       ResourceRecordType = 250   //Transaction Signature
	ResourceRecordIXFR       ResourceRecordType = 251   //incremental transfer
	ResourceRecordAXFR       ResourceRecordType = 252   //transfer of an entire zone
	ResourceRecordMAILB      ResourceRecordType = 253   //mailbox-related RRs (MB, MG or MR)
	ResourceRecordMAILA      ResourceRecordType = 254   //mail agent RRs (OBSOLETE - see MX)
	ResourceRecordALL        ResourceRecordType = 255   //A request for some or all records the server has available
	ResourceRecordURI        ResourceRecordType = 256   //URI
	ResourceRecordCAA        ResourceRecordType = 257   //Certification Authority Restriction
	ResourceRecordAVC        ResourceRecordType = 258   //Application Visibility and Control
	ResourceRecordDOA        ResourceRecordType = 259   //Digital Object Architecture
	ResourceRecordAMTRELAY   ResourceRecordType = 260   //Automatic Multicast Tunneling Relay
	ResourceRecordTA         ResourceRecordType = 32768 //DNSSEC Trust Authorities
	ResourceRecordDLV        ResourceRecordType = 32769 //DNSSEC Lookaside Validation
)

func (rr ResourceRecordType) String() string {
	switch rr {
	case ResourceRecordA:
		return "A"
	case ResourceRecordNS:
		return "NS"
	case ResourceRecordMD:
		return "MD"
	case ResourceRecordMF:
		return "MF"
	case ResourceRecordCNAME:
		return "CNAME"
	case ResourceRecordSOA:
		return "SOA"
	case ResourceRecordMB:
		return "MB"
	case ResourceRecordMG:
		return "MG"
	case ResourceRecordMR:
		return "MR"
	case ResourceRecordNULL:
		return "NULL"
	case ResourceRecordWKS:
		return "WKS"
	case ResourceRecordPTR:
		return "PTR"
	case ResourceRecordHINFO:
		return "HINFO"
	case ResourceRecordMINFO:
		return "MINFO"
	case ResourceRecordMX:
		return "MX"
	case ResourceRecordTXT:
		return "TXT"
	case ResourceRecordRP:
		return "RP"
	case ResourceRecordAFSDB:
		return "AFSDB"
	case ResourceRecordX25:
		return "X25"
	case ResourceRecordISDN:
		return "ISDN"
	case ResourceRecordRT:
		return "RT"
	case ResourceRecordNSAP:
		return "NSAP"
	case ResourceRecordNSAPPTR:
		return "NSAPPTR"
	case ResourceRecordSIG:
		return "SIG"
	case ResourceRecordKEY:
		return "KEY"
	case ResourceRecordPX:
		return "PX"
	case ResourceRecordGPOS:
		return "GPOS"
	case ResourceRecordAAAA:
		return "AAAA"
	case ResourceRecordLOC:
		return "LOC"
	case ResourceRecordNXT:
		return "NXT"
	case ResourceRecordEID:
		return "EID"
	case ResourceRecordNIMLOC:
		return "NIMLOC"
	case ResourceRecordSRV:
		return "SRV"
	case ResourceRecordATMA:
		return "ATMA"
	case ResourceRecordNAPTR:
		return "NAPTR"
	case ResourceRecordKX:
		return "KX"
	case ResourceRecordCERT:
		return "CERT"
	case ResourceRecordA6:
		return "A6"
	case ResourceRecordDNAME:
		return "DNAME"
	case ResourceRecordSINK:
		return "SINK"
	case ResourceRecordOPT:
		return "OPT"
	case ResourceRecordAPL:
		return "APL"
	case ResourceRecordDS:
		return "DS"
	case ResourceRecordSSHFP:
		return "SSHFP"
	case ResourceRecordIPSECKEY:
		return "IPSECKEY"
	case ResourceRecordRRSIG:
		return "RRSIG"
	case ResourceRecordNSEC:
		return "NSEC"
	case ResourceRecordDNSKEY:
		return "DNSKEY"
	case ResourceRecordDHCID:
		return "DHCID"
	case ResourceRecordNSEC3:
		return "NSEC3"
	case ResourceRecordNSEC3PARAM:
		return "NSEC3PARAM"
	case ResourceRecordTLSA:
		return "TLSA"
	case ResourceRecordSMIMEA:
		return "SMIMEA"
	case ResourceRecordHIP:
		return "HIP"
	case ResourceRecordNINFO:
		return "NINFO"
	case ResourceRecordRKEY:
		return "RKEY"
	case ResourceRecordTALINK:
		return "TALINK"
	case ResourceRecordCDS:
		return "CDS"
	case ResourceRecordCDNSKEY:
		return "CDNSKEY"
	case ResourceRecordOPENPGPKEY:
		return "OPENPGPKEY"
	case ResourceRecordCSYNC:
		return "CSYNC"
	case ResourceRecordZONEMD:
		return "ZONEMD"
	case ResourceRecordSPF:
		return "SPF"
	case ResourceRecordUINFO:
		return "UINFO"
	case ResourceRecordUID:
		return "UID"
	case ResourceRecordGID:
		return "GID"
	case ResourceRecordUNSPEC:
		return "UNSPEC"
	case ResourceRecordNID:
		return "NID"
	case ResourceRecordL32:
		return "L32"
	case ResourceRecordL64:
		return "L64"
	case ResourceRecordLP:
		return "LP"
	case ResourceRecordEUI48:
		return "EUI48"
	case ResourceRecordEUI64:
		return "EUI64"
	case ResourceRecordTKEY:
		return "TKEY"
	case ResourceRecordTSIG:
		return "TSIG"
	case ResourceRecordIXFR:
		return "IXFR"
	case ResourceRecordAXFR:
		return "AXFR"
	case ResourceRecordMAILB:
		return "MAILB"
	case ResourceRecordMAILA:
		return "MAILA"
	case ResourceRecordALL:
		return "ALL"
	case ResourceRecordURI:
		return "URI"
	case ResourceRecordCAA:
		return "CAA"
	case ResourceRecordAVC:
		return "AVC"
	case ResourceRecordDOA:
		return "DOA"
	case ResourceRecordAMTRELAY:
		return "AMTRELAY"
	case ResourceRecordTA:
		return "TA"
	case ResourceRecordDLV:
		return "DLV"

	}
	return fmt.Sprint(uint16(rr))
}
