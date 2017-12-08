package main

import "math"

/***************************

Interface for chess pieces

***************************/
type Piece struct {
	name      string
	shortName string
	alive     bool
	team      int
}

type action interface {
	create(pos coord, team int)
	move(cPos coord, nPos coord, board *Board) bool
	specialMove(cPos coord, nPos coord, board *Board) bool
}

/***************************

Pawn implementation

***************************/

type Pawn struct {
	Piece
	moved bool
}

func (p Pawn) create(pos coord, team int) {
	p.name = "Pawn"
	p.shortName = "P"
	p.alive = true
	p.team = team
	p.moved = false
}

func (p Pawn) move(cPos coord, nPos coord, board *Board) bool {
	x_offset := float(nPos.x - cPos.x)
	y_offset := float(nPos.y - cPos.y)
	nPos_Piece := pieceAtCell(board, nPos)

	// Check that the offsets are valid
	if math.Abs(y_offset) > 1 {
		return false
	}

	if p.team == 1 && x_offset != -1 {
		return false
	}

	if p.team == 0 && x_offset != 1 {
		return false
	}

	// If the move is forward check that the forward square is empty
	if y_offset == 0 && nPos_Piece == nil {
		return true
	}

	// If the move is attacking check that the attacked squares non-empty
	if y_offset != 0 && nPos_Piece != nil {
		return true
	}

	// If the move is plausible return true
	return false
}

func (p Pawn) specialMove(cPos coord, nPos coord, board *Board) bool {

	// If the pawn has been moved already return false
	if p.moved {
		return false
	}

	x_offset := float(nPos.x - cPos.x)
	y_offset := float(nPos.y - cPos.y)

	// Check that the offsets are valid
	if math.Abs(y_offset) != 0 {
		return false
	}

	if p.team == 1 && -1*x_offset == 2 {

		// Check to see if both positions in front are empty
		ternPos := coord{x: cPos.x - 1, y: cPos.y}
		tPos_Piece := pieceAtCell(board, ternPos)
		nPos_Piece := pieceAtCell(board, nPos)

		if nPos_Piece == nil && tPos_Piece == nil {
			return true
		}
	}

	if p.team == 0 && x_offset == 2 {

		// Check to see if both positions in front are empty
		ternPos := coord{x: cPos.x + 1, y: cPos.y}
		tPos_Piece := pieceAtCell(board, ternPos)
		nPos_Piece := pieceAtCell(board, nPos)

		if nPos_Piece == nil && tPos_Piece == nil {
			return true
		}
	}

	return false
}

/***************************

Rook implementation

***************************/

type Rook struct {
	Piece
	moved bool
}

func (r Rook) create(pos coord, team int) {
	r.name = "Rook"
	r.shortName = "R"
	r.alive = true
	r.team = team
	r.moved = false
}

func (r Rook) move(cPos coord, nPos coord, board *Board) bool {
	return true
}

func (r Rook) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

Bishop implementation

***************************/

type Bishop struct {
	Piece
}

func (b Bishop) create(pos coord, team int) {
	b.name = "Bishop"
	b.shortName = "B"
	b.alive = true
	b.team = team
}

func (b Bishop) move(cPos coord, nPos coord, board *Board) bool {
	return true
}

func (b Bishop) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

Knight implementation

***************************/

type Knight struct {
	Piece
}

func (n Knight) create(pos coord, team int) {
	n.name = "Knight"
	n.shortName = "N"
	n.alive = true
	n.team = team
}

func (n Knight) move(cPos coord, nPos coord, board *Board) bool {
	return true
}

func (n Knight) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

Queen implementation

***************************/

type Queen struct {
	Piece
}

func (q Queen) create(pos coord, team int) {
	q.name = "Queen"
	q.shortName = "Q"
	q.alive = true
	q.team = team
}

func (q Queen) move(cPos coord, nPos coord, board *Board) bool {
	return true
}

func (q Queen) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

King implementation

***************************/

type King struct {
	Piece
	moved bool
}

func (k King) create(pos coord, team int) {
	k.name = "King"
	k.shortName = "k"
	k.alive = true
	k.team = team
	k.moved = false
}

func (k King) move(cPos coord, nPos coord, board *Board) bool {
	return true
}

func (k King) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}
