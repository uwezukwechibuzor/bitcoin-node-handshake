package networkprotocol

import (
	"io"

	"github.com/uwezukwechibuzor/bitcoin-node-handshake/pkg/binary"
)

// UnmarshalBinary implements binary.Unmarshaler interface.
func (inv *MsgInv) UnmarshalBinary(r io.Reader) error {
	d := binary.NewDecoder(r)

	if err := d.Decode(&inv.Count); err != nil {
		return err
	}

	for i := uint8(0); i < inv.Count; i++ {
		var v InvVector

		if err := d.Decode(&v); err != nil {
			return err
		}

		inv.Inventory = append(inv.Inventory, v)
	}

	return nil
}
