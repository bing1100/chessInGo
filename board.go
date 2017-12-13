package main

/***************************

Board and Cell structure

***************************/

type Cell struct {
	vPos   coord
	pPiece Piece
}

type Board struct {
	pCells    [8][8]*Cell
	pieces    [2][16]Piece
	whoseMove int
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

func createPieces(team int, board *Board) [16]Piece {
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
		for y_coord := 0; y_coord < 8; y_coord++ {
			x_coord := shift*offset + pawnRow

			board.pCells[x_coord][y_coord].pPiece = pieces[currIdx]
			if !pieces[currIdx].setLoc(board, coordData{x: x_coord, y: y_coord}) {
				print("ERROR: Setting locations when creating board failed")
				return pieces
			}
			currIdx++
		}
	}

	return pieces
}

func createBoard() *Board {

	board := new(Board)

	board.pCells = *(new([8][8]*Cell))
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			board.pCells[x][y] = createCell(x, y)
		}
	}

	board.pieces[0] = createPieces(0, board)
	board.pieces[1] = createPieces(1, board)
	board.whoseMove = 1

	return board
}

/***************************

Interface to keep track of board cells

***************************/

func pieceAtCell(cBoa *Board, coor coord) Piece {
	return (*((*cBoa).pCells)[coor.getX()][coor.getY()]).pPiece
}

func cellAttacked(cBoa *Board, coor coord, team int) bool {

	for pieceIdx := 0; pieceIdx <= 15; pieceIdx++ {
		piece := cBoa.pieces[team][pieceIdx]
		if piece.move(coor, cBoa) {
			return true
		}
	}

	return false
}

func killPiece(cBoa *Board, coor coord) {
	cell := (*((*cBoa).pCells)[coor.getX()][coor.getY()])
	if cell.pPiece != nil {
		if !(cell.pPiece.setDead()) {
			println("Piece in cell was already dead")
		}
	} else {
		println("Tried to kill empty cell")
	}

	cell.pPiece = nil
}

/***************************

Interface to keep track of board state

***************************/

func boardValidState(cBoa *Board) bool {
	kLoc := cBoa.pieces[cBoa.whoseMove][12].getLoc()
	return cellAttacked(cBoa, kLoc, (cBoa.whoseMove+1)%2)
}

func checkStaleMate(cBoa *Board) bool {
	kLoc := cBoa.pieces[cBoa.whoseMove][12].getLoc()
	for x_offset := -1; x_offset <= -1; x_offset++ {
		for y_offset := -1; y_offset <= -1; y_offset++ {
			if x_offset != 0 && y_offset != 0 {
				x := kLoc.getX() + x_offset
				y := kLoc.getY() + y_offset

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

func checkCheckMate(cBoa *Board) bool {
	kLoc := cBoa.pieces[cBoa.whoseMove][12].getLoc()
	if checkStaleMate(cBoa) {
		return cellAttacked(cBoa, kLoc, (cBoa.whoseMove+1)%2)
	}
	return false
}

/****************************

Move function that moves piece

****************************/
func move(cPos coord, nPos coord, board *Board) res {
	cPiece := pieceAtCell(board, cPos)
	failure := res{false, board}
	nBoard := Copy(board)
	success := res{true, nBoard.(*Board)}

	if cPiece == nil {

		return failure

	}
	return success

}
