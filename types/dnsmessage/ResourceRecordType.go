package dnsmessage

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
