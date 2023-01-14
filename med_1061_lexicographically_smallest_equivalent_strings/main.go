package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "parker"
	s2 := "morris"
	baseStr := "parser"

	fmt.Printf("%v\n", smallestEquivalentString(s1, s2, baseStr))
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := newUF(len(s1))

	for i := 0; i < len(s1); i++ {
		uf.Union(s1[i], s2[i])
	}

	var builder strings.Builder

	for i := 0; i < len(baseStr); i++ {
		s, ok := uf.Find(baseStr[i])
		if ok {
			builder.WriteByte(s.val)
		} else {
			builder.WriteByte(baseStr[i])
		}
	}

	return builder.String()
}

type UFItem struct {
	parent *UFItem
	val    byte
}

type UF struct {
	roots map[byte]*UFItem
	m     map[byte]byte // [val]tree_id
}

func newUF(n int) *UF {
	return &UF{
		roots: make(map[byte]*UFItem, n),
		m:     make(map[byte]byte, n),
	}
}

func (uf *UF) Find(b byte) (*UFItem, bool) {
	r, ok := uf.m[b]
	if !ok {
		return nil, false
	}

	if uf.roots[r].parent != nil {
		return uf.Find(uf.roots[r].parent.val)
	}

	return uf.roots[r], true
}

func (uf *UF) Union(t1, t2 byte) {
	if t1 == t2 {
		return
	}

	r1, ok1 := uf.Find(t1)
	if !ok1 {
		r1 = &UFItem{val: t1}
	}

	r2, ok2 := uf.Find(t2)
	if !ok2 {
		r2 = &UFItem{val: t2}
	}

	if r1.val < r2.val {
		r2.parent = r1
		r1.parent = nil

		uf.roots[r1.val] = r1
		uf.m[t1] = r1.val
		uf.m[t2] = r1.val
	} else {
		r1.parent = r2
		r2.parent = nil

		uf.roots[r2.val] = r2
		uf.m[t1] = r2.val
		uf.m[t2] = r2.val
	}
}
