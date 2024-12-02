package day9

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Function []int

func NewFunction(line string) Function {
	var f Function
	line = strings.TrimRight(line, "\n")
	split := strings.Split(line, " ")
	for _, s := range split {
		n, _ := strconv.Atoi(s)
		f = append(f, n)
	}
	return f
}

func (f Function) Diff() Function {
	var diff Function
	for i := 1; i < len(f); i++ {
		diff = append(diff, f[i]-f[i-1])
	}
	return diff
}

func (f Function) IsZero() bool {
	for _, v := range f {
		if v != 0 {
			return false
		}
	}
	return true
}

type Differences []Function

func (d Differences) Diff() Differences {
	diff := d[len(d)-1].Diff()
	return append(d, diff)
}

func (d Differences) NextValues() ([]int, error) {
	if !d[len(d)-1].Diff().IsZero() {
		return nil, errors.New("last difference is not zero")
	}
	nextValues := make([]int, len(d))
	last := 0
	for i := len(d) - 1; i >= 0; i-- {
		diff := d[i]
		nextValues[i] = diff[len(diff)-1] + last
		last = nextValues[i]
	}
	return nextValues, nil
}

func (d Differences) PrevValues() ([]int, error) {
	if !d[len(d)-1].Diff().IsZero() {
		return nil, errors.New("last difference is not zero")
	}
	prevValues := make([]int, len(d))
	last := 0
	for i := len(d) - 1; i >= 0; i-- {
		diff := d[i]
		prevValues[i] = diff[0] - last
		last = prevValues[i]
	}
	return prevValues, nil
}

type Histories []Differences

func NewHistories(lines []string) Histories {
	var histories Histories
	for _, line := range lines {
		f := NewFunction(line)
		d := Differences{f}
		histories = append(histories, d)
	}
	return histories
}

func (h Histories) Diff() Histories {
	var diffed Histories
	for _, ds := range h {
		if ds[0].IsZero() {
			continue
		}
		for d := ds[0].Diff(); !d.IsZero(); d = d.Diff() {
			ds = append(ds, d)
		}
		diffed = append(diffed, ds)
	}
	return diffed
}

func (h Histories) SumNextValues() int {
	sum := 0
	for _, ds := range h {
		nextValues, err := ds.NextValues()
		if err != nil {
			log.Fatal(err)
		}
		sum += nextValues[0]
	}
	return sum
}

func (h Histories) SumPrevValues() int {
	sum := 0
	for _, ds := range h {
		prevValues, err := ds.PrevValues()
		if err != nil {
			log.Fatal(err)
		}
		sum += prevValues[0]
	}
	return sum
}

func Sum(lines []string) int {
	histories := NewHistories(lines)
	diffed := histories.Diff()
	return diffed.SumNextValues()
}

func SumPrev(lines []string) int {
	histories := NewHistories(lines)
	diffed := histories.Diff()
	return diffed.SumPrevValues()
}
