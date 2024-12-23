package day22

import (
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
