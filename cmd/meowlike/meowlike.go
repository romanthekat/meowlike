package main

import (
	"fmt"
	physics2 "github.com/romanthekat/meowlike/pkg/system/physics"
)

func main() {
	//smallBowEnergy := physics.Energy(35)
	//mediumBowEnergy := physics.Energy(55)
	//largeBowEnergy := physics.Energy(75)
	//greatBowEnergy := physics.Energy(100)

	crossbowEnergy := physics2.Energy(120)

	boltMass := physics2.Mass(0.025)
	//arrowMass := physics.Mass(0.016)

	//shortSwordMass := physics.Mass(1.3)

	var boltSpeed = physics2.GetMovementSpeed(crossbowEnergy, boltMass)
	fmt.Println(boltSpeed)
}
