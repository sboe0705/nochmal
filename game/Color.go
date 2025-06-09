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

func (s Color) Code() string {
	switch s {
	case Green:
		return "\033[32m"
	case Yellow:
		return "\033[33m"
	case Blue:
		return "\033[34m"
	case Pink:
		return "\033[35m"
	case Orange:
		return "\033[38;5;208m"
	default:
		return "Unknown"
	}
}
