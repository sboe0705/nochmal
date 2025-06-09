package game

import (
	"testing"
)

func TestBoardColorsRandomly(t *testing.T) {
	// given
	board := NewBoard()

	// when ... then
	assertColor(board, 'A', 1, false, Green, t)
	assertColor(board, 'A', 2, false, Orange, t)
	assertColor(board, 'A', 3, false, Blue, t)

	assertColor(board, 'E', 5, false, Orange, t)
	assertColor(board, 'F', 5, false, Pink, t)
	assertColor(board, 'G', 5, false, Blue, t)
}

func assertColor(board Board, col rune, row int, checked bool, color Color, t *testing.T) {
	cell := board.GetCell(col, row)
	if cell.Checked != checked {
		if cell.Checked {
			t.Errorf("Cell %c%v is checked but should not be!", col, row)
		} else {
			t.Errorf("Cell %c%v is not checked but should be!", col, row)
		}
	}
	if cell.Color != color {
		t.Errorf("Cell %c%v has the wrong color (expected: %v, actual: %v)!", col, row, color.ToString(), cell.Color.ToString())
	}
}
