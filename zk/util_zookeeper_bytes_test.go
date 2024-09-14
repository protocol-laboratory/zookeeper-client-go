package zk

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodePassword(t *testing.T) {
	bytes := testx.Hex2Bytes(t, "0000001000000000000000000000000000000000")
	password, idx := readPassword(bytes, 0)
	assert.Equal(t, 20, idx)
	assert.Len(t, password, 16)
}

func TestEncodePassword(t *testing.T) {
	bytes := make([]byte, 20)
	putPassword(bytes, 0, testx.Hex2Bytes(t, "00000000000000000000000000000000"))
	assert.Equal(t, testx.Hex2Bytes(t, "0000001000000000000000000000000000000000"), bytes)
}
