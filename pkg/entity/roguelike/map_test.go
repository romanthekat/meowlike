package roguelike

import (
	"testing"
)

func TestNewTestMap(t *testing.T) {
	m := NewTestMap()

	if len(m.Map) != 10 || len(m.Map[0]) != 10 {
		t.Fatalf("generated test map is incorrect")
	}
}
