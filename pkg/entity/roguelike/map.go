package roguelike

type Cell struct {
	Creature *Creature
}

type Map struct {
	Map [][]*Cell
}
