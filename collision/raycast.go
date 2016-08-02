package collision

import (
	"bitbucket.org/oakmoundstudio/plasticpiston/plastic/event"
	"math"
)

// RayCast returns the set of points where a line
// from x,y going at a certain angle, for a certain length, intersects
// with existing rectangles in the rtree.
// It converts the ray into a series of points which are themselves
// used to check collision at a miniscule width and height.
func RayCast(x, y, degrees, length float64) []CollisionPoint {
	results := []CollisionPoint{}
	resultHash := make(map[*Space]bool)

	s := math.Sin(degrees * math.Pi / 180)
	c := math.Cos(degrees * math.Pi / 180)
	for i := 0.0; i < length; i++ {
		loc := NewRect(x, y, .1, .1)

		next := rt.SearchIntersect(loc)

		for k := 0; k < len(next); k++ {
			nx := (next[k].(*Space))
			if _, ok := resultHash[nx]; !ok {
				resultHash[nx] = true
				results = append(results, CollisionPoint{nx, x, y})
			}
		}
		x += c
		y += s
	}
	return results
}

// RatCastSingle acts as RayCast, but it returns only the first collision
// that the generated ray intersects, ignoring entities
// in the given invalidIDs list.
// Example Use case: shooting a bullet, hitting the first thing that isn't yourself.
func RayCastSingle(x, y, degrees, length float64, invalidIDS []event.CID) CollisionPoint {

	s := math.Sin(degrees * math.Pi / 180)
	c := math.Cos(degrees * math.Pi / 180)
	for i := 0.0; i < length; i++ {
		loc := NewRect(x, y, .1, .1)
		next := rt.SearchIntersect(loc)
	output:
		for k := 0; k < len(next); k++ {
			nx := (next[k].(*Space))
			for e := 0; e < len(invalidIDS); e++ {
				if nx.CID == invalidIDS[e] {
					continue output
				}
			}
			return CollisionPoint{nx, x, y}
		}
		x += c
		y += s

	}
	return CollisionPoint{}
}

func RayCastSingleLabel(x, y, degrees, length float64, label int) CollisionPoint {

	s := math.Sin(degrees * math.Pi / 180)
	c := math.Cos(degrees * math.Pi / 180)
	for i := 0.0; i < length; i++ {
		loc := NewRect(x, y, .1, .1)
		next := rt.SearchIntersect(loc)
		for k := 0; k < len(next); k++ {
			nx := (next[k].(*Space))
			if nx.Label == label {
				return CollisionPoint{nx, x, y}
			}
		}
		x += c
		y += s

	}
	return CollisionPoint{}
}