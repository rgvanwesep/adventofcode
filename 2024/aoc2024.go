package main

import (
	"aoc2024/day1"
	"aoc2024/day10"
	"aoc2024/day11"
	"aoc2024/day12"
	"aoc2024/day13"
	"aoc2024/day14"
	"aoc2024/day15"
	"aoc2024/day16"
	"aoc2024/day17"
	"aoc2024/day18"
	"aoc2024/day19"
	"aoc2024/day2"
	"aoc2024/day20"
	"aoc2024/day21"
	"aoc2024/day22"
	"aoc2024/day23"
	"aoc2024/day24"
	"aoc2024/day25"
	"aoc2024/day3"
	"aoc2024/day4"
	"aoc2024/day5"
	"aoc2024/day6"
	"aoc2024/day7"
	"aoc2024/day8"
	"aoc2024/day9"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)

func main() {
	dayFlag := flag.Int("d", 0, "Day to run")
	partFlag := flag.Int("p", 1, "Part to run")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile := flag.String("memprofile", "", "write memory profile to file")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("Could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	day, part := *dayFlag, *partFlag
	log.Printf("Running day %d, part %d", day, part)

	reader := bufio.NewReader(os.Stdin)
	inputLines := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		inputLines = append(inputLines, strings.TrimRight(line, "\n"))
		if err != nil {
			break
		}
	}
	log.Printf("Read %d lines\n", len(inputLines))

	writer := bufio.NewWriter(os.Stdout)

	switch [2]int{day, part} {
	case [2]int{1, 1}:
		writer.WriteString(fmt.Sprintln(day1.SumDistances(inputLines)))
	case [2]int{1, 2}:
		writer.WriteString(fmt.Sprintln(day1.CalcSimilarity(inputLines)))
	case [2]int{2, 1}:
		writer.WriteString(fmt.Sprintln(day2.CountSafeReports(inputLines)))
	case [2]int{2, 2}:
		writer.WriteString(fmt.Sprintln(day2.CountSafeReportsDamped(inputLines)))
	case [2]int{3, 1}:
		writer.WriteString(fmt.Sprintln(day3.SumMul(inputLines)))
	case [2]int{3, 2}:
		writer.WriteString(fmt.Sprintln(day3.SumConditionalMul(inputLines)))
	case [2]int{4, 1}:
		writer.WriteString(fmt.Sprintln(day4.CountOccurances(inputLines)))
	case [2]int{4, 2}:
		writer.WriteString(fmt.Sprintln(day4.CountOccurancesX(inputLines)))
	case [2]int{5, 1}:
		writer.WriteString(fmt.Sprintln(day5.SumMiddlePages(inputLines)))
	case [2]int{5, 2}:
		writer.WriteString(fmt.Sprintln(day5.SumCorrectedMiddlePages(inputLines)))
	case [2]int{6, 1}:
		writer.WriteString(fmt.Sprintln(day6.CountVisited(inputLines)))
	case [2]int{6, 2}:
		writer.WriteString(fmt.Sprintln(day6.CountCyclingObstructions(inputLines)))
	case [2]int{7, 1}:
		writer.WriteString(fmt.Sprintln(day7.SumCorrected(inputLines)))
	case [2]int{7, 2}:
		writer.WriteString(fmt.Sprintln(day7.SumCorrectedWithConcat(inputLines)))
	case [2]int{8, 1}:
		writer.WriteString(fmt.Sprintln(day8.CountAntiNodes(inputLines)))
	case [2]int{8, 2}:
		writer.WriteString(fmt.Sprintln(day8.CountAntiNodesHarmonics(inputLines)))
	case [2]int{9, 1}:
		writer.WriteString(fmt.Sprintln(day9.CalcChecksum(inputLines)))
	case [2]int{9, 2}:
		writer.WriteString(fmt.Sprintln(day9.CalcChecksumFileSwap(inputLines)))
	case [2]int{10, 1}:
		writer.WriteString(fmt.Sprintln(day10.SumTrailScores(inputLines)))
	case [2]int{10, 2}:
		writer.WriteString(fmt.Sprintln(day10.SumTrailRatings(inputLines)))
	case [2]int{11, 1}:
		writer.WriteString(fmt.Sprintln(day11.CountPebbles(inputLines, 25)))
	case [2]int{11, 2}:
		writer.WriteString(fmt.Sprintln(day11.CountPebbles(inputLines, 75)))
	case [2]int{12, 1}:
		writer.WriteString(fmt.Sprintln(day12.SumFencePrice(inputLines)))
	case [2]int{12, 2}:
		writer.WriteString(fmt.Sprintln(day12.SumFencePriceDiscount(inputLines)))
	case [2]int{13, 1}:
		writer.WriteString(fmt.Sprintln(day13.MinCost(inputLines)))
	case [2]int{13, 2}:
		writer.WriteString(fmt.Sprintln(day13.MinCostBig(inputLines)))
	case [2]int{14, 1}:
		writer.WriteString(fmt.Sprintln(day14.CalcSafetyFactor(inputLines, 103, 101, 100)))
	case [2]int{14, 2}:
		writer.WriteString(fmt.Sprintln(day14.FindSignal(inputLines, 103, 101)))
	case [2]int{15, 1}:
		writer.WriteString(fmt.Sprintln(day15.SumCoordinates(inputLines)))
	case [2]int{15, 2}:
		writer.WriteString(fmt.Sprintln(day15.SumCoordinatesWide(inputLines)))
	case [2]int{16, 1}:
		writer.WriteString(fmt.Sprintln(day16.MinScore(inputLines)))
	case [2]int{16, 2}:
		writer.WriteString(fmt.Sprintln(day16.CountTiles(inputLines)))
	case [2]int{17, 1}:
		writer.WriteString(fmt.Sprintln(day17.ExecProgram(inputLines)))
	case [2]int{17, 2}:
		writer.WriteString(fmt.Sprintln(day17.FindRegisterAValue(inputLines)))
	case [2]int{18, 1}:
		writer.WriteString(fmt.Sprintln(day18.CountSteps(inputLines, 71, 71, 1024)))
	case [2]int{18, 2}:
		writer.WriteString(fmt.Sprintln(day18.FindFinalInput(inputLines, 71, 71)))
	case [2]int{19, 1}:
		writer.WriteString(fmt.Sprintln(day19.CountPossible(inputLines)))
	case [2]int{19, 2}:
		writer.WriteString(fmt.Sprintln(day19.SumCombinations(inputLines)))
	case [2]int{20, 1}:
		writer.WriteString(fmt.Sprintln(day20.CountCheats(inputLines, 2, 100)))
	case [2]int{20, 2}:
		writer.WriteString(fmt.Sprintln(day20.CountCheats(inputLines, 20, 100)))
	case [2]int{21, 1}:
		writer.WriteString(fmt.Sprintln(day21.CalcComplexity(inputLines, 3)))
	case [2]int{21, 2}:
		writer.WriteString(fmt.Sprintln(day21.CalcComplexity(inputLines, 26)))
	case [2]int{22, 1}:
		writer.WriteString(fmt.Sprintln(day22.SumSecrets(inputLines)))
	case [2]int{22, 2}:
		writer.WriteString(fmt.Sprintln(day22.SumSellPrices(inputLines)))
	case [2]int{23, 1}:
		writer.WriteString(fmt.Sprintln(day23.CountLANs(inputLines)))
	case [2]int{23, 2}:
		writer.WriteString(fmt.Sprintln(day23.FindPassword(inputLines)))
	case [2]int{24, 1}:
		writer.WriteString(fmt.Sprintln(day24.Evaluate(inputLines)))
	case [2]int{24, 2}:
		writer.WriteString(fmt.Sprintln(day24.FindSwapped(inputLines)))
	case [2]int{25, 1}:
		writer.WriteString(fmt.Sprintln(day25.CountFits(inputLines)))
	default:
		log.Fatal("Invalid day or part")
	}

	writer.Flush()

	log.Print("Done")

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("Could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("Could not write memory profile: ", err)
		}
	}
}
