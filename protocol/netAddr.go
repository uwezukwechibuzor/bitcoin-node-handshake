package protocol

import (
	"fmt"

	types "github.com/uwezukwechibuzor/bitcoin-node-handshake/types/proto"
)

// NewIPv4 ...
func NewIPv4(ip []byte) *types.IPv4 {
	return &types.IPv4{Ip: ip}
}

// ToIPv6 converts an IPv4 address to a mapped IPv6 address
func ToIPv6(ip *types.IPv4) ([16]byte, error) {
	if ip == nil {
		return [16]byte{}, fmt.Errorf("nil IPv4 pointer provided")
	}
	var v6 [16]byte
	copy(v6[10:], ip.Ip[:])
	return v6, nil
}
