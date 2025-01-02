package day21

import (
	"iter"
	"log"
	"slices"
	"strconv"
)

const (
	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
)

type stack[T any] struct {
	size   int
	values []T
}

func (s *stack[T]) push(value T) {
	s.size++
	s.values = append(s.values, value)
}

func (s *stack[T]) pop() (value T, ok bool) {
	ok = s.size > 0
	if ok {
		s.size--
		value = s.values[s.size]
		s.values = s.values[:s.size]
	}
	return
}

func getNumericPart(input string) int {
	inputBytes := []byte(input)
	numericString := string(inputBytes[:len(input)-1])
	if i, err := strconv.Atoi(numericString); err == nil {
		return i
	}
	log.Panicf("Could not parse numeric part from input %q", input)
	return 0
}

type vector struct {
	x, y int
}

type grid struct {
	nrows, ncols int
	values       []byte
}

func newGrid(nrows, ncols int) *grid {
	values := make([]byte, nrows*ncols)
	return &grid{
		nrows:  nrows,
		ncols:  ncols,
		values: values,
	}
}

func (g *grid) get(v vector) (byte, bool) {
	if v.x < 0 || v.x >= g.ncols || v.y < 0 || v.y >= g.nrows {
		return 0, false
	}
	return g.values[v.y*g.ncols+v.x], true
}

func (g *grid) set(v vector, c byte) {
	g.values[v.y*g.ncols+v.x] = c
}

func (g *grid) all() iter.Seq2[vector, byte] {
	return func(yield func(vector, byte) bool) {
		for i, c := range g.values {
			row := i / g.ncols
			col := i % g.ncols
			v := vector{col, row}
			if !yield(v, c) {
				break
			}
		}
	}
}

func (g grid) String() string {
	s := []byte{}
	for i, c := range g.values {
		s = append(s, c)
		if i%g.ncols == g.ncols-1 {
			s = append(s, '\n')
		}
	}
	return string(s)
}

type numericKeyPad struct {
	layout *grid
	keyMap map[byte]vector
}

/*
 * +---+---+---+
 * | 7 | 8 | 9 |
 * +---+---+---+
 * | 4 | 5 | 6 |
 * +---+---+---+
 * | 1 | 2 | 3 |
 * +---+---+---+
 *     | 0 | A |
 *     +---+---+
 */
func newNumericKeyPad() numericKeyPad {
	layout := newGrid(4, 3)

	layout.set(vector{0, 0}, '7')
	layout.set(vector{1, 0}, '8')
	layout.set(vector{2, 0}, '9')

	layout.set(vector{0, 1}, '4')
	layout.set(vector{1, 1}, '5')
	layout.set(vector{2, 1}, '6')

	layout.set(vector{0, 2}, '1')
	layout.set(vector{1, 2}, '2')
	layout.set(vector{2, 2}, '3')

	layout.set(vector{0, 3}, 0)
	layout.set(vector{1, 3}, '0')
	layout.set(vector{2, 3}, 'A')

	keyMap := map[byte]vector{}
	for v, c := range layout.all() {
		keyMap[c] = v
	}

	return numericKeyPad{layout, keyMap}
}

func (k numericKeyPad) getDirectionalSequences(from, to byte) []string {
	finalSeqs := []string{}
	seqs := new(stack[struct {
		from, to byte
		seq      []byte
	}])
	seqs.push(struct {
		from byte
		to   byte
		seq  []byte
	}{from, to, []byte{}})
	for {
		if current, ok := seqs.pop(); ok {
			seq := current.seq
			from := current.from
			to := current.to
			if from == to {
				finalSeqs = append(finalSeqs, string(append(seq, 'A')))
				continue
			}
			fromVec := k.keyMap[from]
			toVec := k.keyMap[to]
			xDiff := toVec.x - fromVec.x
			yDiff := toVec.y - fromVec.y
			if xDiff < 0 {
				if newFrom, ok := k.layout.get(vector{fromVec.x - 1, fromVec.y}); ok && newFrom != 0 {
					seqs.push(struct {
						from byte
						to   byte
						seq  []byte
					}{
						from: newFrom,
						to:   to,
						seq:  append(slices.Clone(seq), '<'),
					})
				}
			} else if xDiff > 0 {
				if newFrom, ok := k.layout.get(vector{fromVec.x + 1, fromVec.y}); ok && newFrom != 0 {
					seqs.push(struct {
						from byte
						to   byte
						seq  []byte
					}{
						from: newFrom,
						to:   to,
						seq:  append(slices.Clone(seq), '>'),
					})
				}
			}
			if yDiff < 0 {
				if newFrom, ok := k.layout.get(vector{fromVec.x, fromVec.y - 1}); ok && newFrom != 0 {
					seqs.push(struct {
						from byte
						to   byte
						seq  []byte
					}{
						from: newFrom,
						to:   to,
						seq:  append(slices.Clone(seq), '^'),
					})
				}
			} else if yDiff > 0 {
				if newFrom, ok := k.layout.get(vector{fromVec.x, fromVec.y + 1}); ok && newFrom != 0 {
					seqs.push(struct {
						from byte
						to   byte
						seq  []byte
					}{
						from: newFrom,
						to:   to,
						seq:  append(slices.Clone(seq), 'v'),
					})
				}
			}
		} else {
			break
		}
	}
	return finalSeqs
}

