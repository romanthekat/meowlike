package system

import (
	"github.com/romanthekat/meowlike/pkg/core"
	"github.com/romanthekat/meowlike/pkg/entity/roguelike"
)

const TickLength = 100

type Game struct {
	Maps []*roguelike.Map

	GlobalRules []*roguelike.Rule
	Rules       []*roguelike.Rule

	timeLoop *core.Loop
	exit     bool
}

func NewGame() *Game {
	game := &Game{
		Maps:        nil,
		GlobalRules: nil,
		Rules:       nil,
		timeLoop:    core.NewLoop(),
		exit:        false,
	}

	return game
}

func (g *Game) MainLoop() {
	for !g.exit {
		//handle timeLoop
		for e := g.timeLoop.Next(); e != nil; e = g.timeLoop.Next() {
			sentient := e.S

			sentient.AddEnergy(TickLength)

			if sentient.GetEnergy() > 0 {
				cost := sentient.Act()
				sentient.SubEnergy(cost)
			}
		}

		//handle systems
		//TODO iterate over systems x over entities, apply systems
	}
}
