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

	actions *Piece
}

type Piece interface {
	create(team int)
	move(cPos coord, nPos coord, board *Board) bool
	specialMove(cPos coord, nPos coord, board *Board) bool
	hasMoved() bool
	getTeam() int
	getName() string
	getChar() string
}

func (p *basePiece) create(team int) {
	p.actions = nil
	print("ERROR: trying to create a piece object")
}

func (p *basePiece) move(cPos coord, nPos coord, board *Board) bool {
	return p.move(cPos, nPos, board)
}

func (p *basePiece) specialMove(cPos coord, nPos coord, board *Board) bool {
	return p.specialMove(cPos, nPos, board)
}

func (p *basePiece) hasMoved() bool {
	return p.hasMoved()
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

/***************************

Pawn implementation

***************************/

type Pawn struct {
	basePiece
	moved bool

	actions *Piece
}

func (p *Pawn) create(team int) {
	p.name = "Pawn"
	p.shortName = "P"
	p.alive = true
	p.team = team
	p.moved = false

	p.basePiece.actions = p.actions
}

func (p *Pawn) hasMoved() bool {
	return p.moved
}

func (p *Pawn) move(cPos coord, nPos coord, board *Board) bool {
	x_offset := float64(nPos.x - cPos.x)
	y_offset := float64(nPos.y - cPos.y)
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
		if nPos_Piece.getTeam()+p.team == 1 {
			return true
		}
	}

	// If the move is plausible return true
	return false
}

func (p *Pawn) specialMove(cPos coord, nPos coord, board *Board) bool {

	// If the pawn has been moved already return false
	if p.moved {
		return false
	}

	x_offset := float64(nPos.x - cPos.x)
	y_offset := float64(nPos.y - cPos.y)

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
	basePiece
	moved bool

	actions *Piece
}

func (r *Rook) create(team int) {
	r.name = "Rook"
	r.shortName = "R"
	r.alive = true
	r.team = team
	r.moved = false

	r.basePiece.actions = r.actions
}

func (r *Rook) hasMoved() bool {
	return r.moved
}

func (r *Rook) move(cPos coord, nPos coord, board *Board) bool {
	x_offset := float64(nPos.x - cPos.x)
	y_offset := float64(nPos.y - cPos.y)
	x_sign := int(x_offset / math.Abs(x_offset))
	y_sign := int(y_offset / math.Abs(y_offset))

	if x_offset+y_offset != x_offset || x_offset+y_offset != y_offset {
		return false
	}

	for x := 0; x < int(math.Abs(x_offset)); x++ {

		for y := 0; y < int(math.Abs(y_offset)); y++ {
			ternPos := coord{x: cPos.x + x_sign*x, y: cPos.y + y_sign*y}
			tPos_Piece := pieceAtCell(board, ternPos)

			if tPos_Piece != nil {
				return false
			}
		}
	}

	nPos_Piece := pieceAtCell(board, nPos)

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+r.team == 1 {
		return true
	}

	return false
}

func (r *Rook) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

Bishop implementation

***************************/

type Bishop struct {
	basePiece

	actions *Piece
}

func (b *Bishop) create(team int) {
	b.name = "Bishop"
	b.shortName = "B"
	b.alive = true
	b.team = team

	b.basePiece.actions = b.actions
}

func (b *Bishop) hasMoved() bool {
	return true
}

func (b *Bishop) move(cPos coord, nPos coord, board *Board) bool {
	x_offset := float64(nPos.x - cPos.x)
	y_offset := float64(nPos.y - cPos.y)
	x_sign := int(x_offset / math.Abs(x_offset))
	y_sign := int(y_offset / math.Abs(y_offset))

	if math.Abs(x_offset) != math.Abs(y_offset) {
		return false
	}

	for offset := 0; offset < int(math.Abs(x_offset)); offset++ {

		ternPos := coord{x: cPos.x + x_sign*offset, y: cPos.y + y_sign*offset}
		tPos_Piece := pieceAtCell(board, ternPos)

		if tPos_Piece != nil {
			return false
		}
	}

	nPos_Piece := pieceAtCell(board, nPos)

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+b.team == 1 {
		return true
	}

	return false

}

