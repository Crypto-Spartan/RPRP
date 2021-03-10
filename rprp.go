package rprp

// header
//		version  : version of the RPRP protocol used
// 	groupID  : nonce; random number generated for each group of packets sent together
// 	packetID : packet # within the group of packets
// 	fecType  : labels the packet as containing either data or parity
// 	numData  : number of data packets in the group
// 	numParity: number of parity packets in the group
type header struct {
	version   byte   // 6 bits \ These make
	fecType   byte   // 2 bits / one byte
	packetID  byte   // 1 byte
	groupID   uint32 // 4 bytes
	numData   byte   // 1 byte
	numParity byte   // 1 byte
	// total: 8 bytes
}
