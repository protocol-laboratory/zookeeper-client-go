package zk

func readCredentials(bytes []byte, idx int) (string, int) {
	return readString(bytes, idx)
}

func putCredentials(bytes []byte, idx int, str string) int {
	return putString(bytes, idx, str)
}

func readPath(bytes []byte, idx int) (string, int) {
	return readString(bytes, idx)
}

func putPath(bytes []byte, idx int, str string) int {
	return putString(bytes, idx, str)
}

func readScheme(bytes []byte, idx int) (string, int) {
	return readString(bytes, idx)
}

func putScheme(bytes []byte, idx int, str string) int {
	return putString(bytes, idx, str)
}
