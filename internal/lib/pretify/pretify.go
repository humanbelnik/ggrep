package pretify

import (
	"fmt"
)

const (
	Reset = "\033[0m"

	Black        = 30
	Red          = 31
	Green        = 32
	Yellow       = 33
	Blue         = 34
	Magenta      = 35
	Cyan         = 36
	LightGray    = 37
	DarkGray     = 90
	LightRed     = 91
	LightGreen   = 92
	LightYellow  = 93
	LightBlue    = 94
	LightMagenta = 95
	LightCyan    = 96
	White        = 97
)

func WithColor(code int, s string) string {
	return fmt.Sprintf("\033[%dm%s%s", code, s, Reset)
}

func LanguageWithColor(lang string) string {
	switch lang {
	case "C":
		return fmt.Sprintf("\033[%dm%s%s", DarkGray, lang, Reset)
	case "Python":
		return fmt.Sprintf("\033[%dm%s%s", Green, lang, Reset)
	case "C#":
		return fmt.Sprintf("\033[%dm%s%s", Magenta, lang, Reset)
	case "C++":
		return fmt.Sprintf("\033[%dm%s%s", LightRed, lang, Reset)
	case "Go":
		return fmt.Sprintf("\033[%dm%s%s", Cyan, lang, Reset)
	case "Assembly":
		return fmt.Sprintf("\033[%dm%s%s", LightYellow, lang, Reset)
	default:
		return lang
	}
}
