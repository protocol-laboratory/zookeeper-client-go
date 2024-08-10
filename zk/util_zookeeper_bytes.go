package zk

// This file is for zookeeper code bytes type. Format method as alpha order.

func readData(bytes []byte, idx int) ([]byte, int) {
	return readBytes(bytes, idx)
}

func putData(bytes []byte, idx int, dstBytes []byte) int {
	return putBytes(bytes, idx, dstBytes)
}

func readPassword(bytes []byte, idx int) ([]byte, int) {
	return readBytes(bytes, idx)
}

func putPassword(bytes []byte, idx int, dstBytes []byte) int {
	return putBytes(bytes, idx, dstBytes)
}
