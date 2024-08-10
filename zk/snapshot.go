package zk

import (
	"runtime/debug"
	"strings"
)

type Snapshot struct {
	Magic      string
	Version    int
	DbId       int64
	SessionMap map[int64]int
	AclMap     map[int64][]*Acl
	Root       *DataNode
	NodeMap    map[string]*DataNode
	NodeCount  int
}

func DecodeSnapshot(bytes []byte) (s *Snapshot, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			s = nil
		}
	}()
	s = &Snapshot{}
	idx := 0
	s.Magic, idx = readMagic(bytes, idx)
	s.Version, idx = readInt(bytes, idx)
	s.DbId, idx = readInt64(bytes, idx)
	sessionCount, idx := readInt(bytes, idx)
	s.SessionMap = make(map[int64]int, sessionCount)
	for i := 0; i < sessionCount; i++ {
		var sessionId int64
		var timeout int
		sessionId, idx = readInt64(bytes, idx)
		timeout, idx = readInt(bytes, idx)
		s.SessionMap[sessionId] = timeout
	}
	aclMapSize, idx := readInt(bytes, idx)
	s.AclMap = make(map[int64][]*Acl, aclMapSize)
	for i := 0; i < aclMapSize; i++ {
		var key int64
		key, idx = readInt64(bytes, idx)
		var aclList []*Acl
		var aclListSize int
		aclListSize, idx = readInt(bytes, idx)
		for j := 0; j < aclListSize; j++ {
			var acl *Acl
			acl, idx = readAcl(bytes, idx)
			aclList = append(aclList, acl)
		}
		s.AclMap[key] = aclList
	}
	s.NodeMap = make(map[string]*DataNode)
	path, idx := readString(bytes, idx)
	for {
		if path == "/" {
			break
		}
		var dataNode *DataNode
		dataNode, idx = readDataNode(bytes, idx)
		dataNode.Path = path
		s.NodeCount++
		s.NodeMap[path] = dataNode
		if !strings.Contains(path, "/") {
			s.Root = dataNode
		} else {
			parentPath := path[:strings.LastIndex(path, "/")]
			parent := s.NodeMap[parentPath]
			parent.Children = append(parent.Children, dataNode)
		}
		path, idx = readString(bytes, idx)
	}
	return s, nil
}
