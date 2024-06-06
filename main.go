package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bousii/babbelas/parser"
)

func main() {
	parser := parser.Parser{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" || text == "" {
			os.Exit(1)
		}
		program, err := parser.ProduceAST(text)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, stmt := range program.Body {
			fmt.Printf("%s : %v\n", stmt.GetKind(), stmt)
		}
	}
}
