package cid_test

import (
	"testing"
	"time"

	"github.com/qorm/cid"
)

func BenchmarkBytesID(b *testing.B) {
	t := time.Now()
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = cid.NewBytesID(t)
	}
}

func BenchmarkStringID(b *testing.B) {
	t := time.Now()
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = cid.NewStringID(t)
	}
}
func BenchmarkStringIDToBytesID(b *testing.B) {
	id := cid.NewStringID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = cid.StringIDToBytesID(id)
	}
}
func BenchmarkBytesIDToStringID(b *testing.B) {
	id := cid.NewBytesID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = cid.BytesIDToStringID(id)
	}
}
func BenchmarkStringIDToTimestamp(b *testing.B) {
	id := cid.NewStringID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = cid.StringIDToTimestamp(id)
	}
}
func BenchmarkBytesIDToTimestamp(b *testing.B) {
	id := cid.NewBytesID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = cid.BytesIDToTimestamp(id)
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
