package day21

const (
	upChar       = '^'
	downChar     = 'v'
	leftChar     = '<'
	rightChar    = '>'
	activateChar = 'A'
)

type presser interface {
	move(direction byte)
	press()
}

type controller interface {
	signal(value byte)
}

type link struct {
	c *controller
	p *presser
}

func CalcComplexity(inputs []string) int {
	return 0
}
