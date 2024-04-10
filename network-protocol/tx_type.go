package networkprotocol

// MsgTx represents 'tx' message.
type MsgTx struct {
	Version    int32
	Flag       uint16
	TxInCount  uint8
	TxIn       []TxInput
	TxOutCount uint8
	TxOut      []TxOutput
	TxWitness  TxWitnessData
	LockTime   uint32
}

// TxInput represents transaction input.
type TxInput struct {
	PreviousOutput  OutPoint
	ScriptLength    uint8
	SignatureScript []byte
	Sequence        uint32
}

// TxOutput represents transaction output.
type TxOutput struct {
	Value          int64
	PkScriptLength uint8
	PkScript       []byte
}

// TxWitnessData represents transaction witness data.
type TxWitnessData struct {
	Count   uint8
	Witness []TxWitness
}

// TxWitness represents a component of transaction witness data.
type TxWitness struct {
	Length uint8
	Data   []byte
}

// OutPoint represents previous output transaction reference.
type OutPoint struct {
	Hash  [32]byte
	Index uint32
}
