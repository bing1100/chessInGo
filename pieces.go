package main

import "math"

/***************************

Interface for chess pieces

***************************/
type basePiece struct {
	name      string
	shortName string
	alive     bool
	team      int
	loc       coord
}

/*Piece
 */
type Piece interface {
	create(team int)
	move(nPos coord, board *BoardData) bool
	specialMove(nPos coord, board *BoardData) bool
	hasMoved() bool
	getTeam() int
	getName() string
	getChar() string
	setLoc(board *BoardData, nPos coord) bool
	getLoc() coord
	setDead() bool
	getAlive() bool
	copy(source Piece)
}

func (p *basePiece) specialMove(nPos coord, board *BoardData) bool {
	return false
}

func (p *basePiece) hasMoved() bool {
	return false
}

func (p *basePiece) getName() string {
	return p.name
}

func (p *basePiece) getChar() string {
	return p.shortName
}

func (p *basePiece) getTeam() int {
	return p.team
}

func (p *basePiece) getLoc() coord {
	return p.loc
}

func (p *basePiece) setLoc(board *BoardData, nPos coord) bool {
	pieceAtPos := board.pCells[nPos.getX()][nPos.getY()].pPiece

	if pieceAtPos.getChar() != p.getChar() || pieceAtPos.getTeam() != p.getTeam() {
		return false
	}

	p.loc = nPos
	return true
}

func (p *basePiece) setDead() bool {
	previous := p.alive
	p.alive = false
	return previous
}

func (p *basePiece) getAlive() bool {
	return p.alive
}

/***************************

Pawn implementation

***************************/

type Pawn struct {
	basePiece
	moved bool
}

func (p *Pawn) create(team int) {
	p.name = "Pawn"
	p.shortName = "P"
	p.alive = true
	p.team = team
	p.moved = false
}

func (p *Pawn) hasMoved() bool {
	return p.moved
}

func (p *Pawn) move(nPos coord, board *BoardData) bool {
	x_offset := float64(nPos.getX() - p.loc.getX())
	y_offset := float64(nPos.getY() - p.loc.getY())
	nPos_Piece := board.getCell(nPos).pPiece

	// Check that the offsets are valid
	if math.Abs(x_offset) > 1 {
		return false
	}

	if p.team == 1 && y_offset != -1 {
		return false
	}

	if p.team == 0 && y_offset != 1 {
		return false
	}

	// If the move is forward check that the forward square is empty
	if x_offset == 0 && nPos_Piece == nil {
		return true
	}

	// If the move is attacking check that the attacked squares non-empty
	if x_offset != 0 && nPos_Piece != nil {
		if nPos_Piece.getTeam()+p.team == 1 {
			return true
		}
	}

	// If the move is plausible return true
	return false
}

