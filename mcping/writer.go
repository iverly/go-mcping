package mcping

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

func readResponse(conn *net.Conn) (string, error) {
	read := bufio.NewReader(*conn)
	binary.ReadUvarint(read)

	packetType, _ := read.ReadByte()
	if bytes.Compare([]byte{packetType}, []byte("\x00")) != 0 {
		return "", errors.New("error response packet type")
	}

	//Get data length via Varint
	length, err := binary.ReadUvarint(read)
	if err != nil {
		return "", err
	}
	if length < 10 {
		return "", errors.New("error to small response")
	} else if length > 700000 {
		return "", errors.New("error to big response")
	}

	//Recieve json buffer
	bytesRecieved := uint64(0)
	recBytes := make([]byte, length)
	for bytesRecieved < length {
		n, _ := read.Read(recBytes[bytesRecieved:length])
		bytesRecieved = bytesRecieved + uint64(n)
	}

	return string(recBytes), nil
}

func sendPacket(host string, port uint16, conn *net.Conn) {
	var dataBuf bytes.Buffer
	var finBuf bytes.Buffer

	writeProtocol(&dataBuf, "\x6D") // 1.9 protocol
	writeHost(&dataBuf, host)
	writePort(&dataBuf, port)
	dataBuf.Write([]byte("\x01")) // end of packet

	// Prepend packet length with data
	packetLength := []byte{uint8(dataBuf.Len())}
	finBuf.Write(append(packetLength, dataBuf.Bytes()...))

	// Sending packet
	(*conn).Write(finBuf.Bytes())
	(*conn).Write([]byte("\x01\x00"))
}

func writeProtocol(b *bytes.Buffer, protocol string) {
	b.Write([]byte("\x00")) // Packet ID
	b.Write([]byte(protocol))
}

func writeHost(b *bytes.Buffer, host string) {
	b.Write([]uint8{uint8(len(host))})
	b.Write([]byte(host))
}

func writePort(b *bytes.Buffer, port uint16) {
	a := make([]byte, 2)
	binary.BigEndian.PutUint16(a, port)
	b.Write(a)
}
