package system

import "github.com/romanthekat/meowlike/pkg/entity/roguelike"

type Game struct {
	Maps []*roguelike.Map

	//coefficients

	GlobalRules []*roguelike.Rule
	Rules       []*roguelike.Rule
}

func MainLoop() {
	done := false
	for !done {

	}
}
