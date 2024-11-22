package color

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func RedText(text string) string {
	return Red + text + Reset
}

func YellowText(text string) string {
	return Yellow + text + Reset
}
