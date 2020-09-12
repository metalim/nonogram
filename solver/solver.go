package solver

import (
	"strconv"
	"strings"
)

// Cell content
type Cell = rune

// Cell types
const (
	Empty   Cell = '⬜'
	Set          = '⬛'
	Crossed      = '❌'
)

// State of nonogram
type State struct {
	h, w   int
	xh, yh [][]int
	data   [][]Cell
}

// NewState from string
func NewState(puzzle string) *State {
	s := new(State)
	ss := strings.Split(puzzle, "\n")
	dims := ints(ss[0])
	s.w = dims[0]
	s.h = dims[1]
	for x := 0; x < s.w; x++ {
		s.xh = append(s.xh, ints(ss[1+x]))
	}
	for y := 0; y < s.h; y++ {
		s.yh = append(s.yh, ints(ss[1+s.w+y]))
	}

	return s
}

func ints(s string) []int {
	var out []int
	for _, f := range strings.Fields(s) {
		n, _ := strconv.Atoi(f)
		out = append(out, n)
	}
	return out
}

// Solve nonogram
func (s *State) Solve() bool {
	return true
}

// ToString converts nonogram to text
func (s *State) ToString() string {
	var out strings.Builder
	for _, row := range s.data {
		out.WriteString(string(row))
		out.WriteRune('\n')
	}
	return out.String()
}
