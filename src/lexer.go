package main

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenType int

const (
	Identifier TokenType = iota
	Bou                  //initializer token, similar to let for javascript
	Equals
	LeftParen
	RightParen
	BinaryOp //doing basic math, remember we are going to flip these hehehehehe
	Number
)

type Token struct { //
	Value     string
	tokentype TokenType
}

func (t Token) display() {
	fmt.Sprintf("{ Value ->  %s, Type -> %d }", t.Value, t.tokentype)
}

func isAlpha(value string) bool {
	pattern := "^[A-Za-z]+$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func isNum(value string) bool {
	pattern := "^[0-9]+$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func tokenize(source string) []Token {
	var tokens []Token
	src := strings.Split(source, "")
	for len(src) > 0 {
		var tempToken Token
		if src[0] == "(" {
			tempToken = Token{src[0], LeftParen}
		} else if src[0] == ")" {
			tempToken = Token{src[0], RightParen}
		} else if src[0] == "+" || src[0] == "-" || src[0] == "*" || src[0] == "/" {
			tempToken = Token{src[0], BinaryOp}
		} else if src[0] == "=" {
			tempToken = Token{src[0], Equals}
		} else {
			//For multicharacter tokens
			if isNum(src[0]) {
				num := ""
				for len(src) > 0 && isNum(src[0]) {
					num += src[0]
					src = src[1:]
				}
				tokens = append(tokens, Token{src[0], Number})
			} else if isAlpha(src[0]) {
				ident := ""
				for len(src) > 0 && isAlpha(src[0]) {
					ident += src[0]
					src = src[1:]
				}
				tokens = append(tokens, Token{src[0], Identifier})
			}
		}

		tokens = append(tokens, tempToken)
		src = src[1:] //deleting first index, moving forward
	}

	return tokens
}
