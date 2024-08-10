package zk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeProtocolVersion(t *testing.T) {
	bytes := testHex2Bytes(t, "00000000")
	protocolVersion, idx := readProtocolVersion(bytes, 0)
	assert.Equal(t, 4, idx)
	assert.Equal(t, 0, protocolVersion)
}

func TestEncodeProtocolVersion(t *testing.T) {
	bytes := make([]byte, 4)
	putProtocolVersion(bytes, 0, 0)
	assert.Equal(t, testHex2Bytes(t, "00000000"), bytes)
}