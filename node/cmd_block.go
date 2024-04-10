package node

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/pkg/binary"
)

func (no Node) handleBlock(header *networkprotocol.MessageHeader, conn io.ReadWriter) error {
	var block networkprotocol.MsgBlock

	// LimitReader returns a Reader that reads from r but stops with EOF after n bytes.
	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&block); err != nil {
		return err
	}

	// Calculate the hash of the block
	hash, err := block.Hash()
	if err != nil {
		return fmt.Errorf("block.Hash: %+v", err)
	}

	logrus.Debugf("block: %x", hash)

	// check that the block hash is valid
	if err := block.Verify(); err != nil {
		return fmt.Errorf("rejected invalid block %x", hash)
	}

	// Send the valid block to the mempool
	no.mempool.NewBlockCh <- block

	return nil
}
