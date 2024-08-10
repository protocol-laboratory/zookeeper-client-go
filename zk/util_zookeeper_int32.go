package zk

// This file is for zookeeper code int32 type. Format method as alpha order.

func readOpCode(bytes []byte, idx int) (OpCode, int) {
	aux, i := readInt32(bytes, idx)
	return OpCode(aux), i
}

func putOpCode(bytes []byte, idx int, x OpCode) int {
	return putInt32(bytes, idx, int32(x))
}
