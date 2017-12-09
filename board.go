package main

/***************************

Board and Cell structure

***************************/

type Cell struct {
	vPos   coord
	pPiece Piece
}

type Board struct {
	pCells       [8][8]*Cell
	pWhitePieces [16]Piece
	pBlackPieces [16]Piece
	whoseMove    int
}

/***************************

Interface to create new board
and setup pieces

***************************/
func createCell(x int, y int) *Cell {
	cell := new(Cell)

	coord := coord{x: x, y: y}

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

	pawnRow := 2
	offset := -1
	if team == 1 {
		pawnRow = 6
		offset = 1
	}
	currIdx := 0

	for shift := 0; shift <= 1; shift++ {
		for x_coord := 0; x_coord < 8; x_coord++ {
			y_coord := shift*offset + pawnRow

			board.pCells[x_coord][y_coord].pPiece = pieces[currIdx]
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

	board.pBlackPieces = createPieces(0, board)
	board.pWhitePieces = createPieces(1, board)
	board.whoseMove = 1

	return board
}

/***************************

Interface to keep track of board cells

***************************/

func pieceAtCell(cBoa *Board, coor coord) Piece {
	return (*((*cBoa).pCells)[coor.x][coor.y]).pPiece
}

func cellAttacked(cBoa *Board, coor coord) bool {
	return true
}

func boardValidState(cBoa *Board) {

}
