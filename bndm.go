// Package bdnm implements the Backward Nondeterministic Dawg Matching algorithm
// currently only works on strings
package bndm

type Pattern struct {
	t           [256]uint
	x           uint
	pattern_len int
}

func Compile(input string) *Pattern {
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

func (t *Pattern) Search(subject string) int {
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