func (b *Bishop) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

Knight implementation

***************************/

type Knight struct {
	basePiece

	actions *Piece
}

func (n *Knight) create(team int) {
	n.name = "Knight"
	n.shortName = "N"
	n.alive = true
	n.team = team

	n.basePiece.actions = n.actions
}

func (n *Knight) hasMoved() bool {
	return true
}

func (n *Knight) move(cPos coord, nPos coord, board *Board) bool {
	x_offset := float64(nPos.x - cPos.x)
	y_offset := float64(nPos.y - cPos.y)
	total_offset := int(math.Abs(x_offset) + math.Abs(y_offset))

	if x_offset == 0 || y_offset == 0 {
		return false
	}

	if total_offset != 3 {
		return false
	}

	nPos_Piece := pieceAtCell(board, nPos)

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+n.team == 1 {
		return true
	}

	return false
}

func (n *Knight) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

Queen implementation

***************************/

type Queen struct {
	basePiece
	Bishop
	Rook

	actions *Piece
}

func (q *Queen) create(team int) {
	q.name = "Queen"
	q.shortName = "Q"
	q.alive = true
	q.team = team

	q.basePiece.actions = q.actions
}

func (q *Queen) hasMoved() bool {
	return true
}

func (q *Queen) move(cPos coord, nPos coord, board *Board) bool {
	return q.Bishop.move(cPos, nPos, board) || q.Rook.move(cPos, nPos, board)
}

func (q *Queen) specialMove(cPos coord, nPos coord, board *Board) bool {
	return false
}

/***************************

King implementation

***************************/

type King struct {
	basePiece
	moved bool

	actions *Piece
}

func (k *King) create(team int) {
	k.name = "King"
	k.shortName = "K"
	k.alive = true
	k.team = team
	k.moved = false

	k.basePiece.actions = k.actions
}

func (k *King) hasMoved() bool {
	return k.moved
}

func (k *King) move(cPos coord, nPos coord, board *Board) bool {
	x_offset := float64(nPos.x - cPos.x)
	y_offset := float64(nPos.y - cPos.y)
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

	nPos_Piece := pieceAtCell(board, nPos)

	if nPos_Piece == nil {
		return true
	} else if nPos_Piece.getTeam()+k.team == 1 {
		return true
	}

	return false
}

func (k *King) specialMove(cPos coord, nPos coord, board *Board) bool {
	if k.moved {
		return false
	}

	castleK_x := 6
	castleQ_x := 2
	castle_y := 0
	if k.team == 1 {
		castle_y = 7
	}

	if nPos.y != castle_y {
		return false
	}

	if nPos.x != castleQ_x && nPos.x != castleK_x {
		return false
	}

	begin := int(math.Min(float64(cPos.x), float64(nPos.x)))
	end := int(math.Max(float64(cPos.x), float64(nPos.x)))

	for x := begin; x <= end; x++ {
		ternPos := coord{x: x, y: castle_y}
		tPos_Piece := pieceAtCell(board, ternPos)

		if tPos_Piece != nil {
			return false
		}
		if cellAttacked(board, ternPos) {
			return false
		}
	}

	if nPos.x == 2 {
		rookPos := coord{x: 0, y: castle_y}
		rook_Piece := pieceAtCell(board, rookPos)

		if rook_Piece.getName() == "Rook" {
			if !rook_Piece.hasMoved() {
				return true
			}
		}
	} else {
		rookPos := coord{x: 7, y: castle_y}
		rook_Piece := pieceAtCell(board, rookPos)

		if rook_Piece.getName() == "Rook" {
			if !rook_Piece.hasMoved() {
				return true
			}
		}
	}

	return false
}
