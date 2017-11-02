package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	//fmt.Printf("-----\nword: %v\nbit: %v\n-----\n", word, bit) //bitがAddした値
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		fmt.Printf("i: %v, tword: %v\n", i, tword)
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	//要素数を返す
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	//セットからxを取り除く

}

func (s *IntSet) Clear() {
	//セットから全ての要素を取り除く
	a := new(IntSet)
	s.UnionWith(a)
}

//
func (s *IntSet) Copy() *IntSet {
	//セットのコピーを返す
	return s
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	a := new(IntSet)
	b := new(IntSet)
	a.Add(1)
	a.Add(3) //要素2を追加
	fmt.Printf("a.Has(5) : %v\n", a.Has(5))
	a.Add(5)
	a.Add(90)
	fmt.Printf("a.Has(5) : %v\n", a.Has(5))
	fmt.Printf("a : %v\n", a)
	fmt.Printf("b : %v\n", b)
	b.UnionWith(a)
	fmt.Printf("b : %v\n", b)
	fmt.Println(a.String())
	fmt.Println(b.String())

	fmt.Println()

	fmt.Println(a.Len())
	a.Clear()
	fmt.Println(a)
}
