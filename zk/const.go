package zk

const (
	LenAclVersion       = 4
	LenArray            = 4
	LenChildVersion     = 4
	LenCreated          = LenTime
	LenCreatedZxId      = LenZxId
	LenDataLength       = 4
	LenEphemeralOwner   = 8
	LenError            = 4
	LenFlags            = 4
	LenLastModified     = LenTime
	LenLastModifiedZxId = LenZxId
	LenLastZxidSeen     = LenZxId
	LenLength           = 4
	LenNumberOfChildren = 4
	LenOpCode           = 4
	LenPeerZxId         = LenZxId
	LenPermission       = 4
	LenProtocolVersion  = 4
	LenReadonly         = 1
	LenSessionId        = 8
	LenTime             = 8
	LenTimeout          = 4
	LenTransactionId    = 4
	LenVersion          = 4
	LenWatch            = 1
	LenZxId             = 8
)

var (
	PasswordEmpty = make([]byte, 16)
)
