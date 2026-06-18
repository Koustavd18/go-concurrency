package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("test message")
	wg.Wait()

	if msg != "test message" {
		t.Errorf("Expected 'test message', but got '%s'", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "test output"
	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "test output") {
		t.Errorf("Expected 'test output', but got '%s'", output)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("Expected Hello, universe!, not found")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("Expected Hello, cosmos!, not found")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected Hello, world!, not found")
	}
}
