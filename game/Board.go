package game

import (
	"fmt"
)

func NewBoard() Board {
	board := &Board{}
	board.init()
	return *board
}

type Board struct {
	cells [15][7]Cell
}

func (b *Board) init() {
	colors := []Color{
		Green, Orange, Blue, Blue, Pink, Pink, Yellow, // A
		Green, Green, Green, Pink, Orange, Blue, Yellow, // B
		Green, Yellow, Pink, Pink, Orange, Blue, Blue, // C
		Yellow, Green, Green, Green, Orange, Pink, Blue, // D
		Yellow, Yellow, Green, Orange, Orange, Pink, Blue, // E
		Yellow, Yellow, Green, Orange, Pink, Pink, Blue, // F
		Yellow, Orange, Green, Blue, Blue, Pink, Pink, // G
		Green, Orange, Pink, Blue, Blue, Yellow, Yellow, // H
		Blue, Pink, Pink, Green, Orange, Yellow, Yellow, // I
		Blue, Blue, Pink, Green, Orange, Orange, Yellow, // J
		Blue, Blue, Yellow, Yellow, Orange, Pink, Green, // K
		Orange, Orange, Yellow, Yellow, Pink, Blue, Green, // L
		Yellow, Orange, Orange, Orange, Pink, Blue, Green, // M
		Yellow, Green, Green, Pink, Pink, Blue, Orange, // N
		Yellow, Green, Green, Blue, Pink, Orange, Orange, // O
	}
	colorIndex := 0
	for col := 0; col < len(b.cells); col++ {
		for row := 0; row < len(b.cells[col]); row++ {
			cell := b.getCell(col, row)
			cell.Checked = false
			cell.Color = colors[colorIndex]
			colorIndex++
		}
	}
}

func (b *Board) GetCell(col rune, row int) *Cell {
	return b.getCell(int(col-'A'), row-1)
}

func (b *Board) getCell(col, row int) *Cell {
	return &b.cells[col][row]
}

func (b *Board) Print() {
	fmt.Println("  A B C D E F G H I J K L M N O")
	for row := 0; row < len(b.cells[0]); row++ {
		fmt.Print(row+1, " ")
		for col := 0; col < len(b.cells); col++ {
			cell := b.getCell(col, row)
			color := cell.Color
			fmt.Print(color.Code())
			if cell.Checked {
				fmt.Print("X ")
			} else {
				fmt.Print("□ ")
			}
			fmt.Print("\033[0m")
		}
		fmt.Println("")
	}
}
