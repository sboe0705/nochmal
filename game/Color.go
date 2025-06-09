package game

type Color int

const (
	Green Color = iota
	Yellow
	Blue
	Pink
	Orange
)

func (s Color) ToString() string {
	switch s {
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Pink:
		return "Pink"
	case Orange:
		return "Orange"
	default:
		return "Unknown"
	}
}
