package zk

func readBool(bytes []byte, idx int) (bool, int) {
	return bytes[idx] == 1, idx + 1
}

func putBool(bytes []byte, idx int, x bool) int {
	if x {
		bytes[idx] = 1
	}
	return idx + 1
}
