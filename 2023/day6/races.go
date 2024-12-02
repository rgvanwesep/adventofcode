package day6

import (
	"strconv"
	"strings"
)

const (
	timePrefix     = "Time:"
	distancePrefix = "Distance:"
)

type Race struct {
	Time     uint64
	Distance uint64
}

func NewRace(inputs []string) Race {
	race := Race{}
	var (
		time     uint64
		distance uint64
	)
	for _, input := range inputs {
		if strings.HasPrefix(input, timePrefix) {
			input = strings.TrimPrefix(input, timePrefix)
			input = strings.TrimSpace(input)
			timeStr := strings.Join(strings.Split(input, " "), "")
			for _, digit := range timeStr {
				time *= 10
				time += uint64(digit - '0')
			}
			race.Time = time
		} else if strings.HasPrefix(input, distancePrefix) {
			input = strings.TrimPrefix(input, distancePrefix)
			input = strings.TrimSpace(input)
			distanceStr := strings.Join(strings.Split(input, " "), "")
			for _, digit := range distanceStr {
				distance *= 10
				distance += uint64(digit - '0')
			}
			race.Distance = distance
		}
	}
	return race
}

func (r Race) CanWin(impulse uint64) bool {
	if impulse == 0 {
		return false
	}
	remainingTime := r.Time - impulse
	if remainingTime <= 0 {
		return false
	}
	return (impulse * remainingTime) > r.Distance
}

func (r Race) NumWins() uint64 {
	var wins uint64 = 0
	for i := uint64(1); i < r.Time; i++ {
		if r.CanWin(i) {
			wins++
		}
	}
	return wins
}

type Races []Race

func NewRaces(inputs []string) Races {
	times := make([]int, 0)
	distances := make([]int, 0)
	for _, input := range inputs {
		input = strings.TrimSpace(input)
		if strings.HasPrefix(input, timePrefix) {
			split := strings.Split(input, timePrefix)
			split = strings.Split(split[1], " ")
			for _, s := range split {
				if s != "" {
					time, _ := strconv.Atoi(s)
					times = append(times, time)
				}
			}
		} else if strings.HasPrefix(input, distancePrefix) {
			split := strings.Split(input, distancePrefix)
			split = strings.Split(split[1], " ")
			for _, s := range split {
				if s != "" {
					distance, _ := strconv.Atoi(s)
					distances = append(distances, distance)
				}
			}
		}
	}
	races := make(Races, 0)
	for i, time := range times {
		races = append(races, Race{Time: uint64(time), Distance: uint64(distances[i])})
	}
	return races
}

func Product(inputs []string) uint64 {
	races := NewRaces(inputs)
	var product uint64 = 1
	for _, race := range races {
		product *= race.NumWins()
	}
	return product
}

func Count(inputs []string) uint64 {
	race := NewRace(inputs)
	return race.NumWins()
}
