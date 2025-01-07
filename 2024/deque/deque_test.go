package deque

import (
	"slices"
	"testing"
)

func TestDeque(t *testing.T) {
	manyBytes := make([]byte, blockLen/2+1)
	for i := range manyBytes {
		b := byte(int('0') + i)
		manyBytes[i] = b
	}
	cases := []struct {
		values string
	}{
		{
			values: "ab",
		},
		{
			values: string(manyBytes[:blockLen/2]),
		},
		{
			values: string(manyBytes),
		},
	}
	for _, c := range cases {
		var (
			i          int
			b          byte
			bs         []byte
			s, reverse string
			ok         bool
		)
		bs = []byte(c.values)
		slices.Reverse(bs)
		reverse = string(bs)

		d := NewDeque[byte](0)
		ok = d.Append(0)
		if ok {
			t.Errorf("Append(0) == %t, expected false", ok)
		}
		ok = d.AppendLeft(0)
		if ok {
			t.Errorf("AppendLeft(0) == %t, expected false", ok)
		}
		ok = d.Extend(func(yield func(byte) bool) { yield(0) })
		if ok {
			t.Errorf("Extend(0) == %t, expected false", ok)
		}
		ok = d.ExtendLeft(func(yield func(byte) bool) { yield(0) })
		if ok {
			t.Errorf("ExtendLeft(0) == %t, expected false", ok)
		}

		d = NewDeque[byte](len(c.values))

		i = d.Len()
		if i != 0 {
			t.Errorf("Len() == %d, expected 0", i)
		}
		i = d.Cap()
		if i != blockLen {
			t.Errorf("Cap() == %d, expected %d", i, blockLen)
		}
		b, ok = d.Pop()
		if ok || b != 0 {
			t.Errorf("Pop() == (%q, %t), expected (0, false)", b, ok)
		}
		b, ok = d.PopLeft()
		if ok || b != 0 {
			t.Errorf("PopLeft() == (%q, %t), expected (0, false)", b, ok)
		}

		d.Rotate(1)
		i = d.Len()
		if i != 0 {
			t.Errorf("Len() == %d, expected 0", i)
		}
		i = d.Cap()
		if i != blockLen {
			t.Errorf("Cap() == %d, expected %d", i, blockLen)
		}
		b, ok = d.Pop()
		if ok || b != 0 {
			t.Errorf("Pop() == (%q, %t), expected (0, false)", b, ok)
		}
		b, ok = d.PopLeft()
		if ok || b != 0 {
			t.Errorf("PopLeft() == (%q, %t), expected (0, false)", b, ok)
		}

		dClone := d.Clone()

		i = dClone.Len()
		if i != 0 {
			t.Errorf("Len() == %d, expected 0", i)
		}
		i = dClone.Cap()
		if i != blockLen {
			t.Errorf("Cap() == %d, expected %d", i, blockLen)
		}
		b, ok = dClone.Pop()
		if ok || b != 0 {
			t.Errorf("Pop() == (%q, %t), expected (0, false)", b, ok)
		}
		b, ok = dClone.PopLeft()
		if ok || b != 0 {
			t.Errorf("PopLeft() == (%q, %t), expected (0, false)", b, ok)
		}
		bs = make([]byte, d.Len())
		for i, b = range d.All() {
			bs[i] = b
		}
		s = string(bs)
		if s != "" {
			t.Errorf("All() == %q, expected \"\"", s)
		}

		ok = d.Append(c.values[0])
		if !ok {
			t.Errorf("Append(%q) == %t, expected true", c.values[0], ok)
		}
		b, ok = d.Pop()
		if !ok || b != c.values[0] {
			t.Errorf("Pop() == (%q, %t), expected (%q, true)", b, ok, c.values[0])
		}
		i = d.Len()
		if i != 0 {
			t.Errorf("Len() == %d, expected 0", i)
		}
		i = d.Cap()
		if i != blockLen {
			t.Errorf("Cap() == %d, expected %d", i, blockLen)
		}

		ok = d.AppendLeft(c.values[0])
		if !ok {
			t.Errorf("AppendLeft(%q) == %t, expected true", c.values[0], ok)
		}
		b, ok = d.PopLeft()
		if !ok || b != c.values[0] {
			t.Errorf("PopLeft() == (%q, %t), expected (%q, true)", b, ok, c.values[0])
		}
		i = d.Len()
		if i != 0 {
			t.Errorf("Len() == %d, expected 0", i)
		}
		i = d.Cap()
		if i != blockLen {
			t.Errorf("Cap() == %d, expected %d", i, blockLen)
		}

		ok = d.Extend(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		if !ok {
			t.Errorf("Extend(%q) == %t, expected true", c.values, ok)
		}
		i = d.Len()
		if i != len(c.values) {
			t.Errorf("Len() == %d, expected %d", i, len(c.values))
		}
		bs = make([]byte, d.Len())
		for i, b = range d.All() {
			bs[i] = b
		}
		s = string(bs)
		if s != c.values {
			t.Errorf("All() == %q, expected %q", s, c.values)
		}
		dClone = d.Clone()
		for i = len(c.values) - 1; i >= 0; i-- {
			b, ok = d.Pop()
			if !ok || b != c.values[i] {
				t.Errorf("Pop() == (%q, %t), expected (%q, true)", b, ok, c.values[i])
			}
		}

		i = dClone.Len()
		if i != len(c.values) {
			t.Errorf("Len() == %d, expected %d", i, len(c.values))
		}
		bs = make([]byte, dClone.Len())
		for i, b = range dClone.All() {
			bs[i] = b
		}
		s = string(bs)
		if s != c.values {
			t.Errorf("All() == %q, expected %q", s, c.values)
		}
		for i = 0; i < len(c.values); i++ {
			b, ok = dClone.PopLeft()
			if !ok || b != c.values[i] {
				t.Errorf("PopLeft() == (%q, %t), expected (%q, true)", b, ok, c.values[i])
			}
		}

		ok = d.ExtendLeft(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		if !ok {
			t.Errorf("ExtendLeft(%q) == %t, expected true", c.values, ok)
		}
		i = d.Len()
		if i != len(c.values) {
			t.Errorf("Len() == %d, expected %d", i, len(c.values))
		}
		bs = make([]byte, d.Len())
		for i, b = range d.All() {
			bs[i] = b
		}
		s = string(bs)
		if s != reverse {
			t.Errorf("All() == %q, expected %q", s, reverse)
		}
		dClone = d.Clone()
		for i = len(c.values) - 1; i >= 0; i-- {
			b, ok = d.PopLeft()
			if !ok || b != c.values[i] {
				t.Errorf("PopLeft() == (%q, %t), expected (%q, true)", b, ok, c.values[i])
			}
		}
		for i = 0; i < len(c.values); i++ {
			b, ok = dClone.Pop()
			if !ok || b != c.values[i] {
				t.Errorf("Pop() == (%q, %t), expected (%q, true)", b, ok, c.values[i])
			}
		}

		d.Extend(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		d.Clear()
		i = d.Len()
		if i != 0 {
			t.Errorf("Len() == %d, expected 0", i)
		}
		i = d.Cap()
		if i != blockLen {
			t.Errorf("Cap() == %d, expected %d", i, blockLen)
		}
		b, ok = d.Pop()
		if ok || b != 0 {
			t.Errorf("Pop() == (%q, %t), expected (0, false)", b, ok)
		}
		b, ok = d.PopLeft()
		if ok || b != 0 {
			t.Errorf("PopLeft() == (%q, %t), expected (0, false)", b, ok)
		}

		d = NewDeque[byte](2 * len(c.values))

		d.Extend(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		d.ExtendLeft(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		for i = len(c.values) - 1; i >= 0; i-- {
			b, ok = d.Pop()
			if !ok || b != c.values[i] {
				t.Errorf("Pop() == (%q, %t), expected (%q, true)", b, ok, c.values[i])
			}
		}
		for i = 0; i < len(c.values); i++ {
			b, ok = d.Pop()
			if !ok || b != c.values[i] {
				t.Errorf("Pop() == (%q, %t), expected (%q, true)", b, ok, c.values[i])
			}
		}
		d.Clear()

		d.Extend(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		dClone = d.Clone()
		d.Rotate(1)
		v, _ := dClone.Pop()
		dClone.AppendLeft(v)
		bs = make([]byte, d.Len())
		for i, v := range d.All() {
			bs[i] = v
		}
		for i, v := range dClone.All() {
			if bs[i] != v {
				t.Error("Rotate(1) not equivalent to AppendLeft(Pop())")
			}
		}
		d.Clear()
		dClone.Clear()

		d.Extend(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		dClone = d.Clone()
		d.Rotate(-1)
		v, _ = dClone.PopLeft()
		dClone.Append(v)
		bs = make([]byte, d.Len())
		for i, v := range d.All() {
			bs[i] = v
		}
		for i, v := range dClone.All() {
			if bs[i] != v {
				t.Error("Rotate(-1) not equivalent to Append(PopLeft())")
			}
		}
		d.Clear()
		dClone.Clear()

		d.Extend(func(yield func(byte) bool) {
			for i := range c.values {
				if !yield(c.values[i]) {
					break
				}
			}
		})
		for range d.All() {
			break
		}
	}
}
