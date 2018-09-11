package ogame

import "fmt"

// Coordinate ...
type Coordinate struct {
	Galaxy   int
	System   int
	Position int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("[%d:%d:%d]", c.Galaxy, c.System, c.Position)
}

// Equal ...
func (c Coordinate) Equal(v Coordinate) bool {
	return c.Galaxy == v.Galaxy &&
		c.System == v.System &&
		c.Position == v.Position
}