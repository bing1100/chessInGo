package main

import (
	"os"
)

// ASCII print of the board object
func showASCII(board *Board) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := board.pCells[x][y].pPiece
			if piece != nil {
				print(piece.getChar())
			} else {
				print("+")
			}
		}
		print("\n")
	}
}

// HTML print of the board object
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func showHTML(board *Board) {
	f, err := os.Create("/home/bhux/Downloads/chess/index.html")
	check(err)
	defer f.Close()

	//Print the header
	f.WriteString("<! DOCTYPE html>\n")
	f.WriteString("<html>\n")

	//style elements
	f.WriteString("<head><style>\n")
	f.WriteString("table#board {width: 800px;height: 800px;}\n")
	f.WriteString("td#cB {background-color: brown;text-align: center;font-size: 40px;width: 100px;height: 100px}\n")
	f.WriteString("td#cW {background-color: rgb(196, 154, 103);text-align: center;font-size: 40px;width: 100px;height: 100px}\n")
	f.WriteString("</style></head>\n")

	//Begining of body
	f.WriteString("<body><table id=\"board\">\n")
	cellTypes := []string{"cB", "cW"}
	textColor := []string{"black", "white"}
	idx := 0
	for x := 0; x < 8; x++ {
		f.WriteString("<tr>\n")
		idx++
		for y := 0; y < 8; y++ {
			piece := board.pCells[x][y].pPiece
			cType := cellTypes[idx%2]

			f.WriteString("<td id=\"" + cType + "\">")
			if piece != nil {
				tColor := textColor[piece.getTeam()]
				f.WriteString("<font color=\"" + tColor + "\">" + piece.getChar() + "</font>")
			}
			f.WriteString("</td>\n")
			idx++
		}
		f.WriteString("</tr>\n")
	}

	f.WriteString("</table></body>\n")

	// End of html
	f.WriteString("</html>")

	f.Sync()

}
