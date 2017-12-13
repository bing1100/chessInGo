package main

/***************************

Board and Cell structure

***************************/

type Cell struct {
	vPos   coord
	pPiece Piece
}

type BoardData struct {
	pCells    [8][8]*Cell
	pieces    [2][16]Piece
	whoseMove int
}

type Board interface {
	createBoard()
	getCell(coord coord) *Cell
	getPiece(team int, idx int) Piece
	setPiece(cPos coord, piece Piece) bool
	getNext() int
	move(cPos coord, nPos coord) bool
}

func (b *BoardData) getCell(coord coord) *Cell {
	return b.pCells[coord.getX()][coord.getY()]
}

func (b *BoardData) getPiece(team int, idx int) Piece {
	return b.pieces[team][idx]
}

func (b *BoardData) setPiece(cPos coord, piece Piece) bool {
	cell := b.pCells[cPos.getX()][cPos.getY()]
	if cell.pPiece != nil {
		cell.pPiece.setDead()
		b.pCells[cPos.getX()][cPos.getY()].pPiece = piece
		return false
	}
	b.pCells[cPos.getX()][cPos.getY()].pPiece = piece
	return true
}

func (b *BoardData) getNext() int {
	return b.whoseMove
}

/***************************

Interface to create new board
and setup pieces

***************************/
func createCell(x int, y int) *Cell {
	cell := new(Cell)

	coord := coordData{x: x, y: y}

	cell.vPos = coord
	cell.pPiece = nil

	return cell
}

func createPieces(team int, board *BoardData) [16]Piece {
	pieces := *(new([16]Piece))
	for i := 0; i < 8; i++ {
		pieces[i] = new(Pawn)
	}

	pieces[8] = new(Rook)
	pieces[9] = new(Knight)
	pieces[10] = new(Bishop)
	pieces[11] = new(Queen)
	pieces[12] = new(King)
	pieces[13] = new(Bishop)
	pieces[14] = new(Knight)
	pieces[15] = new(Rook)

	for i := 0; i < 16; i++ {
		pieces[i].create(team)
	}

	pawnRow := 1
	offset := -1
	if team == 1 {
		pawnRow = 6
		offset = 1
	}
	currIdx := 0

	for shift := 0; shift <= 1; shift++ {
		for x_coord := 0; x_coord < 8; x_coord++ {
			y_coord := shift*offset + pawnRow
			board.setPiece(coordData{x: x_coord, y: y_coord}, pieces[currIdx])
			if !pieces[currIdx].setLoc(board, coordData{x: x_coord, y: y_coord}) {
				print("ERROR: Setting locations when creating board failed")
				return pieces
			}
			currIdx++
		}
	}

	return pieces
}

func (b *BoardData) createBoard() {

	b.pCells = *(new([8][8]*Cell))
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			b.pCells[x][y] = createCell(x, y)
		}
	}

	b.pieces[0] = createPieces(0, b)
	b.pieces[1] = createPieces(1, b)
	b.whoseMove = 1
}

/***************************

Interface to keep track of board cells

***************************/

func cellAttacked(cBoa *BoardData, coor coord, team int) bool {

	for pieceIdx := 0; pieceIdx <= 15; pieceIdx++ {
		piece := cBoa.getPiece(team, pieceIdx)
		if piece.move(coor, cBoa) {
			return true
		}
	}

	return false
}

/***************************

Interface to keep track of board state

***************************/

func boardValidState(cBoa *BoardData) bool {
	loc := cBoa.pieces[cBoa.whoseMove][12].getLoc()
	return cellAttacked(cBoa, loc, (cBoa.whoseMove+1)%2)
}

func checkStaleMate(cBoa *BoardData) bool {
	loc := cBoa.pieces[cBoa.whoseMove][12].getLoc()
	for x_offset := -1; x_offset <= -1; x_offset++ {
		for y_offset := -1; y_offset <= -1; y_offset++ {
			if x_offset != 0 && y_offset != 0 {
				x := loc.getX() + x_offset
				y := loc.getY() + y_offset

				if (x >= 0 && x <= 7) && (y >= 0 && y <= 7) {
					if !cellAttacked(cBoa, coordData{x: x, y: y}, (cBoa.whoseMove+1)%2) {
						return false
					}
				}
			}
		}
	}
	return true
}

func checkCheckMate(cBoa *BoardData) bool {
	loc := cBoa.pieces[cBoa.whoseMove][12].getLoc()
	if checkStaleMate(cBoa) {
		return cellAttacked(cBoa, loc, (cBoa.whoseMove+1)%2)
	}
	return false
}

/****************************

Move function that moves piece

****************************/
func (b *BoardData) move(cPos coord, nPos coord) bool {
	cPiece := b.getCell(cPos).pPiece
	tempBoard := Copy(b).(*BoardData)

	if cPiece == nil {

		return false

	}

	if cPiece.move(nPos, b) || cPiece.specialMove(nPos, b) {
		tempBoard.setPiece(nPos, cPiece)
		tempBoard.pCells[cPos.getX()][cPos.getY()].pPiece = nil

		if boardValidState(tempBoard) {
			b.setPiece(nPos, cPiece)
			b.pCells[cPos.getX()][cPos.getY()].pPiece = nil
			return true
		}
	}

	return false
}
