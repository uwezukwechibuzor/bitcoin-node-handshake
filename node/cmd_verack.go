package node

import (
	"io"

	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
)

func (n Node) handleVerack(header *networkprotocol.MessageHeader, conn io.ReadWriter) error {
	return nil
}
