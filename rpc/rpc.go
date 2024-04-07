package rpc

// Node defines the interface of interaction between JSON-RPC server and a node.
type Node interface {
  // add field	
}

// RPC implements RPC interface of the node.
type RPC struct {
	node Node
}

// MempoolArgs are arguments of Mempool method.
type MempoolArgs interface{}

// MempoolReply is reply of Mempool method.
type MempoolReply string
