package day5part2

import (
	"strconv"
	"strings"
)

const (
	seedsPrefix                 = "seeds: "
	seedToSoilHeader            = "seed-to-soil map:"
	soilToFertilizerHeader      = "soil-to-fertilizer map:"
	fertilizerToWaterHeader     = "fertilizer-to-water map:"
	waterToLightHeader          = "water-to-light map:"
	lightToTemperatureHeader    = "light-to-temperature map:"
	temperatureToHumidityHeader = "temperature-to-humidity map:"
	humidityToLocationHeader    = "humidity-to-location map:"
)

type Range struct {
	Start  int
	Length int
}

func (r Range) End() int {
	return r.Start + r.Length - 1
}

type Ranges []Range

func (rs Ranges) Equal(other Ranges) bool {
	if len(rs) != len(other) {
		return false
	}
	set := make(map[Range]bool)
	for _, r := range rs {
		set[r] = true
	}
	for _, r := range other {
		if !set[r] {
			return false
		}
	}
	return true
}

type Seeds Ranges

func NewSeeds(line string) Seeds {
	seeds := Seeds{}
	line = strings.TrimRight(line, "\n")
	line = strings.TrimPrefix(line, seedsPrefix)
	split := strings.Split(line, " ")
	for i, j := 0, 1; j < len(split); i, j = i+2, j+2 {
		start, _ := strconv.Atoi(split[i])
		length, _ := strconv.Atoi(split[j])
		seeds = append(seeds, Range{Start: start, Length: length})
	}
	return seeds
}

type RangeMapItem struct {
	DstStart int
	SrcStart int
	Length   int
}

func (rmi RangeMapItem) SrcEnd() int {
	return rmi.SrcStart + rmi.Length - 1
}

func (rmi RangeMapItem) DstEnd() int {
	return rmi.DstStart + rmi.Length - 1
}

func (rmi RangeMapItem) Shift() int {
	return rmi.DstStart - rmi.SrcStart
}

func NewRangeMapItem(input string) RangeMapItem {
	input = strings.TrimRight(input, "\n")
	split := strings.Split(input, " ")
	dstStart, _ := strconv.Atoi(split[0])
	srcStart, _ := strconv.Atoi(split[1])
	length, _ := strconv.Atoi(split[2])
	return RangeMapItem{dstStart, srcStart, length}
}

type RangeMap []RangeMapItem

func NewRangeMap(inputs []string) RangeMap {
	items := make(RangeMap, 0)
	for _, input := range inputs {
		if strings.TrimSpace(input) == "" {
			break
		}
		items = append(items, NewRangeMapItem(input))
	}
	return items
}

func (rm RangeMap) Apply(rs Ranges) Ranges {
	inputs := rs
	outputs := make(Ranges, 0)
	newInputs := make(Ranges, 0)
	for _, item := range rm {
		for _, input := range inputs {
			if input.End() < item.SrcStart || input.Start > item.SrcEnd() {
				newInputs = append(newInputs, input)
			} else if input.Start >= item.SrcStart {
				if input.End() <= item.SrcEnd() {
					outputs = append(outputs, Range{input.Start + item.Shift(), input.Length})
				} else {
					length := item.SrcEnd() - input.Start + 1
					if length > 0 {
						outputs = append(outputs, Range{input.Start + item.Shift(), length})
					}
					if input.Length-length > 0 {
						newInputs = append(newInputs, Range{item.SrcEnd() + 1, input.Length - length})
					}
				}
			} else {
				if input.End() <= item.SrcEnd() {
					length := input.End() - item.SrcStart + 1
					if length > 0 {
						outputs = append(outputs, Range{item.DstStart, length})
					}
					if input.Length-length > 0 {
						newInputs = append(newInputs, Range{input.Start, input.Length - length})
					}
				} else {
					outputs = append(outputs, Range{item.DstStart, item.Length})
					leftLength := item.SrcStart - input.Start
					rightLength := input.End() - item.SrcEnd()
					if leftLength > 0 {
						newInputs = append(newInputs, Range{input.Start, leftLength})
					}
					if rightLength > 0 {
						newInputs = append(newInputs, Range{item.SrcEnd() + 1, rightLength})
					}
				}
			}
		}
		inputs = newInputs
		newInputs = make(Ranges, 0)
	}
	resultSet := make(map[Range]bool)
	for _, output := range outputs {
		resultSet[output] = true
	}
	for _, input := range inputs {
		resultSet[input] = true
	}
	result := make(Ranges, 0)
	for rangeItem := range resultSet {
		result = append(result, rangeItem)
	}
	return result
}

type Atlas struct {
	Seeds Seeds
	Maps  []RangeMap
}

func NewAtlas(inputs []string) Atlas {
	atlas := Atlas{
		Seeds: NewSeeds(inputs[0]),
		Maps:  make([]RangeMap, 0),
	}
	for i := 1; i < len(inputs); {
		if strings.HasPrefix(inputs[i], seedToSoilHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else if strings.HasPrefix(inputs[i], soilToFertilizerHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else if strings.HasPrefix(inputs[i], fertilizerToWaterHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else if strings.HasPrefix(inputs[i], waterToLightHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else if strings.HasPrefix(inputs[i], lightToTemperatureHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else if strings.HasPrefix(inputs[i], temperatureToHumidityHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else if strings.HasPrefix(inputs[i], humidityToLocationHeader) {
			rangeMap := NewRangeMap(inputs[i+1:])
			atlas.Maps = append(atlas.Maps, rangeMap)
			i += len(rangeMap)
		} else {
			i++
		}
	}
	return atlas
}

func (a Atlas) FindLocations() Ranges {
	locations := Ranges(a.Seeds)
	for _, rangeMap := range a.Maps {
		locations = rangeMap.Apply(locations)
	}
	return locations
}

func MinLocation(lines []string) int {
	atlas := NewAtlas(lines)
	locations := atlas.FindLocations()
	min := locations[0].Start
	for _, location := range locations {
		if location.Start < min {
			min = location.Start
		}
	}
	return min
}
