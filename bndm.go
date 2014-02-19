// Package bdnm implements the Backward Nondeterministic Dawg Matching algorithm.
// Currently only tested on strings.
// This implemntation is based on the one in Apache HTTPD mod_include
package bndm

// Pattern is an opaque struct that represents a compiled bndm pattern
type Pattern struct {
	t           [256]uint
	x           uint
	pattern_len int
}

// Compile sets up a bndm pattern
func Compile(input []byte) *Pattern {
	t := &Pattern{}
	length := len(input)
	t.pattern_len = length
	n := 0
	var x uint
	for x = 1; n < length; x <<= 1 {
		c := uint(input[n])
		n++
		t.t[c] |= x
	}
	t.x = x - 1
	return t
}

// Search in a string for the pattern given to Compile. returns the index of the first instance of patter in subject,
// or -1 if it is not found
func (t *Pattern) Search(subject []byte) int {
	hl := len(subject)
	nl := t.pattern_len
	x := t.x

	pi := -1
	p := pi + nl
	he := hl
	var skip int
	var d uint
	for p < he {
		skip = p
		d = x
		for d != 0 {
			d &= t.t[subject[p]]
			p--
			if d == 0 {
				break
			}
			if d&1 != 0 {
				if p != pi {
					skip = p
				} else {
					return p + 1
				}
			}
			d >>= 1
		}
		pi = skip
		p = pi + nl
	}

	return -1
}
