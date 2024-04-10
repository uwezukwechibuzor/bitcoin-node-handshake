package node

import (
	"encoding/hex"

	"github.com/sirupsen/logrus"
	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
)

// Mempool represents mempool.
type Mempool struct {
	NewBlockCh chan networkprotocol.MsgBlock
	NewTxCh    chan networkprotocol.MsgTx

	txs map[string]*networkprotocol.MsgTx
}

// NewMempool returns a new Mempool.
func NewMempool() *Mempool {
	return &Mempool{
		NewBlockCh: make(chan networkprotocol.MsgBlock),
		NewTxCh:    make(chan networkprotocol.MsgTx),
		txs:        make(map[string]*networkprotocol.MsgTx),
	}
}

// Run starts mempool state handling.
func (m Mempool) Run() {
	for {
		select {
		case tx := <-m.NewTxCh:
			hash, err := tx.Hash()
			if err != nil {
				logrus.Errorf("failed to calculate tx hash: %+v", err)
				break
			}

			txid := hex.EncodeToString(hash)
			m.txs[txid] = &tx
		case block := <-m.NewBlockCh:
			for _, tx := range block.Txs {
				hash, err := tx.Hash()
				if err != nil {
					logrus.Errorf("failed to calculate tx hash: %+v", err)
					break
				}

				txid := hex.EncodeToString(hash)
				delete(m.txs, txid)
			}
		}
	}
}

// Mempool
func (n Node) Mempool() map[string]*networkprotocol.MsgTx {
	m := make(map[string]*networkprotocol.MsgTx)

	for k, v := range n.mempool.txs {
		m[string(k)] = v
	}

	return m
}
