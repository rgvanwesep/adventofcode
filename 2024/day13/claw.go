package day13

import (
	"log"
	"regexp"
	"strconv"
)

const (
	aCost         = 3
	bCost         = 1
	buttonPattern = `^Button (A|B): X\+(\d+), Y\+(\d+)$`
	prizePattern  = `^Prize: X=(\d+), Y=(\d+)$`
)

type vector struct {
	x, y int
}

type machine struct {
	aVec, bVec, prizeVec vector
}

func parseMachines(inputs []string) []*machine {
	buttonMatcher := regexp.MustCompile(buttonPattern)
	prizeMatcher := regexp.MustCompile(prizePattern)
	machines := []*machine{}
	for i := 0; i < len(inputs); i += 4 {
		matched := buttonMatcher.FindStringSubmatch(inputs[i])
		x, err := strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse Button A x value from input %q", inputs[i])
		}
		y, err := strconv.Atoi(matched[3])
		if err != nil {
			log.Panicf("Could not parse Button A y value from input %q", inputs[i])
		}
		aVec := vector{x, y}

		matched = buttonMatcher.FindStringSubmatch(inputs[i+1])
		x, err = strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse Button B x value from input %q", inputs[i+1])
		}
		y, err = strconv.Atoi(matched[3])
		if err != nil {
			log.Panicf("Could not parse Button B y value from input %q", inputs[i+1])
		}
		bVec := vector{x, y}

		matched = prizeMatcher.FindStringSubmatch(inputs[i+2])
		x, err = strconv.Atoi(matched[1])
		if err != nil {
			log.Panicf("Could not parse Prize x value from input %q", inputs[i+2])
		}
		y, err = strconv.Atoi(matched[2])
		if err != nil {
			log.Panicf("Could not parse Prize B y value from input %q", inputs[i+2])
		}
		prizeVec := vector{x, y}

		machines = append(machines, &machine{aVec, bVec, prizeVec})
	}
	return machines
}

func MinCost(inputs []string) int {
	sum := 0
	return sum
}
