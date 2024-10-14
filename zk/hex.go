package zk

import (
	"encoding/hex"
	"github.com/libgox/buffer"
	"github.com/stretchr/testify/require"
	"testing"
)

func hex2Bytes(t *testing.T, str string) []byte {
	bytes, err := hex.DecodeString(str)
	require.NoError(t, err)
	return bytes
}

func hex2Buffer(t *testing.T, str string) *buffer.Buffer {
	bytes, err := hex.DecodeString(str)
	require.NoError(t, err)
	return buffer.NewBufferFromBytes(bytes)
}
