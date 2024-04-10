package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/node"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/rpc"
)

const (
	userAgent = "/Satoshi:5.64/tinybit:0.0.1/"
)

var (
	network     string
	nodeAddr    string
	jsonrpcPort int
)

func init() {
	btchandshakeCmd.Flags().IntVar(&jsonrpcPort, "jsonrpc-port", 9334, "Port to listen JSON-RPC connections on")
	btchandshakeCmd.Flags().StringVar(&nodeAddr, "node-addr", "127.0.0.1:9333", "TCP address of a Bitcoin node to connect to")
	btchandshakeCmd.Flags().StringVar(&network, "network", "simnet", "Bitcoin network (simnet, mainnet)")
}

var btchandshakeCmd = &cobra.Command{
	Use: "btchandshake",
	RunE: func(cmd *cobra.Command, args []string) error {
		node, err := node.New(network, userAgent)
		if err != nil {
			return err
		}

		rpc, err := rpc.NewServer(jsonrpcPort, node)
		if err != nil {
			return err
		}

		logrus.Infof("Running JSON-RPC server on port %d", jsonrpcPort)
		go rpc.Run()

		return node.Run(nodeAddr)
	},
}

// Execute
func Execute() {
	if err := btchandshakeCmd.Execute(); err != nil {
		logrus.Fatalln(err)
		os.Exit(1)
	}
}
