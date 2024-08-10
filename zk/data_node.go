package zk

type DataNode struct {
	Path     string
	Data     []byte
	Acl      int64
	Stat     *StatPersisted
	Children []*DataNode
}

func readDataNode(bytes []byte, idx int) (dataNode *DataNode, nextIdx int) {
	dataNode = &DataNode{}
	dataNode.Data, idx = readBytes(bytes, idx)
	dataNode.Acl, idx = readInt64(bytes, idx)
	dataNode.Stat, idx = readStatPersisted(bytes, idx)
	return dataNode, idx
}
