package zk

import "runtime/debug"

type Id struct {
	Scheme string
	Id     string
}

func DecodeId(bytes []byte) (id *Id, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			id = nil
		}
	}()
	id = &Id{}
	idx := 0
	id.Scheme, idx = readString(bytes, idx)
	id.Id, idx = readString(bytes, idx)
	return id, nil
}

func (i *Id) ByteLength() int {
	return StrLen(i.Scheme) + StrLen(i.Id)
}

func readId(bytes []byte, idx int) (*Id, int) {
	id, err := DecodeId(bytes[idx:])
	if err != nil {
		panic(err)
	}
	return id, idx + id.ByteLength()
}
