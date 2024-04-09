package protocol

import (
	"crypto/sha256"
	"fmt"

	types "github.com/uwezukwechibuzor/bitcoin-node-handshake/types/proto"
	"google.golang.org/protobuf/proto"
)

const (
	checksumLength = 4
	nodeNetwork    = 1
	magicLength    = 4
)

var (
	magicMainnet = [magicLength]byte{0xf9, 0xbe, 0xb4, 0xd9}
	magicSimnet  = [magicLength]byte{0x16, 0x1c, 0x14, 0x12}
	networks     = map[string][magicLength]byte{
		"mainnet": magicMainnet,
		"simnet":  magicSimnet,
	}
)

type MessagePayload interface {
	Serialize() ([]byte, error)
}

type Message types.Message

// Serialize serializes the MyPayload into a byte slice.
func (m *Message) Serialize() ([]byte, error) {
	// Serialize the payload using Protocol Buffers
	serialized, err := proto.Marshal((*types.Message)(m))
	if err != nil {
		return nil, err
	}
	return serialized, nil
}

type MsgVersion types.MsgVersion

// Serialize serializes the message version into a byte slice.
func (m *MsgVersion) Serialize() ([]byte, error) {
	// Serialize the payload using Protocol Buffers
	serialized, err := proto.Marshal((*types.MsgVersion)(m))
	if err != nil {
		return nil, err
	}
	return serialized, nil
}

// NewMessage creates a new Message using Protocol Buffers.
func NewMessage(cmd string, network string, payload MessagePayload) (*types.Message, error) {

	serializedPayload, _ := payload.Serialize()

	command, ok := commands[cmd]
	if !ok {
		return nil, fmt.Errorf("unsupported command %s", cmd)
	}

	magic, ok := networks[network]
	if !ok {
		return nil, fmt.Errorf("unsupported network %s", network)
	}

	message := &types.Message{
		Magic:    magic[:],
		Command:  command[:],
		Length:   uint32(len(serializedPayload)),
		Checksum: checksum(serializedPayload),
		Payload:  serializedPayload,
	}

	return message, nil
}

// checksum calculates the checksum of the data.
func checksum(data []byte) []byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return hash[:checksumLength]
}
