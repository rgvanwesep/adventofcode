package day21

const (
	upChar    = '^'
	downChar  = 'v'
	leftChar  = '<'
	rightChar = '>'
	pressChar = 'A'
	nullChar  = 0
)

type vector struct {
	x, y int
}

type instruction struct {
	input byte
}

type state struct {
	position  vector
	isPressed bool
}

type robot interface {
	update(inst instruction) state
	get() state
	valid() []state
}

type directionalRobot struct {
	state state
}

func (r *directionalRobot) update(inst instruction) state {
	return state{}
}

func (r *directionalRobot) get() state {
	return r.state
}

func (r *directionalRobot) valid() []state {
	return []state{}
}

var _ robot = &directionalRobot{}

type numericRobot struct {
	state state
}

func (r *numericRobot) update(inst instruction) state {
	return state{}
}

func (r *numericRobot) get() state {
	return r.state
}

func (r *numericRobot) valid() []state {
	return []state{}
}

var _ robot = &numericRobot{}

type connection struct {
	from robot
	to   robot
}

func (c *connection) transmit() [2]state {
	return [2]state{c.from.get(), c.to.get()}
}

type systemState struct {
	finger, robotA, robotB, robotC state
}

type system struct {
	state             systemState
	connectionAB      connection
	connectionBC      connection
	directionalKeypad map[vector]byte
	numericKeypad     map[vector]byte
}

func (s *system) valid() []systemState {
	var (
		validFingers, validRobotAs, validRobotBs, validRobotCs []state
	)
	validStates := []systemState{}
	if s.state.finger.isPressed {
		validFingers = []state{
			{
				position:  s.state.finger.position,
				isPressed: false,
			},
		}
	} else {
		validFingers = []state{
			{
				position:  vector{1, 0},
				isPressed: true,
			},
			{
				position:  vector{1, 0},
				isPressed: true,
			},
			{
				position:  vector{0, 1},
				isPressed: true,
			},
			{
				position:  vector{1, 1},
				isPressed: true,
			},
			{
				position:  vector{2, 1},
				isPressed: true,
			},
		}
	}
	return []systemState{}
}

func (s *system) pressButton(value byte) systemState {
	return systemState{}
}

func CalcComplexity(inputs []string) int {
	return 0
}
