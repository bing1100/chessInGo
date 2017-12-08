package main

/***************************

Board and Cell structure

***************************/

type Cell struct {
	vPos   *coord
	pPiece *Piece
}

type Board struct {
	pCells       [8][8]*Cell
	pWhitePieces [16]*Piece
	pBlackPieces [16]*Piece
	whoseMove    int
}

/***************************

Interface to create new board
and setup pieces

***************************/
func createCell(x int, y int) *Cell {
	cell := new(Cell)

	coord := createCoord(x, y)

	cell.vPos = coord
	cell.pPiece = nil

	return cell
}

func createPieces(team int) [16]*Piece {
	return *(new([16]*Piece))
}

func setPieces(boar *Board, bPieces [16]*Piece, wPieces [16]*Piece) {

}

func createBoard() *Board {

	board := new(Board)

	board.pCells = *(new([8][8]*Cell))
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			board.pCells[x][y] = createCell(x, y)
		}
	}

	board.pBlackPieces = createPieces(0)
	board.pWhitePieces = createPieces(1)
	board.whoseMove = 1

	setPieces(board, board.pBlackPieces, board.pWhitePieces)

	return board
}

/***************************

Interface to keep track of board cells

***************************/

func pieceAtCell(cBoa *Board, coor coord) *Piece {
	return (*((*cBoa).pCells)[coor.x][coor.y]).pPiece
}

func cellAttacked(cBoa *Board, coor coord) bool {
	return true
}

func boardValidState(cBoa *Board) {

}
