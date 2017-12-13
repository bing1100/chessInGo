package main

/***************************

Useful custom types

***************************/

// Coordinate custom type
type coordData struct {
	x int
	y int
}

type coord interface {
	getX() int
	getY() int
}

func (c coordData) getX() int {
	return c.x
}

func (c coordData) getY() int {
	return c.y
}

// move ret type
type res struct {
	ret   bool
	board *Board
}
