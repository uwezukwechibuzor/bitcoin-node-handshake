package protocol

import types "github.com/uwezukwechibuzor/bitcoin-node-handshake/types/proto"

func newVarStr(str string) *types.VarStr {
	return &types.VarStr{
		Length: uint32(len(str)),
		Str:    str,
	}
}
