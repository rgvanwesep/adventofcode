package day3

import "strings"

const (
	zero  = byte('0')
	one   = byte('1')
	two   = byte('2')
	three = byte('3')
	four  = byte('4')
	five  = byte('5')
	six   = byte('6')
	seven = byte('7')
	eight = byte('8')
	nine  = byte('9')
	dot   = byte('.')
)

type Coordinate struct {
	Row int
	Col int
}

func (c *Coordinate) Neighbors() map[Coordinate]bool {
	neighbors := make(map[Coordinate]bool)
	neighbors[Coordinate{Row: c.Row - 1, Col: c.Col - 1}] = true
	neighbors[Coordinate{Row: c.Row - 1, Col: c.Col}] = true
	neighbors[Coordinate{Row: c.Row - 1, Col: c.Col + 1}] = true
	neighbors[Coordinate{Row: c.Row, Col: c.Col - 1}] = true
	neighbors[Coordinate{Row: c.Row, Col: c.Col + 1}] = true
	neighbors[Coordinate{Row: c.Row + 1, Col: c.Col - 1}] = true
	neighbors[Coordinate{Row: c.Row + 1, Col: c.Col}] = true
	neighbors[Coordinate{Row: c.Row + 1, Col: c.Col + 1}] = true
	return neighbors
}

type CoordinateRange struct {
	Start Coordinate
	End   Coordinate
}

func (r *CoordinateRange) Extend(c Coordinate) {
	// Update the range coordinates
	if c.Row < r.Start.Row {
		r.Start.Row = c.Row
	}
	if c.Row > r.End.Row {
		r.End.Row = c.Row
	}
	if c.Col < r.Start.Col {
		r.Start.Col = c.Col
	}
	if c.Col > r.End.Col {
		r.End.Col = c.Col
	}
}

type PartNumber struct {
	Value int
	Range CoordinateRange
}

func (p *PartNumber) AppendDigit(digit int, coordinate Coordinate) {
	p.Value = p.Value*10 + digit
	p.Range.Extend(coordinate)
}

type Symbol struct {
	Value    byte
	Position Coordinate
}

type Schematic struct {
	PartNumbers map[Coordinate]PartNumber
	Symbols     map[Coordinate]Symbol
}

func NewSchematic(inputs []string) *Schematic {
	schematic := &Schematic{
		PartNumbers: make(map[Coordinate]PartNumber),
		Symbols:     make(map[Coordinate]Symbol),
	}

	for row, input := range inputs {
		input = strings.TrimRight(input, "\n")
		for col, char := range []byte(input) {
			coordinate := Coordinate{Row: row, Col: col}
			switch char {
			case zero:
				schematic.AddDigit(0, coordinate)
			case one:
				schematic.AddDigit(1, coordinate)
			case two:
				schematic.AddDigit(2, coordinate)
			case three:
				schematic.AddDigit(3, coordinate)
			case four:
				schematic.AddDigit(4, coordinate)
			case five:
				schematic.AddDigit(5, coordinate)
			case six:
				schematic.AddDigit(6, coordinate)
			case seven:
				schematic.AddDigit(7, coordinate)
			case eight:
				schematic.AddDigit(8, coordinate)
			case nine:
				schematic.AddDigit(9, coordinate)
			case dot:
				continue
			default:
				schematic.Symbols[coordinate] = Symbol{
					Value:    char,
					Position: coordinate,
				}
			}
		}
	}

	return schematic
}

func (s *Schematic) AddDigit(digit int, coordinate Coordinate) {
	prevCoordinate := Coordinate{Row: coordinate.Row, Col: coordinate.Col - 1}
	partNumber, ok := s.PartNumbers[prevCoordinate]
	if ok {
		partNumber.AppendDigit(digit, coordinate)
		for i := partNumber.Range.Start.Col; i <= partNumber.Range.End.Col; i++ {
			s.PartNumbers[Coordinate{Row: partNumber.Range.Start.Row, Col: i}] = partNumber
		}
	} else {
		s.PartNumbers[coordinate] = PartNumber{
			Value: digit,
			Range: CoordinateRange{Start: coordinate, End: coordinate},
		}
	}
}

func (s *Schematic) Equals(other *Schematic) bool {
	if len(s.PartNumbers) != len(other.PartNumbers) {
		return false
	}

	for coordinate, partNumber := range s.PartNumbers {
		otherPartNumber, ok := other.PartNumbers[coordinate]
		if !ok {
			return false
		}
		if partNumber != otherPartNumber {
			return false
		}
	}

	if len(s.Symbols) != len(other.Symbols) {
		return false
	}

	for coordinate, symbol := range s.Symbols {
		otherSymbol, ok := other.Symbols[coordinate]
		if !ok {
			return false
		}
		if symbol != otherSymbol {
			return false
		}
	}

	return true
}

func (s *Schematic) Diff(other *Schematic) map[Coordinate]bool {
	diff := make(map[Coordinate]bool)

	for coordinate, partNumber := range s.PartNumbers {
		otherPartNumber, ok := other.PartNumbers[coordinate]
		if !ok {
			diff[coordinate] = true
		}
		if partNumber != otherPartNumber {
			diff[coordinate] = true
		}
	}

	for coordinate, partNumber := range other.PartNumbers {
		otherPartNumber, ok := s.PartNumbers[coordinate]
		if !ok {
			diff[coordinate] = true
		}
		if partNumber != otherPartNumber {
			diff[coordinate] = true
		}
	}

	for coordinate, symbol := range s.Symbols {
		otherSymbol, ok := other.Symbols[coordinate]
		if !ok {
			diff[coordinate] = true
			continue
		}
		if symbol != otherSymbol {
			diff[coordinate] = true

		}
	}

	for coordinate, symbol := range other.Symbols {
		_, ok := s.Symbols[coordinate]
		if !ok {
			diff[coordinate] = true
			continue
		}
		if symbol != s.Symbols[coordinate] {
			diff[coordinate] = true
		}
	}

	return diff
}

func (s *Schematic) AdjacentPartNumbers(coordinate Coordinate) map[PartNumber]bool {
	partNumbers := make(map[PartNumber]bool)
	for neighbor := range coordinate.Neighbors() {
		partNumber, ok := s.PartNumbers[neighbor]
		if !ok {
			continue
		}
		partNumbers[partNumber] = true
	}
	return partNumbers
}

func Sum(inputs []string) int {
	schematic := NewSchematic(inputs)
	sum := 0
	for _, symbol := range schematic.Symbols {
		partNumbers := schematic.AdjacentPartNumbers(symbol.Position)
		for partNumber := range partNumbers {
			sum += partNumber.Value
		}
	}
	return sum
}

func SumGearRatios(inputs []string) int {
	schematic := NewSchematic(inputs)
	sum := 0
	for _, symbol := range schematic.Symbols {
		if symbol.Value == '*' {
			partNumbers := schematic.AdjacentPartNumbers(symbol.Position)
			if len(partNumbers) == 2 {
				product := 1
				for partNumber := range partNumbers {
					product *= partNumber.Value
				}
				sum += product
			}
		}
	}
	return sum
}
