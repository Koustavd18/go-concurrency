package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_balanceBank(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	balaceBank()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "34320") {
		t.Error("[Test]: Result mismatch")
	}

}
