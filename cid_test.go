package cid

import (
	"testing"
	"time"
)

func BenchmarkBytesID(b *testing.B) {
	t := time.Now()
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = NewBytesID(t)
	}
}

func BenchmarkStringID(b *testing.B) {
	t := time.Now()
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = NewStringID(t)
	}
}
func BenchmarkStringIDToBytesID(b *testing.B) {
	id := NewStringID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = StringIDToBytesID(id)
	}
}
func BenchmarkBytesIDToStringID(b *testing.B) {
	id := NewBytesID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = BytesIDToStringID(id)
	}
}
func BenchmarkGetTimestampByStringID(b *testing.B) {
	id := NewStringID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = GetTimestampByStringID(id)
	}
}
func BenchmarkGetTimestampByBytesID(b *testing.B) {
	id := NewBytesID(time.Now())
	b.ReportAllocs()
	b.SetBytes(32)
	for i := 0; i < b.N; i++ {
		_ = GetTimestampByBytesID(id)
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
