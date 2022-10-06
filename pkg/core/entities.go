package core

import "github.com/romanthekat/meowlike/pkg/physics"

type Game struct {
	Map [][]*Cell

	//coefficients

	GlobalRules []Rule
	Rules       []Rule
}

type Entity struct {
	Id  string
	Obj PhysicalObject

	Rules []*Rule
}

type Creature struct {
	E          Entity
	Controller Controller

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
}

type Rule struct {
	IsGlobal bool
	//Condition
	//Action
}

type Cell struct {
	Creature *Creature
}
