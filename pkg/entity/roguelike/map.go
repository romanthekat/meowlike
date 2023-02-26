package roguelike

type Cell struct {
	Creature *Creature
}

type Map struct {
	Map [][]*Cell
}

func NewTestMap() *Map {
	var result [][]*Cell

	for row := 0; row < 10; row++ {
		var row []*Cell
		for col := 0; col < 10; col++ {
			row = append(row, &Cell{})
		}
		result = append(result, row)
	}

	return &Map{Map: result}
}
