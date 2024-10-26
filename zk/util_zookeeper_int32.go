package zk

// This file is for zookeeper code int32 type. Format method as alpha order.

func readOpCode(bytes []byte, idx int) (OpCode, int) {
	aux, i := readInt32(bytes, idx)
	return OpCode(aux), i
}

func putOpCode(bytes []byte, idx int, x OpCode) int {
	return putInt32(bytes, idx, int32(x))
}

func readTransactionId(bytes []byte, idx int) (int32, int) {
	return readInt32(bytes, idx)
}

func putTransactionId(bytes []byte, idx int, x int32) int {
	return putInt32(bytes, idx, x)
}
