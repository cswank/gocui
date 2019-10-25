package gocui

import (
	"fmt"
	"testing"
)

func TestEscape(t *testing.T) {
	input := "\033[48;5;200;38;5;100mHi!!\u001b[0m"
	ei := newEscapeInterpreter(Output256)

	for _, r := range input {
		_, err := ei.parseOne(r)
		if err != nil {
			t.Fatal(err)
		}
	}

	r := ei.runes()
	fmt.Println("out:", string(r))
}
