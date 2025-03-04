package day22

import (
	"iter"
	"log"
	"strconv"
)

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	const modulus = 16777216
	return a % modulus
}

func stepA(current int) int {
	const multiplier = 64
	next := current
	next *= multiplier
	next = mix(next, current)
	next = prune(next)
	return next
}

func stepB(current int) int {
	const divisor = 32
	next := current
	next /= divisor
	next = mix(next, current)
	next = prune(next)
	return next
}

func stepC(current int) int {
	const multiplier = 2048
	next := current
	next *= multiplier
	next = mix(next, current)
	next = prune(next)
	return next
}

func calcNextSecret(current int) int {
	operations := [3]func(int) int{stepA, stepB, stepC}
	for _, op := range operations {
		current = op(current)
	}
	return current
}

func calcFinalSecret(seed, nSecrets int) int {
	secret := seed
	for range nSecrets {
		secret = calcNextSecret(secret)
	}
	return secret
}

func SumSecrets(inputs []string) int {
	const nSecrets = 2000
	sum := 0
	for _, input := range inputs {
		if seed, err := strconv.Atoi(input); err == nil {
			sum += calcFinalSecret(seed, nSecrets)
		} else {
			log.Panicf("Cound not parse seed from input %q", input)
		}
	}
	return sum
}

func generatePrices(seed, nSecrets int) []int {
	secret := seed
	prices := make([]int, nSecrets+1)
	prices[0] = secret % 10
	for i := range nSecrets {
		secret = calcNextSecret(secret)
		prices[i+1] = secret % 10
	}
	return prices
}

func generatePriceChanges(prices []int) []int {
	size := len(prices) - 1
	priceChanges := make([]int, size)
	for i := range size {
		priceChanges[i] = prices[i+1] - prices[i]
	}
	return priceChanges
}

func generatePriceChangeSequences(priceChanges []int) []int32 {
	n := len(priceChanges) - 4
	seqs := make([]int32, n)
	for i := range n {
		seq := int32(priceChanges[i])
		for j := range 3 {
			seq <<= 8
			seq ^= int32(priceChanges[i+j+1])
		}
		seqs[i] = seq
	}
	return seqs
}

func findChangeSequenceFast(seq int32, priceChanges []int32) (int, bool) {
	n := len(priceChanges)
	j := 0
	isMatch := false
	for i := range n {
		isMatch = seq == priceChanges[i]
		if isMatch {
			j = i + 4
			break
		}
	}
	return j, isMatch
}

func findSellPriceFast(seq int32, priceChanges []int32, prices []int) int {
	sellPrice := 0
	if i, ok := findChangeSequenceFast(seq, priceChanges); ok {
		sellPrice = prices[i]
	}
	return sellPrice
}

var changeSequencesFast iter.Seq[int32] = func(yield func(int32) bool) {
	changeValues := [19]int{}
	for i, j := 0, -9; j <= 9; i, j = i+1, j+1 {
		changeValues[i] = j
	}
	for _, i := range changeValues {
		for _, j := range changeValues {
			for _, k := range changeValues {
				for _, l := range changeValues {
					seq := int32(i)
					seq <<= 8
					seq ^= int32(j)
					seq <<= 8
					seq ^= int32(k)
					seq <<= 8
					seq ^= int32(l)
					if !yield(seq) {
						break
					}
				}
			}
		}
	}
}

func SumSellPrices(inputs []string) int {
	const nSecrets = 2000
	nBuyers := len(inputs)
	prices := make([][]int, nBuyers)
	priceChanges := make([][]int, nBuyers)
	priceChangeSequences := make([][]int32, nBuyers)
	for i, input := range inputs {
		if seed, err := strconv.Atoi(input); err == nil {
			prices[i] = generatePrices(seed, nSecrets)
			priceChanges[i] = generatePriceChanges(prices[i])
			priceChangeSequences[i] = generatePriceChangeSequences(priceChanges[i])
		} else {
			log.Panicf("Cound not parse seed from input %q", input)
		}
	}
	maxSum := 0
	for changeSequence := range changeSequencesFast {
		sum := 0
		for i := range nBuyers {
			sum += findSellPriceFast(changeSequence, priceChangeSequences[i], prices[i])
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
