package roguelike

import (
	"github.com/romanthekat/meowlike/pkg/component"
	"github.com/romanthekat/meowlike/pkg/system/physics"
)

type Entity struct {
	Id string

	Rules []*Rule
}

type Creature struct {
	E          Entity
	Controller component.Controller

	Coor     component.Coor
	Physical component.Physical
	Kinetic  component.Kinetic

	Str float64
	Int float64

	//skills
	//inventory

	Items []*Item

	MentalEnergy physics.Energy
}

type Item struct {
	E Entity
	//usage type

	Coor     component.Coor
	Physical component.Physical
	Kinetic  component.Kinetic
}
