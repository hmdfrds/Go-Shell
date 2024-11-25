package shell

import (
	"os"
	"testing"
)

func TestNewShell(t *testing.T) {

	currentDir, _ := os.Getwd()
	shell := NewShell(currentDir)

	if shell.currentDir != currentDir {
		t.Errorf("Expected currentDir to be %s, got %s", currentDir, shell.currentDir)
	}
	if len(shell.commandRegistry) != 0 {
		t.Errorf("Expected commandRegistry to be empty, got %d commands", len(shell.commandRegistry))
	}
	if len(shell.history) != 0 {
		t.Errorf("Expected history to be empty, got %d items", len(shell.history))
	}
}

func TestRegisterCommands(t *testing.T) {
	currentDir, _ := os.Getwd()
	shell := NewShell(currentDir)
	shell.RegisterCommands()

	expectedCommands := []string{"cd", "pwd", "exit", "ls", "echo"}

	for _, cmd := range expectedCommands {
		if _, exits := shell.commandRegistry[cmd]; !exits {
			t.Errorf("Command %s was not registered", cmd)
		}
	}
}

func TestParseCommand(t *testing.T) {
	tests := []struct {
		input       string
		expected    []string
		shouldError bool
	}{
		{"echo Hello World", []string{"echo", "Hello", "World"}, false},
		{"echo \"Hello World\"", []string{"echo", "Hello World"}, false},
		{"echo 'Hello World'", []string{"echo", "Hello World"}, false},
		{"echo \"Unclosed quote", nil, true},
		{"echo 'unclosed quote", nil, true},
	}

	for _, test := range tests {
		result, err := parseCommand(test.input)
		if (err != nil) != test.shouldError {
			t.Errorf("Unexpected error status for input '%s' : %v", test.expected, err)
		}
		if !test.shouldError && !equalSlices(result, test.expected) {
			t.Errorf("Expected %v, got %v for input '%s'", test.expected, result, test.input)
		}
	}
}

func TestParseRedirection(t *testing.T) {
	tests := []struct {
		input        string
		expectedCmd  string
		expectedFile string
		appendMode   bool
	}{
		{"echo Hello > output.txt", "echo Hello", "output.txt", false},
		{"echo Hello >> output.txt", "echo Hello", "output.txt", true},
		{"echo Hello", "echo Hello", "", false},
	}

	for _, test := range tests {
		cmd, file, appendMode := parseRedirection(test.input)
		if cmd != test.expectedCmd || file != test.expectedFile || appendMode != test.appendMode {
			t.Errorf("For input '%s', expected cmd='%s', appendMode=%v; got cmd='%s', file='%s', appendMode=%v",
				test.input, test.expectedCmd, test.appendMode, cmd, file, appendMode)
		}
	}
}

func TestHandleFileRedirection(t *testing.T) {
	testFile := "test_output.txt"
	defer os.Remove(testFile)

	output := "Hello, file!"
	handleFileRedirection(false, testFile, output)

	content, _ := os.ReadFile(testFile)
	if string(content) != output {
		t.Errorf("Expected file content '%s', got '%s'", output, content)
	}
}

func TestGetHistoryInput(t *testing.T) {
	history := []string{"cmd1", "cmd2", "cmd3"}
	temp := "temp"
	tests := []struct {
		currentHistory int
		expected       string
	}{
		{-1, "temp"},
		{0, "cmd3"},
		{1, "cmd2"},
	}

	for _, test := range tests {
		result := getHistoryInput(temp, test.currentHistory, history)
		if result != test.expected {
			t.Errorf("Expected '%s', got '%s'", test.expected, result)
		}
	}
}

func TestShellIntegration(t *testing.T) {
	currentDir, _ := os.Getwd()
	shell := NewShell(currentDir)
	shell.RegisterCommands()

	output, err := shell.executeCommand("pwd")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if output != currentDir {
		t.Errorf("Expected '%s', got '%s'", currentDir, output)
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
