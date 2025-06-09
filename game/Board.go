package game

func NewBoard() Board {
	board := &Board{}
	board.init()
	return *board
}

type Board struct {
	cells [15][7]Cell
}

func (board *Board) init() {
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
	for col := 'A'; col <= 'O'; col++ {
		for row := 1; row <= 7; row++ {
			cell := board.GetCell(col, row)
			cell.Checked = false
			cell.Color = colors[colorIndex]
			colorIndex++
		}
	}
}

func (board *Board) GetCell(col rune, row int) *Cell {
	println("GetCell(rune,int)")
	return board.getCell(int(col-'A'), row-1)
}

func (board *Board) getCell(col, row int) *Cell {
	println("GetCell(int,int)")
	return &board.cells[col][row]
}
