package main

import (
	"bufio"
	"fmt"
	"io"
	"monkey/src"
	"os"
	"os/user"
)

func repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print("< ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := src.NewLexer(line)

		for tok := l.NextToken(); tok.Type != src.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Println("Feel free to type in commands")
	repl(os.Stdin, os.Stdout)
}
