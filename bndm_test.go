package bndm

import (
	"bytes"
	"crypto/rand"
	"fmt"
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

func BenchmarkSearchWithShortSubject(b *testing.B) {
	p := Compile([]byte("needle"))
	subject := "haystackneedlehaystack"
	for i := 0; i < b.N; i++ {
		p.Search([]byte(subject))
	}
}

func BenchmarkIndexWithShortSubject(b *testing.B) {
	subject := "haystackneedlehaystack"
	for i := 0; i < b.N; i++ {
		strings.Index(subject, "needle")
	}
}

// http://stackoverflow.com/a/12772666
func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func BenchmarkSearchWithRandomSubject(b *testing.B) {
	p := Compile([]byte("needle"))
	subject := randString(16384)
	for i := 0; i < b.N; i++ {
		p.Search([]byte(subject))
	}
}

func BenchmarkIndexWithRandomSubject(b *testing.B) {
	subject := randString(16384)
	for i := 0; i < b.N; i++ {
		strings.Index(subject, "needle")
	}
}

func longSubject(n int, haystack string, needle string) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(haystack)
	}
	return fmt.Sprint(buffer.String())
}

func BenchmarkSearchWithLongSubject(b *testing.B) {
	p := Compile([]byte("needle"))
	subject := longSubject(4092, "haystack", "needle")
	for i := 0; i < b.N; i++ {
		p.Search([]byte(subject))
	}
}

func BenchmarkIndexWithLongSubject(b *testing.B) {
	subject := longSubject(4092, "haystack", "needle")
	for i := 0; i < b.N; i++ {
		strings.Index(subject, "needle")
	}
}
