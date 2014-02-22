package bndm

import (
	"strings"
	"testing"
)

func TestCompile(t *testing.T) {
	p := Compile([]byte("needle"))
	if p == nil {
		t.Error("Compile failed")
	}
}

func TestSearch(t *testing.T) {
	p := Compile([]byte("needle"))
	if p == nil {
		t.Error("Compile failed")
	}
	subject := "haystackneedlehaystack"
	index := p.Search([]byte(subject))

	if index != 8 {
		t.Error("Search failed: index: ", index)
	}
	if subject[index] != 'n' {
		t.Error("Search failed: character: ", subject[index])
	}
}

func BenchmarkCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := Compile([]byte("needle"))
		if p == nil {
			b.Error("Compile failed")
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	p := Compile([]byte("needle"))
	if p == nil {
		b.Error("Compile failed")
	}
	subject := "haystackneedlehaystack"

	for i := 0; i < b.N; i++ {
		index := p.Search([]byte(subject))

		if index != 8 {
			b.Error("Search failed: index: ", index)
		}
	}
}

func BenchmarkIndex(b *testing.B) {
	subject := "haystackneedlehaystack"

	for i := 0; i < b.N; i++ {
		index := strings.Index(subject, "needle")
		if index != 8 {
			b.Error("Search failed: index: ", index)
		}
	}
}
