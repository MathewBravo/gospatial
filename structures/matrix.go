package structures

type Steric interface {
	// Return the topological dimensions of the geometry
	// 0 = point
	// 1 = line
	// 2 = polygon
	Dimensions() int

	// Returns true if the geometry fits the OGC standards
	IsValid() bool

	// Returns true if the geometry is non-empty
	IsEmpty() bool

	// Checks for self intersections within the geometry
	IsSimple() bool
}

type Matrix []float64
