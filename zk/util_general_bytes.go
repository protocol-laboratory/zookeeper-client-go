package zk

func readBytes(bytes []byte, idx int) ([]byte, int) {
	length, idx := readInt(bytes, idx)
	return bytes[idx : idx+length], idx + length
}

func putBytes(bytes []byte, idx int, srcBytes []byte) int {
	idx = putInt(bytes, idx, len(srcBytes))
	copy(bytes[idx:], srcBytes)
	return idx + len(srcBytes)
}

func readBytesNum(bytes []byte, idx int, num int) ([]byte, int) {
	return bytes[idx : idx+num], idx + num
}

func putBytesWithoutLen(bytes []byte, idx int, srcBytes []byte) int {
	copy(bytes[idx:], srcBytes)
	return idx + len(srcBytes)
}

func BytesLen(bytes []byte) int {
	return 4 + len(bytes)
}
