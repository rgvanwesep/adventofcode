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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

type vector struct {
	x, y int
}

type machine struct {
	aVec, bVec, prizeVec vector
}

func parseMachines(inputs []string) []machine {
	buttonMatcher := regexp.MustCompile(buttonPattern)
	prizeMatcher := regexp.MustCompile(prizePattern)
	machines := []machine{}
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

		machines = append(machines, machine{aVec, bVec, prizeVec})
	}
	return machines
}

func MinCost(inputs []string) int {
	sum := 0
	machines := parseMachines(inputs)
	for _, m := range machines {
		lcmA := lcm(m.aVec.x, m.aVec.y)
		xFact, yFact := lcmA/m.aVec.x, lcmA/m.aVec.y
		xFactPrize, xFactB := xFact*m.prizeVec.x, xFact*m.bVec.x
		numerB := xFactPrize - yFact*m.prizeVec.y
		denom := xFactB - yFact*m.bVec.y
		bRem := numerB % denom
		if bRem == 0 {
			b := numerB / denom
			numerA := xFactPrize - xFactB*b
			aRem := numerA % lcmA
			if aRem == 0 {
				a := numerA / lcmA
				sum += a*aCost + b*bCost
			}
		}
	}
	return sum
}
