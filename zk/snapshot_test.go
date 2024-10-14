package zk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDecodeSnapshot zookeeper without traffic
func TestDecodeSnapshot0(t *testing.T) {
	bytes, err := os.ReadFile("snapshot.0")
	assert.Nil(t, err)
	s, err := DecodeSnapshot(bytes)
	assert.Nil(t, err)
	assert.Len(t, s.AclMap, 1)
}
