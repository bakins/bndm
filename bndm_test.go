package bndm

import (
	"testing"
)

func TestCompile(t *testing.T) {
	p := Compile("needle")
	if p == nil {
		t.Error("Compile failed")
	}
}

func TestSearch(t *testing.T) {
	p := Compile("needle")
	if p == nil {
		t.Error("Compile failed")
	}
	subject := "haystackneedlehaystack"
	index := p.Search(subject)

	if index != 8 {
		t.Error("Search failed: index: ", index)
	}
	if subject[index] != 'n' {
		t.Error("Search failed: character: ", subject[index])
	}
}
