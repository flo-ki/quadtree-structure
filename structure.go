package main

import "fmt"

type Point struct {
	X, Y int
}

type Node struct {
	Point    *Point
	NE, NW, SE, SW *Node
}

type QuadTree struct {
	Root *Node
}

// Insert a new point into the QuadTree
func (qt *QuadTree) Insert(p *Point) {
	// Insertion logic here
}

// Search for a point in the QuadTree
func (qt *QuadTree) Search(p *Point) *Node {
	// Search logic here
	return nil
}

// Retrieve the nearest point to a given point in the QuadTree
func (qt *QuadTree) Nearest(p *Point) *Point {
	// Nearest logic here
	return nil
}

func main() {
	qt := &QuadTree{}
	
	point := &Point{X: 5, Y: 5}
	qt.Insert(point)
	
	fmt.Println(qt.Search(point))
	fmt.Println(qt.Nearest(point))
}
