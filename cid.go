package cid

//Correct Identifier

// 2059-5-26 1:38:27

import (
	crand "crypto/rand"
	"encoding/binary"
	"math"
	"math/rand"
	"strings"
	"time"
)

const keyString string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	var n uint64
	binary.Read(crand.Reader, binary.LittleEndian, &n)
	rand.Seed(time.Now().UnixMilli() + int64(n))
}

func NewBytesID(t time.Time) []byte {
	b1 := Uint64ToBytes(uint64(t.UnixMilli()))
	b2 := randomPadding()
	return merge(b1[:6], b2[:4])
}

func NewStringID(t time.Time) string {
	b1 := Uint64ToBytes(uint64(t.UnixMilli()))
	b2 := randomPadding()
	return b2s(b1[:6], b2[:4])
}

func BytesIDToTimestamp(bytesID []byte) int64 {
	return int64(bytesToUint64(bytesID[:6]))
}

func StringIDToTimestamp(stringID string) int64 {
	return int64(u36ToTen(stringID[:8]))
}

func StringIDToBytesID(stringID string) []byte {
	t1 := u36ToTen(stringID[:8])
	t2 := u36ToTen(stringID[8:])
	b1 := Uint64ToBytes(t1)
	b2 := Uint64ToBytes(t2)
	return merge(b1[:6], b2[2:6])
}

func BytesIDToStringID(bytesID []byte) string {
	return b2s(bytesID[:6], bytesID[6:10])
}

func b2s(b1, b2 []byte) string {
	i1 := uint64(b1[5]) | uint64(b1[4])<<8 | uint64(b1[3])<<16 |
		uint64(b1[2])<<24 | uint64(b1[1])<<32 | uint64(b1[0])<<40
	i2 := uint64(b2[3]) | uint64(b2[2])<<8 | uint64(b2[1])<<16 |
		uint64(b2[0])<<24
	return tenToU36(i1) + tenToU36(i2)
}

func randomPadding() []byte {
	b := make([]byte, 4)
	nn := rand.Int63()
	b[0] = byte(nn << 32 >> 40)
	b[1] = byte(nn << 24 >> 40)
	b[2] = byte(nn << 16 >> 40)
	b[3] = byte(nn << 8 >> 40)
	return b
}

func merge(b1 []byte, b2 []byte) []byte {
	if len(b1) != 6 || len(b2) != 4 {
		return []byte{}
	}
	id := make([]byte, 10)
	id[0] = b1[0]
	id[1] = b1[1]
	id[2] = b1[2]
	id[3] = b1[3]
	id[4] = b1[4]
	id[5] = b1[5]
	id[6] = b2[0]
	id[7] = b2[1]
	id[8] = b2[2]
	id[9] = b2[3]
	return id
}

func tenToU36(n uint64) string {
	r := ""
	for n != 0 {
		r = string(keyString[n%36]) + r
		n = n / 36
	}
	for i := len(r); i < 8; i++ {
		r = "0" + r
	}
	return r
}

func u36ToTen(s string) uint64 {
	var n uint64
	l := len(s)
	for i := 0; i < l; i++ {
		r := strings.IndexByte(keyString, s[i])
		n = n + uint64(float64(r)*math.Pow(float64(36), float64(l-i-1)))
	}
	return n
}

func bytesToUint64(b []byte) uint64 {
	if len(b) > 6 {
		return 0
	}
	return uint64(b[5]) | uint64(b[4])<<8 | uint64(b[3])<<16 |
		uint64(b[2])<<24 | uint64(b[1])<<32 | uint64(b[0])<<40
}

func Uint64ToBytes(u uint64) []byte {
	if u>>40 > 255 {
		return nil
	}
	b := make([]byte, 6)
	b[5] = byte(u)
	b[4] = byte(u >> 8)
	b[3] = byte(u >> 16)
	b[2] = byte(u >> 24)
	b[1] = byte(u >> 32)
	b[0] = byte(u >> 40)
	return b
}
