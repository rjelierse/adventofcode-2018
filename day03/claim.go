package day03

type Claim struct {
	id int
	top int
	left int
	width int
	height int
}

func (c Claim) Occupies() []Coordinate {
	var coordinates []Coordinate
	for x := c.left; x < c.left + c.width; x++ {
		for y := c.top; y < c.top + c.height; y++ {
			coordinates = append(coordinates, Coordinate{x, y})
		}
	}
	return coordinates
}
