package main

func showASCII(board *Board) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := board.pCells[x][y].pPiece
			if piece != nil {
				print(piece.getChar())
			} else {
				print(" ")
			}
		}
		print("\n")
	}
}