func (p *Pawn) specialMove(nPos coord, board *BoardData) bool {

	// If the pawn has been moved already return false
	if p.moved {
		return false
	}

	x_offset := float64(nPos.getX() - p.loc.getX())
	y_offset := float64(nPos.getY() - p.loc.getY())

	// Check that the offsets are valid
	if math.Abs(x_offset) != 0 {
		return false
	}

	if p.team == 1 && -1*y_offset == 2 {

		// Check to see if both positions in front are empty
		y := p.loc.getY() - 1
		x := nPos.getX()
		ternPos := coordData{x: x, y: y}
		tPos_Piece := board.getCell(ternPos).pPiece
		nPos_Piece := board.getCell(nPos).pPiece

		if nPos_Piece == nil && tPos_Piece == nil {
			return true
		}
	}

	if p.team == 0 && y_offset == 2 {

		// Check to see if both positions in front are empty
		y := p.loc.getY() + 1
		x := nPos.getX()
		ternPos := coordData{x: x, y: y}
		tPos_Piece := board.getCell(ternPos).pPiece
		nPos_Piece := board.getCell(nPos).pPiece

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
	basePiece
	moved bool
}

func (r *Rook) create(team int) {
	r.name = "Rook"
	r.shortName = "R"
	r.alive = true
	r.team = team
	r.moved = false

}

func (r *Rook) hasMoved() bool {
	return r.moved
}

func (r *Rook) move(nPos coord, board *BoardData) bool {
	x_offset := float64(nPos.getX() - r.loc.getX())
	y_offset := float64(nPos.getY() - r.loc.getY())
	x_sign := int(x_offset / math.Abs(x_offset))
	y_sign := int(y_offset / math.Abs(y_offset))

	if x_offset+y_offset != x_offset || x_offset+y_offset != y_offset {
		return false
	}

	for x := 0; x < int(math.Abs(x_offset)); x++ {

		for y := 0; y < int(math.Abs(y_offset)); y++ {
			ternPos := coordData{x: r.loc.getX() + x_sign*x, y: r.loc.getY() + y_sign*y}
			tPos_Piece := board.getCell(ternPos).pPiece

			if tPos_Piece != nil {
				return false
			}
		}
	}

	nPos_Piece := board.getCell(nPos).pPiece

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+r.team == 1 {
		return true
	}

	return false
}

/***************************

Bishop implementation

***************************/

type Bishop struct {
	basePiece
}

func (b *Bishop) create(team int) {
	b.name = "Bishop"
	b.shortName = "B"
	b.alive = true
	b.team = team
}

func (b *Bishop) hasMoved() bool {
	return true
}

func (b *Bishop) move(nPos coord, board *BoardData) bool {
	x_offset := float64(nPos.getX() - b.loc.getX())
	y_offset := float64(nPos.getY() - b.loc.getY())
	x_sign := int(x_offset / math.Abs(x_offset))
	y_sign := int(y_offset / math.Abs(y_offset))

	if math.Abs(x_offset) != math.Abs(y_offset) {
		return false
	}

	for offset := 0; offset < int(math.Abs(x_offset)); offset++ {

		ternPos := coordData{x: b.loc.getX() + x_sign*offset, y: b.loc.getY() + y_sign*offset}
		tPos_Piece := board.getCell(ternPos).pPiece

		if tPos_Piece != nil {
			return false
		}
	}

	nPos_Piece := board.getCell(nPos).pPiece

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+b.team == 1 {
		return true
	}

	return false

}

/***************************

Knight implementation

***************************/

type Knight struct {
	basePiece
}

func (n *Knight) create(team int) {
	n.name = "Knight"
	n.shortName = "N"
	n.alive = true
	n.team = team

}

func (n *Knight) hasMoved() bool {
	return true
}

func (n *Knight) move(nPos coord, board *BoardData) bool {
	x_offset := float64(nPos.getX() - n.loc.getX())
	y_offset := float64(nPos.getY() - n.loc.getY())
	total_offset := int(math.Abs(x_offset) + math.Abs(y_offset))

	if x_offset == 0 || y_offset == 0 {
		return false
	}

	if total_offset != 3 {
		return false
	}

	nPos_Piece := board.getCell(nPos).pPiece

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+n.team == 1 {
		return true
	}

	return false
}

/***************************

Queen implementation

***************************/

type Queen struct {
	basePiece
	Bishop
	Rook
}

func (q *Queen) create(team int) {
	q.name = "Queen"
	q.shortName = "Q"
	q.alive = true
	q.team = team

}

func (q *Queen) hasMoved() bool {
	return true
}

func (q *Queen) move(nPos coord, board *BoardData) bool {
	return q.Bishop.move(nPos, board) || q.Rook.move(nPos, board)
}

/***************************

King implementation

***************************/

type King struct {
	basePiece
	moved bool
}

func (k *King) create(team int) {
	k.name = "King"
	k.shortName = "K"
	k.alive = true
	k.team = team
	k.moved = false

}

func (k *King) hasMoved() bool {
	return k.moved
}

func (k *King) move(nPos coord, board *BoardData) bool {
	x_offset := float64(nPos.getX() - k.loc.getX())
	y_offset := float64(nPos.getY() - k.loc.getY())
	total_offset := int(math.Abs(x_offset) + math.Abs(y_offset))

	if math.Abs(x_offset) > 1 {
		return false
	}

	if math.Abs(y_offset) > 1 {
		return false
	}

	if total_offset == 0 {
		return false
	}

	nPos_Piece := board.getCell(nPos).pPiece

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+k.team == 1 {
		return true
	}

	return false
}

func (k *King) specialMove(nPos coord, board *BoardData) bool {
	if k.moved {
		return false
	}

	castleK_x := 6
	castleQ_x := 2
	castle_y := 0
	if k.team == 1 {
		castle_y = 7
	}

	if nPos.getY() != castle_y {
		return false
	}

	if nPos.getX() != castleQ_x && nPos.getX() != castleK_x {
		return false
	}

	begin := int(math.Min(float64(k.loc.getX()), float64(nPos.getX())))
	end := int(math.Max(float64(k.loc.getX()), float64(nPos.getX())))

	for x := begin; x <= end; x++ {
		ternPos := coordData{x: x, y: castle_y}
		tPos_Piece := board.getCell(ternPos).pPiece

		if tPos_Piece != nil {
			return false
		}
		if cellAttacked(board, ternPos, (k.team+1)%2) {
			return false
		}
	}

	if nPos.getX() == 2 {
		rookPos := coordData{x: 0, y: castle_y}
		rook_Piece := board.getCell(rookPos).pPiece

		if rook_Piece.getName() == "Rook" {
			if !rook_Piece.hasMoved() {
				return true
			}
		}
	} else {
		rookPos := coordData{x: 7, y: castle_y}
		rook_Piece := board.getCell(rookPos).pPiece

		if rook_Piece.getName() == "Rook" {
			if !rook_Piece.hasMoved() {
				return true
			}
		}
	}

	return false
}
