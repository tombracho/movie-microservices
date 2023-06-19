package main

import "testing"

func BenchmarkSerializedToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToJSON(metadata)
	}
}

func BenchmarkSerializedToXML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToXML(metadata)
	}
}

func BenchmarkSerializedToProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeToProto(genMetadata)
	}
}
