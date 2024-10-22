package models

// Cube represents a cube with a given side length
type Cube struct {
    sideLength int
}

// NewCube creates a new Cube instance
func NewCube(sideLength int) *Cube {
    return &Cube{sideLength: sideLength}
}

// Volume calculates the volume of the cube
func (c *Cube) Volume() int {
    return c.sideLength * c.sideLength * c.sideLength
}
