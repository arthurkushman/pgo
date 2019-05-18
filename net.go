package pgo

import (
	"encoding/binary"
	"errors"
	"net"
)

// IP2long converts a string containing an (IPv4) Internet Protocol dotted address into a long integer
func IP2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

// Long2ip converts an long integer address into a string in (IPv4) Internet standard dotted format
func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}
