package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// Client is a JSON-RPC client.
// Connection to the server.
type Client struct {
	conn    net.Conn
	jsonrpc *rpc.Client
}

// NewClient returns a new Client.
func NewClient(port int) (*Client, error) {
	// Get the server address from the environment variable SERVER_ADDRESS.
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "127.0.0.1" // Default to localhost if SERVER_ADDRESS is not set.
	}

	// Establish TCP connection to the server.
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddress, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}

	// Create JSON-RPC client using the connection.
	rpcClient := jsonrpc.NewClient(conn)

	// Initialize client struct.
	client := &Client{
		conn:    conn,
		jsonrpc: rpcClient,
	}

	return client, nil
}

// Call calls the remote RPC method.
func (c *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	// Call the remote RPC method using the embedded JSON-RPC client.
	err := c.jsonrpc.Call(serviceMethod, args, reply)
	if err != nil {
		return fmt.Errorf("RPC call failed: %v", err)
	}
	return nil
}

// Close closes the connection to the server.
func (c *Client) Close() error {
	// Close the TCP connection.
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %v", err)
	}
	return nil
}
