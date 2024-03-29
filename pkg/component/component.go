package component

import (
	"github.com/romanthekat/meowlike/pkg/system/physics"
)

type Coor struct {
	X, Y     int
	Altitude int
}

type Kinetic struct {
	Vector physics.Vector
	//Speed  physics.Speed

	Energy          physics.Energy
	PotentialEnergy physics.Energy
}

type Form string

const (
	Solid  Form = "Solid"
	Liquid Form = "Liquid"
	Gas    Form = "Gas"
	Plasma Form = "Plasma"
)

type Material struct {
	Density  float64
	Hardness float64

	IsMetallic bool

	HeatCapacity float64
	FreezePoint  physics.Temperature
	IgnitePoint  physics.Temperature
	MeltingPoint physics.Temperature
	//GasPoint     physics.Temperature
	//PlasmaPoint  physics.Temperature

	HeatDamagePoint physics.Temperature
	ColdDamagePoint physics.Temperature

	Tags map[string]string
}

type Physical struct {
	Integrity Integrity

	Material Material
	Mass     physics.Mass

	Form Form
	Temp physics.Temperature

	ElectricCharge ElectricCharge
}

// Integrity abstracts is the object is still intact, health, etc.
type Integrity float64

type Force physics.Vector

type Field struct {
	Vector physics.Vector
}

type ElectricCharge physics.Charge

//well, there are no magnetic monopoles
//type MagneticCharge Charge

type Controller string

const (
	Player Controller = "player"
	AI                = "AI"
	None              = "none"
)
