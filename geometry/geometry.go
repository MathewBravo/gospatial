package geometry

const (
	Point      = "Point"
	LineString = "LineString"
	Polygon    = "Polygon"

	MultiPoint      = "MultiPoint"
	MultiLineString = "MultiLineString"
	MultiPolygon    = "MultiPolygon"

	CircularString               = "CircularString"
	CompountCurve                = "CompoundCurve"
	CurvePolygon                 = "CurvePolygon"
	PolyhedralSurface            = "PolyhedralSurface"
	TriangulatedIrregularNetwork = "TriangulatedIrregularNetwork"
	Triangle                     = "Triangle"
)

type Geometry interface {
	// Returns true if the given geometries are "topologically equal".
	// Topological equality means that the geometries have the same dimension,
	// and their point-sets occupy the same space.
	STEquals(g *Geometry) (bool, error)

	// Returns true if two geometries are disjoint.
	// Geometries are disjoint if they have no point in common.
	STDisjoint(g *Geometry) (bool, error)

	// Return true if two geometries intersect.
	// Geometries intersect if they have any point in common.
	STIntersects(g *Geometry) (bool, error)

	// Returns TRUE if A and B intersect, but their interiors
	// do not intersect. Equivalently, A and B have at least
	// one point in common, and the common points lie in at least one boundary.
	STTouches(g *Geometry) (bool, error)

	// Compares two geometry objects and returns true if their intersection
	// "spatially crosses"; that is, the geometries have some, but not all
	// interior points in common.
	STCrosses(g *Geometry) (bool, error)

	// Returns TRUE if geometry A is within geometry B. A is within B if
	// and only if all points of A lie inside (i.e. in the interior or boundary of)
	// B (or equivalently, no points of A lie in the exterior of B),
	// and the interiors of A and B have at least one point in common.
	STWithin(g *Geometry) (bool, error)

	// Returns TRUE if geometry A contains geometry B. A contains B if and
	// only if all points of B lie inside (i.e. in the interior or boundary of)
	// A (or equivalently, no points of B lie in the exterior of A), and the
	// interiors of A and B have at least one point in common.
	STContains(g *Geometry) (bool, error)

	// Returns TRUE if geometry A and B "spatially overlap". Two geometries
	// overlap if they have the same dimension, their interiors intersect in
	// that dimension. and each has at least one point inside the other.
	STOverlaps(g *Geometry) (bool, error)

	// For geometry types returns the minimum 2D Cartesian (planar) distance
	// between two geometries, in projected units (spatial ref units).
	STDistance(g *Geometry) (float64, error)

	// Returns the area of a polygonal geometry.
	STArea() (float64, error)

	// For geometry types: returns the 2D Cartesian length of the geometry if it is a
	// LineString, MultiLineString, ST_Curve, ST_MultiCurve.
	STLength() (float64, error)

	//Returns the 2D perimeter of the geometry/geography if it is a Polygon, MultiPolygon
	STPerimeter() (float64, error)

	// Computes a POLYGON or MULTIPOLYGON that represents all points whose distance from a
	// geometry/geography is less than or equal to a given distance.
	STBuffer(radius_to_buffer float64, quad_segs int)

	// Computes the convex hull of a geometry. The convex hull is the smallest convex
	// geometry that encloses all geometries in the input.
	STConvexHull() (Geometry, error)

	// Returns a geometry representing the point-set intersection of two geometries.
	// In other words, that portion of geometry A and geometry B that is shared between the two geometries.
	STIntersection(g *Geometry) (Geometry, error)

	// Unions the input geometries, merging geometry to produce a result geometry with no overlaps.
	// The output may be an atomic geometry, a MultiGeometry, or a Geometry Collection.
	STUnion(g *Geometry) (Geometry, error)

	// Returns a geometry representing the part of geometry A that does not intersect geometry B.
	STDifference(g *Geometry) (Geometry, error)

	// Computes a point which is the geometric center of mass of a geometry.
	STCentroid(g *Geometry) (Geometry, error)

	// Returns a POINT which is guaranteed to lie in the interior of a surface.
	STPointOnSurface() (Geometry, error)

	// Returns a LINESTRING representing the exterior ring (shell) of a POLYGON.
	// Returns an Error if the geometry is not a polygon.
	STExteriorRing() (Geometry, error)

	// Returns the Nth interior ring (hole) of a POLYGON geometry as a LINESTRING.
	// The index starts at 1. Returns an Error if the geometry is not a polygon or
	// the index is out of range.
	STInteriorRingN(interior_ring int) (Geometry, error)

	// Return the number of interior rings of a polygon geometry.
	// Returns an Error if the geometry is not a polygon.
	STNumInteriorRings() (int, error)

	// Return the 1-based Nth element geometry of an input geometry which is a GEOMETRYCOLLECTION,
	// MULTIPOINT, MULTILINESTRING, MULTICURVE, MULTIPOLYGON, or POLYHEDRALSURFACE.
	// Otherwise, returns error.
	STGeometryN(n_element_geo int) (Geometry, error)

	// Returns the number of elements in a geometry collection (GEOMETRYCOLLECTION or MULTI*).
	// For non-empty atomic geometries returns 1. For empty geometries returns 0.
	STNumGeometries() (int, error)

	// Computes a simplified representation of a geometry using the Douglas-Peucker algorithm.
	// The simplification tolerance is a distance value, in the units of the input SRS.
	// Simplification removes vertices which are within the tolerance distance of the simplified linework.
	// The result may not be valid even if the input is.
	STSimplify(tolerence int) (Geometry, error)

	// Future Support
	// STRelate
	// STSnapToGrid
	// STTranslate
	// STScale
	// STRotate
	// STAffine
}
