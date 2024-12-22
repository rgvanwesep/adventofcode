package day21

const (
	upChar    = '^'
	downChar  = 'v'
	leftChar  = '<'
	rightChar = '>'
	pressChar = 'A'
)

type vector struct {
	x, y int
}

type state struct {
	position  vector
	isPressed bool
}

func CalcComplexity(inputs []string) int {
	return 0
}
