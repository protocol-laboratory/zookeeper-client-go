package zk

func readMagic(bytes []byte, idx int) (string, int) {
	bytesNum, i := readBytesNum(bytes, idx, 4)
	return string(bytesNum), i
}
