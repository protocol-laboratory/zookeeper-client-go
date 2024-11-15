package zk

import "runtime/debug"

type CreateReq struct {
	TransactionId int32
	OpCode        OpCode
	Path          string
	Data          []byte
	Permissions   []int
	Scheme        string
	Credentials   string
	Flags         int
}

func DecodeCreateReq(bytes []byte) (req *CreateReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &CreateReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Data, idx = readData(bytes, idx)
	var length int
	length, idx = readInt(bytes, idx)
	for i := 0; i < length; i++ {
		var permission int
		permission, idx = readPermission(bytes, idx)
		req.Permissions = append(req.Permissions, permission)
	}
	req.Scheme, idx = readScheme(bytes, idx)
	req.Credentials, idx = readCredentials(bytes, idx)
	req.Flags, idx = readFlags(bytes, idx)
	return req, nil
}

func (c *CreateReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(c.Path) + BytesLen(c.Data)
	length += LenArray + LenPermission*len(c.Permissions)
	length += StrLen(c.Scheme) + StrLen(c.Credentials) + LenFlags
	return length
}

func (c *CreateReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, c.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putOpCode(bytes, idx, OpCreate)
	idx = putPath(bytes, idx, c.Path)
	idx = putData(bytes, idx, c.Data)
	idx = putInt(bytes, idx, len(c.Permissions))
	for i := 0; i < len(c.Permissions); i++ {
		idx = putPermission(bytes, idx, c.Permissions[i])
	}
	idx = putScheme(bytes, idx, c.Scheme)
	idx = putCredentials(bytes, idx, c.Credentials)
	idx = putFlags(bytes, idx, c.Flags)
	return bytes
}
