package networkprotocol

// MsgPing describes 'ping' message.
type MsgPing struct {
	Nonce uint64
}

// MsgPong describes 'pong' message.
type MsgPong struct {
	Nonce uint64
}
