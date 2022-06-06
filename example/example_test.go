package example

import "testing"

func TestRaceCondition(t *testing.T) {
	example(3)
}
