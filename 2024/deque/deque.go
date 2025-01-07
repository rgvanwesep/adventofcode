package deque

import (
	"iter"
)

type Deque[T any] interface {
	All() iter.Seq2[int, T]
	Append(T) bool
	AppendLeft(T) bool
	Cap() int
	Clear()
	Clone() Deque[T]
	Extend(iter.Seq[T]) bool
	ExtendLeft(iter.Seq[T]) bool
	Len() int
	Pop() (T, bool)
	PopLeft() (T, bool)
	Rotate(int)
}

const (
	blockLen = 64
	center   = (blockLen - 1) / 2
)

type block[T any] struct {
	data        [blockLen]T
	left, right *block[T]
}

type dequeDLL[T any] struct {
	length, capacity, maxLen int
	leftIndex, rightIndex    int
	leftBlock, rightBlock    *block[T]
}

func (d *dequeDLL[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		leftIndex := d.leftIndex
		leftBlock := d.leftBlock
		for i := 0; i < d.length; i++ {
			value := leftBlock.data[leftIndex]
			if !yield(i, value) {
				break
			}
			leftIndex++

			if leftIndex == blockLen {
				leftIndex = 0
				leftBlock = leftBlock.right
			}
		}
	}
}

func (d *dequeDLL[T]) Append(value T) bool {
	var b *block[T]
	if d.length == d.maxLen {
		return false
	}
	if d.rightIndex == blockLen-1 {
		b = new(block[T])
		b.left = d.rightBlock
		d.rightBlock.right = b
		d.rightBlock = b
		d.rightIndex = -1
		d.capacity += blockLen
	}
	d.rightIndex++
	d.rightBlock.data[d.rightIndex] = value
	d.length++
	return true
}

func (d *dequeDLL[T]) AppendLeft(value T) bool {
	var b *block[T]
	if d.length == d.maxLen {
		return false
	}
	if d.leftIndex == 0 {
		b = new(block[T])
		b.right = d.leftBlock
		d.leftBlock.left = b
		d.leftBlock = b
		d.leftIndex = blockLen
		d.capacity += blockLen
	}
	d.leftIndex--
	d.leftBlock.data[d.leftIndex] = value
	d.length++
	return true
}

func (d *dequeDLL[T]) Cap() int {
	return d.capacity
}

func (d *dequeDLL[T]) Clear() {
	d.leftBlock = d.rightBlock
	d.leftIndex = center + 1
	d.rightIndex = center
	d.length = 0
	d.capacity = blockLen
}

func (d *dequeDLL[T]) Clone() Deque[T] {
	dClone := new(dequeDLL[T])
	b := new(block[T])
	b.data = d.leftBlock.data
	dClone.leftBlock = b
	dClone.rightBlock = b
	for last, current := b, d.leftBlock.right; current != nil; last, current = current, current.right {
		b = new(block[T])
		last.right = b
		b.data = current.data
		b.left = last
		dClone.rightBlock = b
	}
	dClone.maxLen = d.maxLen
	dClone.length = d.length
	dClone.capacity = d.capacity
	dClone.leftIndex = d.leftIndex
	dClone.rightIndex = d.rightIndex
	return dClone
}

func (d *dequeDLL[T]) Extend(values iter.Seq[T]) bool {
	var ok bool
	for v := range values {
		ok = d.Append(v)
		if !ok {
			break
		}
	}
	return ok
}

func (d *dequeDLL[T]) ExtendLeft(values iter.Seq[T]) bool {
	var ok bool
	for v := range values {
		ok = d.AppendLeft(v)
		if !ok {
			break
		}
	}
	return ok
}

func (d *dequeDLL[T]) Len() int {
	return d.length
}

func (d *dequeDLL[T]) Pop() (T, bool) {
	var value T
	if d.length == 0 {
		return value, false
	}
	value = d.rightBlock.data[d.rightIndex]
	d.rightIndex--
	d.length--

	if d.rightIndex < 0 {
		if d.length > 0 {
			d.rightBlock = d.rightBlock.left
			d.rightBlock.right = nil
			d.rightIndex = blockLen - 1
			d.capacity -= blockLen
		} else {
			d.leftIndex = center + 1
			d.rightIndex = center
		}
	}
	return value, true
}

func (d *dequeDLL[T]) PopLeft() (T, bool) {
	var value T
	if d.length == 0 {
		return value, false
	}
	value = d.leftBlock.data[d.leftIndex]
	d.leftIndex++
	d.length--

	if d.leftIndex == blockLen {
		if d.length > 0 {
			d.leftBlock = d.leftBlock.right
			d.leftBlock.left = nil
			d.leftIndex = 0
			d.capacity -= blockLen
		} else {
			d.leftIndex = center + 1
			d.rightIndex = center
		}
	}
	return value, true
}

func (d *dequeDLL[T]) Rotate(n int) {
	if d.length == 0 {
		return
	}
	if n > 0 {
		for range n {
			v, _ := d.Pop()
			d.AppendLeft(v)
		}
	} else {
		for range -n {
			v, _ := d.PopLeft()
			d.Append(v)
		}
	}
}

func NewDeque[T any](maxLen int) Deque[T] {
	d := new(dequeDLL[T])
	b := new(block[T])
	d.leftBlock = b
	d.rightBlock = b
	d.leftIndex = center + 1
	d.rightIndex = center
	d.maxLen = maxLen
	d.length = 0
	d.capacity = blockLen
	return d
}
