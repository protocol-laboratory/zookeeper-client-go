package zk

import "runtime/debug"

type Acl struct {
	Perms int
	Id    *Id
}

func DecodeAcl(bytes []byte) (acl *Acl, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			acl = nil
		}
	}()
	acl = &Acl{}
	idx := 0
	acl.Perms, idx = readInt(bytes, idx)
	acl.Id, idx = readId(bytes, idx)
	return acl, nil
}

func (a *Acl) ByteLength() int {
	return 4 + a.Id.ByteLength()
}

func readAcl(bytes []byte, idx int) (*Acl, int) {
	acl, err := DecodeAcl(bytes[idx:])
	if err != nil {
		panic(err)
	}
	return acl, idx + acl.ByteLength()
}
