package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeProtocolVersion(t *testing.T) {
	bytes := hex2Bytes(t, "00000000")
	protocolVersion, idx := readProtocolVersion(bytes, 0)
	assert.Equal(t, 4, idx)
	assert.Equal(t, 0, protocolVersion)
}

func TestEncodeProtocolVersion(t *testing.T) {
	bytes := make([]byte, 4)
	putProtocolVersion(bytes, 0, 0)
	assert.Equal(t, hex2Bytes(t, "00000000"), bytes)
}
