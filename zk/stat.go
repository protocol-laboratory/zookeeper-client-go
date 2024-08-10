package zk

import "runtime/debug"

type Stat struct {
	CreatedZxId              int64
	LastModifiedZxId         int64
	Created                  int64
	LastModified             int64
	Version                  int
	ChildVersion             int
	AclVersion               int
	EphemeralOwner           int64
	DataLength               int
	NumChildren              int
	LastModifiedChildrenZxId int64
}

func DecodeStat(bytes []byte) (stat *Stat, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			stat = nil
		}
	}()
	stat = &Stat{}
	idx := 0
	stat.CreatedZxId, idx = readCreatedZxId(bytes, idx)
	stat.LastModifiedZxId, idx = readLastModifiedZxId(bytes, idx)
	stat.Created, idx = readCreated(bytes, idx)
	stat.LastModified, idx = readLastModified(bytes, idx)
	stat.Version, idx = readVersion(bytes, idx)
	stat.ChildVersion, idx = readChildVersion(bytes, idx)
	stat.AclVersion, idx = readAclVersion(bytes, idx)
	stat.EphemeralOwner, idx = readEphemeralOwner(bytes, idx)
	stat.DataLength, idx = readDataLength(bytes, idx)
	stat.NumChildren, idx = readNumChildren(bytes, idx)
	stat.LastModifiedChildrenZxId, idx = readLastModifiedChildrenZxId(bytes, idx)
	return stat, nil
}

func (s *Stat) BytesLength() int {
	length := 0
	length += LenCreatedZxId + LenLastModifiedZxId
	length += LenCreated + LenLastModified + LenVersion + LenChildVersion + LenAclVersion + LenEphemeralOwner
	length += LenDataLength + LenNumberOfChildren + LenLastModifiedZxId
	return length
}

func (s *Stat) Bytes() []byte {
	bytes := make([]byte, s.BytesLength())
	idx := 0
	idx = putCreatedZxId(bytes, idx, s.CreatedZxId)
	idx = putLastModifiedZxId(bytes, idx, s.LastModifiedZxId)
	idx = putCreated(bytes, idx, s.Created)
	idx = putLastModified(bytes, idx, s.LastModified)
	idx = putVersion(bytes, idx, s.Version)
	idx = putChildVersion(bytes, idx, s.ChildVersion)
	idx = putAclVersion(bytes, idx, s.AclVersion)
	idx = putEphemeralOwner(bytes, idx, s.EphemeralOwner)
	idx = putDataLength(bytes, idx, s.DataLength)
	idx = putNumChildren(bytes, idx, s.NumChildren)
	idx = putLastModifiedChildrenZxId(bytes, idx, s.LastModifiedChildrenZxId)
	return bytes
}

func readStat(bytes []byte, idx int) (*Stat, int) {
	stat, err := DecodeStat(bytes[idx:])
	if err != nil {
		panic(err)
	}
	return stat, idx + stat.BytesLength()
}

func putStat(bytes []byte, idx int, stat *Stat) int {
	idx = putBytesWithoutLen(bytes, idx, stat.Bytes())
	return idx
}
