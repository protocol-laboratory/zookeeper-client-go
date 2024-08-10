package zk

// This file is for zookeeper code bool type. Format method as alpha order.

func putReadOnly(bytes []byte, idx int, x bool) int {
	return putBool(bytes, idx, x)
}

func readReadOnly(bytes []byte, idx int) (bool, int) {
	return readBool(bytes, idx)
}

func putWatch(bytes []byte, idx int, x bool) int {
	return putBool(bytes, idx, x)
}

func readWatch(bytes []byte, idx int) (bool, int) {
	return readBool(bytes, idx)
}
