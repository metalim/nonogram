package solver

import (
	"fmt"
	"testing"
)

func expect(t *testing.T, ex, got string) {
	if got != ex {
		t.Fatalf("Expected: %s\nGot: %s", ex, got)
	}
}

func TestFormat(t *testing.T) {
	var s State
	s.data = [][]Cell{
		[]Cell("❌⬛⬜"),
		[]Cell("⬜❌⬛"),
		[]Cell("⬛⬜❌"),
	}
	expect(t, `❌⬛⬜
⬜❌⬛
⬛⬜❌
`,
		s.ToString())
}

func TestNew(t *testing.T) {
	s := NewState(`3 3
		1 2 3
		2 3 1
		3 1 2
		1 2 3
		2 3 1
		3 1 2
		`)
	s.Solve()
	fmt.Print(s)
}
