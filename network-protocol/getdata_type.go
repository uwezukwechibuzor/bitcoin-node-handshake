package networkprotocol

// MsgGetData represents 'getdata' message.
type MsgGetData struct {
	Count     uint8
	Inventory []InvVector
}
