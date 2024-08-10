package zk

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testHex2Bytes(t *testing.T, str string) []byte {
	bytes, err := hex.DecodeString(str)
	assert.Nil(t, err)
	return bytes
}
