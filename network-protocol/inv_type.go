package networkprotocol

const (
	DataObjectError = iota
	DataObjectTx
	DataObjectBlock
	DataObjectFilterBlock
	DataObjectCmpctBlock
)

// MsgInv represents 'inv' message.
type MsgInv struct {
	Count     uint8
	Inventory []InvVector
}

// InvVector represents inventory vector.
type InvVector struct {
	Type uint32
	Hash [32]byte
}
