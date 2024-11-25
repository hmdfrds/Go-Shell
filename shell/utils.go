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
	currentHistory := -1
	var input string
	var temp string
	for {
		buf, n, err := readKeyPress()
		if err != nil {
			break
		}
		switch {
		case n == 1:
			key := buf[0]
			switch key {
			case '\r': // return key
				println()
				resetSuggestion(&suggestions, &currentSuggestion)
				return input
			case '\t': // tab key
				input = handleTab(input, &suggestions, &currentSuggestion)
				rePrint(s.currentDir + "> " + input)
			case 127: // backspace key
				if len(input) > 0 {
					input = input[:len(input)-1]
					rePrint(s.currentDir + "> " + input)
				}
				resetSuggestion(&suggestions, &currentSuggestion)
			default: // any other keys
				input += string(key)
				fmt.Print(string(key))
				resetSuggestion(&suggestions, &currentSuggestion)
			}
			temp = input
		case n == 3 && buf[0] == '\x1b' && buf[1] == '[' && len(s.history) > 0:
			switch buf[2] {
			case 'A':
				if currentHistory < len(s.history)-1 {
					currentHistory++
				}
			case 'B':
				if currentHistory > -1 {
					currentHistory--
				}
			}
			input = getHistoryInput(temp, currentHistory, s.history)

			rePrint(s.currentDir + "> " + input)
		}

	}

	return input
}

func getHistoryInput(temp string, currentHistory int, history []string) string {
	if currentHistory == -1 {
		return temp
	}
	return history[len(history)-1-currentHistory]
}

func handleTab(input string, suggestions *[]string, currentSuggestion *int) string {

	strSlice := strings.Split(input, " ")
	base := strings.Join(strSlice[:len(strSlice)-1], " ") // get all the first part
	path := strSlice[len(strSlice)-1]                     // last part of input

	if len(*suggestions) == 0 {
		*suggestions = getSuggestionDirectories(path)
	}
	if len(*suggestions) > 0 {
		suggestion := (*suggestions)[*currentSuggestion]
		*currentSuggestion = (*currentSuggestion + 1) % len(*suggestions)
		input = formatCompletedInput(base, suggestion)
	}
	return input
}
func formatCompletedInput(base, suggestion string) string {
	if base != "" {
		base += " "
	}
	return base + suggestion
}

func resetSuggestion(suggestion *[]string, currentSuggestion *int) {
	*suggestion = nil
	*currentSuggestion = 0
}
func rePrint(str string) {
	// \033[2K clear the whole line
	// \r put cursor at start of the line
	// \b put cursor at the last of the word
	fmt.Print("\033[2K\r" + str + " \b")
}

// read one key at a time
func readKeyPress() ([]byte, int, error) {
	buf := make([]byte, 3)
	n, err := os.Stdin.Read(buf)
	return buf, n, err
}

func parseCommand(input string) ([]string, error) {
	var args []string
	var current strings.Builder
	inDoubleQuotes := false
	inSingleQuotes := false

	for i, ch := range input {
		switch ch {
		case ' ':
			// Append to args if not inside quotes
			if !inDoubleQuotes && !inSingleQuotes {
				if current.Len() > 0 {
					args = append(args, current.String())
					current.Reset()
				}
			} else {
				current.WriteRune(ch)
			}

		case '"':
			if inSingleQuotes {
				// Inside single quotes, treat as regular character
				current.WriteRune(ch)
			} else {
				// Toggle double quotes
				inDoubleQuotes = !inDoubleQuotes
				if !inDoubleQuotes && current.Len() > 0 {
					args = append(args, current.String())
					current.Reset()
				}
			}

		case '\'':
			if inDoubleQuotes {
				// Inside double quotes, treat as regular character
				current.WriteRune(ch)
			} else {
				// Toggle single quotes
				inSingleQuotes = !inSingleQuotes
				if !inSingleQuotes && current.Len() > 0 {
					args = append(args, current.String())
					current.Reset()
				}
			}

		case '\\':
			// Check for escaped quotes
			if i+1 < len(input) {
				next := input[i+1]
				if (next == '"' && inDoubleQuotes) || (next == '\'' && inSingleQuotes) {
					current.WriteRune(rune(next))
					i++ // Skip the escaped character
				} else {
					current.WriteRune(ch)
				}
			} else {
				current.WriteRune(ch)
			}
		default:
			current.WriteRune(ch)
		}

		// Add the last argument if at the end
		if i == len(input)-1 && current.Len() > 0 {
			args = append(args, current.String())
		}
	}

	// Check for unclosed quotes
	if inDoubleQuotes || inSingleQuotes {
		return nil, fmt.Errorf("unclosed quote in input")
	}

	return args, nil
}
