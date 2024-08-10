package zk

func readByte(bytes []byte, idx int) (byte, int) {
	return bytes[idx], idx + 1
}

func putByte(bytes []byte, idx int, x byte) int {
	bytes[idx] = x
	return idx + 1
}
