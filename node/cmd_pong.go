package node

import (
	"io"

	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/pkg/binary"
)

func (n Node) handlePong(header *networkprotocol.MessageHeader, conn io.ReadWriter) error {
	var pong networkprotocol.MsgPing

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&pong); err != nil {
		return err
	}

	n.PongCh <- pong.Nonce

	return nil
}
