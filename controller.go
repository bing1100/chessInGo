package main

func main() {
	board := new(BoardData)
	board.createBoard()
	showASCII(board)
	showHTML(board)

	println(board.getCell(coordData{x: 0, y: 1}).pPiece.getChar())
	board.move(coordData{x: 0, y: 1}, coordData{x: 0, y: 3})
	showASCII(board)
}
