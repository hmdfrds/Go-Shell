package shell

import (
	"os"
	"testing"
)

func TestCmdPwd(t *testing.T) {
	currentDir, _ := os.Getwd()
	output, err := cmdPwd(nil)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if output != currentDir {
		t.Errorf("Expected output '%s', got '%s'", currentDir, output)
	}
}

func TestCmdLs(t *testing.T) {
	os.Mkdir("test_dir", 0755)
	defer os.RemoveAll("test_dir")

	_, err := cmdLs(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCmdEcho(t *testing.T) {
	args := []string{"Hello", "World"}

	expected := "Hello World"
	output, err := cmdEcho(args)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if output != expected {
		t.Errorf("Expected output '%s', got '%s'", expected, output)
	}
}
