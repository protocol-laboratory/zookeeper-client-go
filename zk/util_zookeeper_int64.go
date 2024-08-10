package zk

// This file is for zookeeper code int64 type. Format method as alpha order.

func readCreated(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putCreated(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readCreatedZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putCreatedZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readEphemeralOwner(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putEphemeralOwner(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastModified(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastModified(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastModifiedChildrenZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastModifiedChildrenZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastModifiedZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastModifiedZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastZxidSeen(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastZxidSeen(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readPeerZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putPeerZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readSessionId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putSessionId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}
