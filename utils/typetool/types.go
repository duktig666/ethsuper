// description:
// @author renshiwei
// Date: 2022/11/15 18:32

package util

import (
	"bytes"
	"encoding/binary"
)

// IntToBytes 整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// Int32ToBytes 整形转换成字节
func Int32ToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

// Int64ToBytes 整形转换成字节
func Int64ToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

// PadByteTo 补0
func PadByteTo(n int) []byte {
	return make([]byte, n)
}

// PadByteTo32 补0到32byte
func PadByteTo32(n int) []byte {
	return make([]byte, 32-n)
}
