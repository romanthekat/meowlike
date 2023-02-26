package roguelike

import "github.com/romanthekat/meowlike/pkg/component"

func NewHuman(controller component.Controller) *Creature {
	human := &Creature{
		E: NewEntity(), Controller: controller,
		Coor: &component.Coor{X: 0, Y: 0, Altitude: 0},
		Str:  10, Int: 10,
		PhysicalEnergy: 9001, MentalEnergy: 9001,
	}

	return human
}
