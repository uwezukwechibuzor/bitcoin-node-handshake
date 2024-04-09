package cmd

import (
	"io"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/protocol"
	types "github.com/uwezukwechibuzor/bitcoin-node-handshake/types/proto"
	"google.golang.org/protobuf/proto"
)

func init() {
	btcNodeCmd.Flags().String("node-url", "127.0.0.1:9333", "TCP address of a Bitcoin node to connect to")
	btcNodeCmd.Flags().String("network", "simnet", "Bitcoin network (simnet, mainnet)")
}

var btcNodeCmd = &cobra.Command{
	Use: "btcNode",
	Run: func(cmd *cobra.Command, args []string) {
		nodeURL, err := cmd.Flags().GetString("node-url")
		if err != nil {
			logrus.Fatalln(err)
		}
		network, err := cmd.Flags().GetString("network")
		if err != nil {
			logrus.Fatalln(err)
		}

		version := &protocol.MsgVersion{
			Version:   protocol.Version,
			Services:  protocol.SrvNodeNetwork,
			Timestamp: time.Now().UTC().Unix(),
			AddrRecv: &types.NetAddr{
				Services: protocol.SrvNodeNetwork,
				Ip:       &types.IPv4{Ip: []byte{127, 0, 0, 1}},
				Port:     9333,
			},
			AddrFrom: &types.NetAddr{
				Services: protocol.SrvNodeNetwork,
				Ip:       &types.IPv4{Ip: []byte{127, 0, 0, 1}},
				Port:     9334,
			},
			Nonce:       nonce(),
			UserAgent:   &types.VarStr{Str: protocol.NewUserAgent().Str},
			StartHeight: -1,
			Relay:       true,
		}

		msg, err := protocol.NewMessage("version", network, version)
		if err != nil {
			logrus.Fatalln(err)
		}

		msgSerialized, err := proto.Marshal(msg)
		if err != nil {
			logrus.Fatalln(err)
		}

		conn, err := net.Dial("tcp", nodeURL)
		if err != nil {
			logrus.Fatalln(err)
		}
		defer conn.Close()

		_, err = conn.Write(msgSerialized)
		if err != nil {
			logrus.Fatalln(err)
		}

		tmp := make([]byte, 256)

		for {
			n, err := conn.Read(tmp)
			if err != nil {
				if err != io.EOF {
					logrus.Fatalln(err)
				}
				return
			}
			logrus.Infof("received: %x", tmp[:n])
		}
	},
}

func nonce() uint64 {
	return rand.Uint64()
}

// Execute ...
func Execute() {
	if err := btcNodeCmd.Execute(); err != nil {
		logrus.Fatalln(err)
		os.Exit(1)
	}
}
