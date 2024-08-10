package zk

type ErrorCode int32

const (
	EC_NoNodeError     ErrorCode = -101
	EC_NodeExistsError ErrorCode = -110
	EC_OK              ErrorCode = 0
)
