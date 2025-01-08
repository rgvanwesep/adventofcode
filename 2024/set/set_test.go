package set

import (
	"slices"
	"testing"
)

func TestSet(t *testing.T) {
	var (
		i  int
		b  byte
		bs []byte
		ok bool
	)
	s := NewSet[byte]()

	i = s.Len()
	if i != 0 {
		t.Errorf("Len() == %d, expected 0", i)
	}
	ok = s.Contains('a')
	if ok {
		t.Errorf("Contains('a') == %t, expected false", ok)
	}
	bs = []byte{}
	for b = range s.All() {
		bs = append(bs, b)
	}
	if len(bs) != 0 {
		t.Errorf("All() == %q, expected \"\"", string(bs))
	}

	s.Add('a')
	i = s.Len()
	if i != 1 {
		t.Errorf("Len() == %d, expected 1", i)
	}
	ok = s.Contains('a')
	if !ok {
		t.Errorf("Contains('a') == %t, expected true", ok)
	}
	ok = s.Contains('b')
	if ok {
		t.Errorf("Contains('b') == %t, expected false", ok)
	}
	bs = []byte{}
	for b = range s.All() {
		bs = append(bs, b)
	}
	if len(bs) != 1 || bs[0] != 'a' {
		t.Errorf("All() == %q, expected \"a\"", string(bs))
	}

	s.Remove('b')
	i = s.Len()
	if i != 1 {
		t.Errorf("Len() == %d, expected 1", i)
	}
	ok = s.Contains('a')
	if !ok {
		t.Errorf("Contains('a') == %t, expected true", ok)
	}
	ok = s.Contains('b')
	if ok {
		t.Errorf("Contains('b') == %t, expected false", ok)
	}
	bs = []byte{}
	for b = range s.All() {
		bs = append(bs, b)
	}
	if len(bs) != 1 || bs[0] != 'a' {
		t.Errorf("All() == %q, expected \"a\"", string(bs))
	}

	s.Remove('a')
	i = s.Len()
	if i != 0 {
		t.Errorf("Len() == %d, expected 0", i)
	}
	ok = s.Contains('a')
	if ok {
		t.Errorf("Contains('a') == %t, expected false", ok)
	}
	bs = []byte{}
	for b = range s.All() {
		bs = append(bs, b)
	}
	if len(bs) != 0 {
		t.Errorf("All() == %q, expected \"\"", string(bs))
	}

	s.Add('a')
	s.Add('b')
	sClone := s.Clone()

	s.Clear()
	i = s.Len()
	if i != 0 {
		t.Errorf("Len() == %d, expected 0", i)
	}
	ok = s.Contains('a')
	if ok {
		t.Errorf("Contains('a') == %t, expected false", ok)
	}
	bs = []byte{}
	for b = range s.All() {
		bs = append(bs, b)
	}
	if len(bs) != 0 {
		t.Errorf("All() == %q, expected \"\"", string(bs))
	}

	i = sClone.Len()
	if i != 2 {
		t.Errorf("Len() == %d, expected 2", i)
	}
	ok = sClone.Contains('a')
	if !ok {
		t.Errorf("Contains('a') == %t, expected true", ok)
	}
	ok = sClone.Contains('b')
	if !ok {
		t.Errorf("Contains('b') == %t, expected true", ok)
	}
	bs = []byte{}
	for b = range sClone.All() {
		bs = append(bs, b)
	}
	slices.Sort(bs)
	if len(bs) != 2 || bs[0] != 'a' || bs[1] != 'b' {
		t.Errorf("All() == %q, expected \"ab\"", string(bs))
	}
	for range sClone.All() {
		break
	}
}
