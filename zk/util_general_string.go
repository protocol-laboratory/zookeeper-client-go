package zk

func readString(bytes []byte, idx int) (string, int) {
	length, idx := readInt(bytes, idx)
	return string(bytes[idx : idx+length]), idx + length
}

func putString(bytes []byte, idx int, str string) int {
	strBytes := []byte(str)
	idx = putInt(bytes, idx, len(strBytes))
	copy(bytes[idx:idx+len(strBytes)], strBytes)
	return idx + len(strBytes)
}

func StrLen(str string) int {
	return 4 + len([]byte(str))
}
