package physics

import "math"

// GetMovementSpeed uses kinetic energy and mass to calculate movement speed of an object.
//E=m*v^2/2 -> v=(2E/m)^0.5
func GetMovementSpeed(energy Energy, mass Mass) Speed {
	return Speed(math.Sqrt(float64(energy) * 2 / float64(mass)))
}
