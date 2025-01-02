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

type keyPad struct {
	layout *grid
	keyMap map[byte]vector
}

func (k keyPad) getDirectionalSequences(from, to byte) []string {
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

type numericKeyPad struct {
	keyPad
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

	k := new(numericKeyPad)
	k.layout = layout
	k.keyMap = keyMap
	return *k
}

type directionalKeyPad struct {
	keyPad
	directionalSequences map[string]int
	expansionMap         [][][]int
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

	k := new(directionalKeyPad)
	k.layout = layout
	k.keyMap = keyMap
	k.setDirectionalSequences()
	k.setExpansionMap()
	return *k
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
							for _, seq := range k.getDirectionalSequences(ci, cj) {
								if _, ok := k.directionalSequences[seq]; !ok {
									k.directionalSequences[seq] = index
									index++
								}
							}
						}
					}
				}
			}
		}
	}
}

func (k *directionalKeyPad) setExpansionMap() {
	k.expansionMap = make([][][]int, len(k.directionalSequences))
	for directionalSequence, i := range k.directionalSequences {
		k.expansionMap[i] = [][]int{}
		s := new(stack[struct {
			index    int
			from, to byte
			indices  []int
		}])
		var from byte = 'A'
		to := directionalSequence[0]
		for _, seq := range k.getDirectionalSequences(from, to) {
			l := k.directionalSequences[seq]
			s.push(struct {
				index   int
				from    byte
				to      byte
				indices []int
			}{
				0,
				from,
				to,
				[]int{l},
			})
		}
		for {
			if current, ok := s.pop(); ok {
				if current.index < len(directionalSequence)-1 {
					index := current.index + 1
					from := current.to
					to := directionalSequence[index]
					for _, seq := range k.getDirectionalSequences(from, to) {
						l := k.directionalSequences[seq]
						s.push(struct {
							index   int
							from    byte
							to      byte
							indices []int
						}{
							index,
							from,
							to,
							append(slices.Clone(current.indices), l),
						})
					}
				} else {
					k.expansionMap[i] = append(k.expansionMap[i], current.indices)
				}
			} else {
				break
			}
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
			minSum := maxInt
			for _, ks := range directionalKeyPad.expansionMap[j] {
				sum := 0
				for _, k := range ks {
					sum += shortestSequenceLengths[i-1][k]
				}
				if sum < minSum {
					minSum = sum
				}
			}
			shortestSequenceLengths[i][j] = minSum
		}
	}

	finalSum := 0
	var fromNumeric byte = 'A'
	for i := range input {
		minSum := maxInt
		toNumeric := input[i]
		for _, initialDirectionalSequence := range numericKeyPad.getDirectionalSequences(fromNumeric, toNumeric) {
			index := 0
			var fromDirectional byte = 'A'
			toDirectional := initialDirectionalSequence[index]
			s := new(stack[struct {
				index    int
				from, to byte
				sum      int
			}])
			for _, directionalSequence := range directionalKeyPad.getDirectionalSequences(fromDirectional, toDirectional) {
				s.push(struct {
					index int
					from  byte
					to    byte
					sum   int
				}{
					index,
					fromDirectional,
					toDirectional,
					shortestSequenceLengths[depth-1][directionalKeyPad.directionalSequences[directionalSequence]],
				})
			}
			for {
				if current, ok := s.pop(); ok {
					if current.index < len(initialDirectionalSequence)-1 {
						index := current.index + 1
						from := current.to
						to := initialDirectionalSequence[index]
						for _, directionalSequence := range directionalKeyPad.getDirectionalSequences(from, to) {
							s.push(struct {
								index int
								from  byte
								to    byte
								sum   int
							}{
								index,
								from,
								to,
								current.sum + shortestSequenceLengths[depth-1][directionalKeyPad.directionalSequences[directionalSequence]],
							})
						}
					} else {
						if current.sum < minSum {
							minSum = current.sum
						}
					}
				} else {
					break
				}
			}
		}
		finalSum += minSum
		fromNumeric, toNumeric = toNumeric, fromNumeric
	}
	return finalSum
}

func CalcComplexity(inputs []string, nDirectionalKeypads int) int {
	sum := 0
	for _, input := range inputs {
		sum += getNumericPart(input) * getShortestSequenceLength(input, nDirectionalKeypads)
	}
	return sum
}
