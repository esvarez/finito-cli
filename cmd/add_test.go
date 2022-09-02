package cmd

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestAddCmd(t *testing.T) {
	cmd := newAddCmd(nil)
	b := bytes.NewBufferString("")

	cmd.command().SetOut(b)
	cmd.command().Execute()

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}