type directionalKeyPad struct {
	layout               *grid
	keyMap               map[byte]vector
	directionalSequences map[string]int
	expansionMap         [][]int
}

/*
 *     +---+---+
 *     | ^ | A |
 * +---+---+---+
 * | < | v | > |
 * +---+---+---+
 */
func newDirectionalKeyPad() directionalKeyPad {
	layout := newGrid(2, 3)

	layout.set(vector{0, 0}, 0)
	layout.set(vector{1, 0}, '^')
	layout.set(vector{2, 0}, 'A')

	layout.set(vector{0, 1}, '<')
	layout.set(vector{1, 1}, 'v')
	layout.set(vector{2, 1}, '>')

	keyMap := map[byte]vector{}
	for v, c := range layout.all() {
		keyMap[c] = v
	}

	keyPad := directionalKeyPad{
		layout: layout,
		keyMap: keyMap,
	}
	keyPad.setDirectionalSequences()
	keyPad.setExpansionMap()
	return keyPad
}

func (k directionalKeyPad) getDirectionalSequence(from, to byte) string {
	seq := []byte{}
	fromVec := k.keyMap[from]
	toVec := k.keyMap[to]
	xDiff := toVec.x - fromVec.x
	yDiff := toVec.y - fromVec.y
	switch {
	case xDiff < 0 && yDiff < 0:
		for range -yDiff {
			seq = append(seq, '^')
		}
		for range -xDiff {
			seq = append(seq, '<')
		}
	case xDiff < 0 && yDiff >= 0:
		for range yDiff {
			seq = append(seq, 'v')
		}
		for range -xDiff {
			seq = append(seq, '<')
		}
	case xDiff >= 0 && yDiff < 0:
		for range xDiff {
			seq = append(seq, '>')
		}
		for range -yDiff {
			seq = append(seq, '^')
		}
	case xDiff >= 0 && yDiff >= 0:
		for range xDiff {
			seq = append(seq, '>')
		}
		for range yDiff {
			seq = append(seq, 'v')
		}
	}
	seq = append(seq, 'A')
	return string(seq)
}

func (k *directionalKeyPad) setDirectionalSequences() {
	k.directionalSequences = map[string]int{}
	index := 0
	for xi := range k.layout.ncols {
		for yi := range k.layout.nrows {
			vi := vector{xi, yi}
			if ci, _ := k.layout.get(vi); ci != 0 {
				for xj := range k.layout.ncols {
					for yj := range k.layout.nrows {
						vj := vector{xj, yj}
						if cj, _ := k.layout.get(vj); cj != 0 {
							if _, ok := k.directionalSequences[k.getDirectionalSequence(ci, cj)]; !ok {
								k.directionalSequences[k.getDirectionalSequence(ci, cj)] = index
								index++
							}
						}
					}
				}
			}
		}
	}
}

func (k *directionalKeyPad) setExpansionMap() {
	k.expansionMap = make([][]int, len(k.directionalSequences))
	for directionalSequence, i := range k.directionalSequences {
		k.expansionMap[i] = []int{}
		var from byte = 'A'
		for j := range directionalSequence {
			to := directionalSequence[j]
			d := k.getDirectionalSequence(from, to)
			l := k.directionalSequences[d]
			k.expansionMap[i] = append(k.expansionMap[i], l)
			from = to
		}
	}
}

func getShortestSequenceLength(input string, nDirectionalKeypads int) int {
	numericKeyPad := newNumericKeyPad()
	directionalKeyPad := newDirectionalKeyPad()
	depth := nDirectionalKeypads - 1

	shortestSequenceLengths := make([][]int, depth)
	shortestSequenceLengths[0] = make([]int, len(directionalKeyPad.directionalSequences))
	for directionalSequence, j := range directionalKeyPad.directionalSequences {
		shortestSequenceLengths[0][j] = len(directionalSequence)
	}
	for i := 1; i < depth; i++ {
		shortestSequenceLengths[i] = make([]int, len(directionalKeyPad.directionalSequences))
		for j := range shortestSequenceLengths[i] {
			for _, k := range directionalKeyPad.expansionMap[j] {
				shortestSequenceLengths[i][j] += shortestSequenceLengths[i-1][k]
			}
		}
	}

	finalSum := 0
	var fromNumeric byte = 'A'
	for i := range input {
		minSum := maxInt
		toNumeric := input[i]
		for _, initialDirectionalSequence := range numericKeyPad.getDirectionalSequences(fromNumeric, toNumeric) {
			sum := 0
			var fromDirectional byte = 'A'
			for j := range initialDirectionalSequence {
				toDirectional := initialDirectionalSequence[j]
				directionalSequence := directionalKeyPad.getDirectionalSequence(fromDirectional, toDirectional)
				index := directionalKeyPad.directionalSequences[directionalSequence]
				sum += shortestSequenceLengths[depth-1][index]
				fromDirectional = toDirectional
			}
			if sum < minSum {
				minSum = sum
			}
			fromNumeric = toNumeric
		}
		finalSum += minSum
	}
	return finalSum
}

func CalcComplexity(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += getNumericPart(input) * getShortestSequenceLength(input, 3)
	}
	return sum
}
