package node

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	networkprotocol "github.com/uwezukwechibuzor/bitcoin-node-handshake/network-protocol"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/pkg/binary"
)

const (
	pingIntervalSec = 120
	pingTimeoutSec  = 30
)

// Peer describes a peer node in a network.
type Peer struct {
	Address    net.Addr
	Connection io.ReadWriteCloser
	PongCh     chan uint64
	Services   uint64
	UserAgent  string
	Version    int32
}

// ID returns peer ID.
func (p Peer) ID() PeerID {
	return PeerID(p.Address.String())
}

func (p Peer) String() string {
	return fmt.Sprintf("%s (%s)", p.UserAgent, p.Address)
}

type peerPing struct {
	nonce  uint64
	peerID PeerID
}

func (n Node) monitorPeers() {
	peerPings := make(map[uint64]PeerID)

	for {
		select {
		case nonce := <-n.PongCh:
			peerID := peerPings[nonce]
			if peerID == "" {
				break
			}
			peer := n.Peers[peerID]
			if peer == nil {
				break
			}

			peer.PongCh <- nonce
			delete(peerPings, nonce)

		case pp := <-n.PingCh:
			peerPings[pp.nonce] = pp.peerID

		case peerID := <-n.DisconCh:
			n.disconnectPeer(peerID)

			for k, v := range peerPings {
				if v == peerID {
					delete(peerPings, k)
					break
				}
			}
		}
	}
}

func (n *Node) monitorPeer(peer *Peer) {
	for {
		time.Sleep(pingIntervalSec * time.Second)

		ping, nonce, err := networkprotocol.NewPingMsg(n.Network)
		if err != nil {
			logrus.Fatalf("monitorPeer, NewPingMsg: %v", err)
		}

		msg, err := binary.Marshal(ping)
		if err != nil {
			logrus.Fatalf("monitorPeer, binary.Marshal: %v", err)
		}

		if _, err := peer.Connection.Write(msg); err != nil {
			n.disconnectPeer(peer.ID())
		}

		logrus.Debugf("sent 'ping' to %s", peer)

		n.PingCh <- peerPing{
			nonce:  nonce,
			peerID: peer.ID(),
		}

		t := time.NewTimer(pingTimeoutSec * time.Second)

		select {
		case pn := <-peer.PongCh:
			if pn != nonce {
				logrus.Errorf("nonce doesn't match for %s: want %d, got %d", peer, nonce, pn)
				n.DisconCh <- peer.ID()
				return
			}
			logrus.Debugf("got 'pong' from %s", peer)
		case <-t.C:
			n.DisconCh <- peer.ID()
			return
		}

		t.Stop()
	}
}
