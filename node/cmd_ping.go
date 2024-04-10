package node

import (
	"io"

	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/pkg/binary"
)

func (n Node) handlePing(header *networkprotocol.MessageHeader, conn io.ReadWriter) error {
	var ping networkprotocol.MsgPing

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&ping); err != nil {
		return err
	}

	pong, err := networkprotocol.NewPongMsg(n.Network, ping.Nonce)
	if err != nil {
		return err
	}

	msg, err := binary.Marshal(pong)
	if err != nil {
		return err
	}

	if _, err := conn.Write(msg); err != nil {
		return err
	}

	return nil
}
