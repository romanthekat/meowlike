package roguelike

import (
	"github.com/google/uuid"
	"github.com/romanthekat/meowlike/pkg/component"
)

type Entity struct {
	Id string

	Rules []*Rule
}

func NewEntity() *Entity {
	id := uuid.Must(uuid.NewRandom()).String()
	return &Entity{Id: id}
}

type Creature struct {
	E          *Entity
	Controller component.Controller

	Coor    *component.Coor
	Kinetic *component.Kinetic
	//TODO body system
	//Physical component.Physical

	Str float64
	Int float64

	//skills
	//inventory

	Items []*Item

	PhysicalEnergy float64
	MentalEnergy   float64
}

type Item struct {
	E Entity
	//usage type

	Coor     component.Coor
	Physical component.Physical
	Kinetic  component.Kinetic
}
