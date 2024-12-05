package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func (r *Rule) Apply(u Update) bool {
	if first, ok := u.index[r.before]; ok {
		if second, ok := u.index[r.after]; ok {
			return first < second
		}
	}
	return true
}

type Update struct {
	pages  []int
	index  map[int]int
	middle int
}

func NewUpdate(pages []int) Update {
	index := make(map[int]int)
	for i, page := range pages {
		index[page] = i
	}
	middle := pages[len(pages)/2]
	return Update{pages, index, middle}
}

type UpdateInstructions struct {
	rules   []Rule
	updates []Update
}

func NewUpdateInstructions() UpdateInstructions {
	rules := make([]Rule, 0)
	updates := make([]Update, 0)
	return UpdateInstructions{rules, updates}
}

func ParseUpdateInstructions(inputs []string) UpdateInstructions {
	instructions := NewUpdateInstructions()
	inUpdateSection := false
	for _, input := range inputs {
		if input == "" {
			inUpdateSection = true
			continue
		}
		if inUpdateSection {
			pages := make([]int, 0)
			for _, s := range strings.Split(input, ",") {
				if page, err := strconv.Atoi(s); err == nil {
					pages = append(pages, page)
				} else {
					panic(fmt.Sprintf("Unable to parse page in input: %q", input))
				}
			}
			update := NewUpdate(pages)
			instructions.updates = append(instructions.updates, update)
		} else {
			rInt := [2]int{}
			for i, s := range strings.Split(input, "|") {
				if n, err := strconv.Atoi(s); err == nil {
					rInt[i] = n
				} else {
					panic(fmt.Sprintf("Unable to parse page in input: %q", input))
				}
			}
			rule := Rule{before: rInt[0], after: rInt[1]}
			instructions.rules = append(instructions.rules, rule)
		}
	}
	return instructions
}

func SumMiddlePages(inputs []string) int {
	sum := 0
	instructions := ParseUpdateInstructions(inputs)
	for _, update := range instructions.updates {
		isValid := true
		for _, rule := range instructions.rules {
			if !rule.Apply(update) {
				isValid = false
				break
			}
		}
		if isValid {
			sum += update.middle
		}
	}
	return sum
}
