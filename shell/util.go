package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getSuggestionDirectories(input string) []string {
	path := filepath.FromSlash(input) // convert all '/' to '\'
	baseDir, partial := ".", path

	if strings.Contains(path, "\\") {
		// base dir get all the path except the last after \.
		// Partial get the last after \.
		// If there's nothing after \, it just got empty string.
		baseDir, partial = filepath.Split(path)
	}

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil
	}

	var suggestion []string
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), partial) {
			suggestion = append(suggestion, filepath.Join(baseDir, entry.Name()))
		}
	}
	return suggestion
}

func (s *Shell) readInput() string {
	var suggestions []string
	currentSuggestion := 0
	var input string
	for {
		key, err := readKeyPress()
		if err != nil {
			break
		}

		switch key {
		case '\r': // return key
			println()
			resetSuggestion(&suggestions, &currentSuggestion)
			return input
		case '\t': // tab key
			input = handleTab(input, &suggestions, &currentSuggestion)
			rePrint(s.CurrentDir + "> " + input)
		case 127: // backspace key
			if len(input) > 0 {
				continue
			}
			resetSuggestion(&suggestions, &currentSuggestion)
			input = input[:len(input)-1]
			rePrint(s.CurrentDir + "> " + input)

		default: // any other keys
			resetSuggestion(&suggestions, &currentSuggestion)
			input += string(key)
			fmt.Print(string(key))
		}
	}

	return input
}

func handleTab(input string, suggestions *[]string, currentSuggestion *int) string {

	strSlice := strings.Split(input, " ")
	base := strings.Join(strSlice[:len(strSlice)-1], " ") // get all the first part
	path := strings.TrimSpace(strSlice[len(strSlice)-1])  // last part of input

	if len(*suggestions) == 0 {
		*suggestions = getSuggestionDirectories(path)
	}
	if len(*suggestions) > 0 {
		suggestion := (*suggestions)[*currentSuggestion]
		*currentSuggestion = (*currentSuggestion + 1) % len(*suggestions)
		input = base
		if len(base) > 0 {
			input += " "
		}
		input += suggestion
	}
	return input
}

func resetSuggestion(suggestion *[]string, currentSuggestion *int) {
	*suggestion = nil
	*currentSuggestion = 0
}
func rePrint(str string) {
	fmt.Print("\033[2K") // clear the whole line
	// \r put cursor at start of the line
	// \b put cursor at the last of the word
	fmt.Print("\r", str, " \b")
}

// read one key at a time
func readKeyPress() (byte, error) {
	buf := make([]byte, 1)
	_, err := os.Stdin.Read(buf)
	return buf[0], err
}
