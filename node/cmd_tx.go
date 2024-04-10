package node

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/pkg/binary"
)

func (no Node) handleTx(header *networkprotocol.MessageHeader, conn io.ReadWriter) error {
	var tx networkprotocol.MsgTx

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&tx); err != nil {
		return err
	}

	hash, err := tx.Hash()
	if err != nil {
		return fmt.Errorf("tx.Hash: %+v", err)
	}

	logrus.Debugf("transaction: %x", hash)

	if err := tx.Verify(); err != nil {
		return fmt.Errorf("rejected invalid transaction %x", hash)
	}

	no.mempool.NewTxCh <- tx

	return nil
}
