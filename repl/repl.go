package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/maybe-joe/monkey/token"
)

const prompt = ">> "

func Run(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			break
		}

		for _, t := range token.NewTokenizer(scanner.Text()).Tokenize() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}

	return nil
}
