package main

/***************************

coord tuple custom type interface

***************************/

type coord struct {
	x int
	y int
}

func createCoord(x int, y int) *coord {
	coord := new(coord)
	coord.x = x
	coord.y = y
	return coord
}
