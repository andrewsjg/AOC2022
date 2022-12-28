package Puzzle12

type Location struct {
	x, y     int
	topoType int

	Map topoMap
}

type topoMap map[int]map[int]*Location

func (t topoMap) setLocation(l *Location, x, y int) {
	if t[x] == nil {
		t[x] = map[int]*Location{}
	}

	l.x = x
	l.y = y

	t[x][y] = l
	l.Map = t
}

func (l Location) pathToNeighBourCost(to Pather) float64 {
	toLocation := to.(*Location)
	return float64(toLocation.topoType)
}

func (l Location) pathCostEstimate(to Pather) float64 {

	toLocation := to.(*Location)
	absX := toLocation.x - l.x
	if absX < 0 {
		absX = -absX
	}
	absY := toLocation.y - l.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (l Location) neighbours() []Pather {
	neighbours := []Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {

		if neighbour := l.Map.getLocation(l.x+offset[0], l.y+offset[1]); neighbour != nil {

			diff := neighbour.topoType - l.topoType
			//fmt.Printf("Location TopoTypeVal: %d,Neighbour TopoTypeVal: %d diff: %d\n", l.topoType, neighbour.topoType, diff)

			// Exclude neighbours that are to high
			if diff < 2 {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
}

func (t topoMap) getLocation(x, y int) *Location {
	if t[x] == nil {
		return nil
	}
	return t[x][y]
}
