package set

import (
	"fmt"
	"slices"
	"testing"
)

func TestSet(t *testing.T) {
	var (
		i   int
		b   byte
		bs  []byte
		str string
		ok  bool
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
	str = fmt.Sprint(s)
	if str != "{}" {
		t.Errorf("String() == %q, expected \"{}\"", str)
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
	str = fmt.Sprint(s)
	if str != fmt.Sprintf("{%v}", 'a') {
		t.Errorf("String() == %q, expected %q", str, fmt.Sprintf("{%v}", 'a'))
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
	str = fmt.Sprint(s)
	if str != fmt.Sprintf("{%v, %v}", 'a', 'b') && str != fmt.Sprintf("{%v, %v}", 'b', 'a') {
		t.Errorf("String() == %q, expected %q or %q",
			str, fmt.Sprintf("{%v, %v}", 'a', 'b'), fmt.Sprintf("{%v, %v}", 'b', 'a'),
		)
	}

	sClone := s.Clone()

	s.Clear()
	if Equals(s, sClone) {
		t.Errorf("Equals(%v, %v) == true, expected false", s, sClone)
	}
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

	s.Add('b')
	s.Add('c')
	if Equals(s, sClone) {
		t.Errorf("Equals(%v, %v) == true, expected false", s, sClone)
	}

	for range sClone.All() {
		break
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		x, y, expected string
	}{
		{
			x: "", y: "", expected: "",
		},
		{
			x: "a", y: "b", expected: "",
		},
		{
			x: "a", y: "a", expected: "a",
		},
		{
			x: "ab", y: "a", expected: "a",
		},
		{
			x: "b", y: "ab", expected: "b",
		},
		{
			x: "ab", y: "bc", expected: "b",
		},
		{
			x: "abc", y: "ab", expected: "ab",
		},
		{
			x: "bc", y: "abc", expected: "bc",
		},
	}
	for _, c := range cases {
		sx := NewSet[byte]()
		for i := range c.x {
			sx.Add(c.x[i])
		}
		sy := NewSet[byte]()
		for i := range c.y {
			sy.Add(c.y[i])
		}
		se := NewSet[byte]()
		for i := range c.expected {
			se.Add(c.expected[i])
		}
		z := Intersection(sx, sy)
		if !Equals(z, se) {
			t.Errorf("Intersection(%v, %v) == %v, expected %v", sx, sy, z, se)
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		x, y, expected string
	}{
		{
			x: "", y: "", expected: "",
		},
		{
			x: "a", y: "b", expected: "ab",
		},
		{
			x: "a", y: "a", expected: "a",
		},
		{
			x: "ab", y: "a", expected: "ab",
		},
		{
			x: "b", y: "ab", expected: "ab",
		},
		{
			x: "ab", y: "bc", expected: "abc",
		},
		{
			x: "abc", y: "ab", expected: "abc",
		},
		{
			x: "bc", y: "abc", expected: "abc",
		},
	}
	for _, c := range cases {
		sx := NewSet[byte]()
		for i := range c.x {
			sx.Add(c.x[i])
		}
		sy := NewSet[byte]()
		for i := range c.y {
			sy.Add(c.y[i])
		}
		se := NewSet[byte]()
		for i := range c.expected {
			se.Add(c.expected[i])
		}
		z := Union(sx, sy)
		if !Equals(z, se) {
			t.Errorf("Union(%v, %v) == %v, expected %v", sx, sy, z, se)
		}
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		x, y, expected string
	}{
		{
			x: "", y: "", expected: "",
		},
		{
			x: "a", y: "b", expected: "a",
		},
		{
			x: "a", y: "a", expected: "",
		},
		{
			x: "ab", y: "a", expected: "b",
		},
		{
			x: "b", y: "ab", expected: "",
		},
		{
			x: "ab", y: "bc", expected: "a",
		},
		{
			x: "abc", y: "ab", expected: "c",
		},
		{
			x: "bc", y: "abc", expected: "",
		},
	}
	for _, c := range cases {
		sx := NewSet[byte]()
		for i := range c.x {
			sx.Add(c.x[i])
		}
		sy := NewSet[byte]()
		for i := range c.y {
			sy.Add(c.y[i])
		}
		se := NewSet[byte]()
		for i := range c.expected {
			se.Add(c.expected[i])
		}
		z := Difference(sx, sy)
		if !Equals(z, se) {
			t.Errorf("Difference(%v, %v) == %v, expected %v", sx, sy, z, se)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	cases := []struct {
		x, y, expected string
	}{
		{
			x: "", y: "", expected: "",
		},
		{
			x: "a", y: "b", expected: "ab",
		},
		{
			x: "a", y: "a", expected: "",
		},
		{
			x: "ab", y: "a", expected: "b",
		},
		{
			x: "b", y: "ab", expected: "a",
		},
		{
			x: "ab", y: "bc", expected: "ac",
		},
		{
			x: "abc", y: "ab", expected: "c",
		},
		{
			x: "bc", y: "abc", expected: "a",
		},
	}
	for _, c := range cases {
		sx := NewSet[byte]()
		for i := range c.x {
			sx.Add(c.x[i])
		}
		sy := NewSet[byte]()
		for i := range c.y {
			sy.Add(c.y[i])
		}
		se := NewSet[byte]()
		for i := range c.expected {
			se.Add(c.expected[i])
		}
		z := SymmetricDifference(sx, sy)
		if !Equals(z, se) {
			t.Errorf("SymmetricDifference(%v, %v) == %v, expected %v", sx, sy, z, se)
		}
	}
}

func TestCartesianProduct(t *testing.T) {
	x := NewSet[byte]()
	y := NewSet[byte]()
	z := CartesianProduct(x, y)
	expected := NewSet[[2]byte]()
	if !Equals(z, expected) {
		t.Errorf("CartesianProduct(%v, %v) == %v, expected %v", x, y, z, expected)
	}

	x.Add('a')
	z = CartesianProduct(x, y)
	if !Equals(z, expected) {
		t.Errorf("CartesianProduct(%v, %v) == %v, expected %v", x, y, z, expected)
	}

	y.Add('b')
	z = CartesianProduct(x, y)
	expected.Add([2]byte{'a', 'b'})
	if !Equals(z, expected) {
		t.Errorf("CartesianProduct(%v, %v) == %v, expected %v", x, y, z, expected)
	}

	x.Add('b')
	y.Add('a')
	z = CartesianProduct(x, y)
	expected.Add([2]byte{'a', 'a'})
	expected.Add([2]byte{'a', 'b'})
	expected.Add([2]byte{'b', 'a'})
	expected.Add([2]byte{'b', 'b'})
	if !Equals(z, expected) {
		t.Errorf("CartesianProduct(%v, %v) == %v, expected %v", x, y, z, expected)
	}
}
