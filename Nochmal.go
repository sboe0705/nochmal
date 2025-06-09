package main

import (
	"fmt"
	"nochmal/game"
)

func main() {
	fmt.Println("NOCHMAL Game Simulator")
	board := game.NewBoard()
	board.GetCell('G', 4).Checked = true
	board.GetCell('G', 5).Checked = true
	board.GetCell('H', 4).Checked = true
	board.GetCell('H', 5).Checked = true
	board.Print()
}
