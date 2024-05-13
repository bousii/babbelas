package main

//arrays are indexed at -1
import (
	"fmt"
	"regexp"
	"strings"
)

type TokenType int

const (
	Identifier TokenType = 1
	Bou                  = 2 //initializer token, similar to let for javascript
	Equals               = 3
	LeftParen            = 4
	RightParen           = 5
	BinaryOp             = 6 //doing basic math, remember we are going to flip these hehehehehe
	Number               = 7
)

var Keywords = map[string]TokenType{
	"Bou": Bou, //replace with eventual final variable identifier keyword
}

type Token struct {
	Value     string
	tokentype TokenType
}

func (t Token) display() {
	fmt.Printf("{ Value ->  %s, Type -> %d }\n", t.Value, t.tokentype)
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

func isSkippable(value string) bool {
	pattern := "\\s"
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func tokenize(source string) ([]Token, error) {
	var tokens []Token
	src := strings.Split(source, "")
	for len(src) > 0 {
		if src[0] == "(" {
			tokens = append(tokens, Token{src[0], LeftParen})
		} else if src[0] == ")" {
			tokens = append(tokens, Token{src[0], RightParen})
		} else if src[0] == "+" || src[0] == "-" || src[0] == "*" || src[0] == "/" {
			tokens = append(tokens, Token{src[0], BinaryOp})
		} else if src[0] == "=" {
			tokens = append(tokens, Token{src[0], Equals})
		} else {
			//For multicharacter tokens, INDEX OUT OF BOUNDS FOR CASE "400"
			if isNum(src[0]) {
				num := ""
				for len(src) > 0 && isNum(src[0]) {
					num += src[0]
					src = src[1:]
				}
				tokens = append(tokens, Token{num, Number})
			} else if isAlpha(src[0]) {
				ident := ""
				for len(src) > 0 && isAlpha(src[0]) {
					ident += src[0]
					src = src[1:]
				}

				//reserved keywords
				reserved := Keywords[ident]
				if reserved == 0 {
					tokens = append(tokens, Token{ident, Identifier})
				} else {
					tokens = append(tokens, Token{ident, Keywords[ident]})
				}
			} else if isSkippable(src[0]) {
				src = src[1:]
			} else {
				return nil, fmt.Errorf("Unknown character found: %s", src[0])
			}
		}

		src = src[1:] //deleting first index, moving forward
	}

	return tokens, nil
}

func main() {
	tokens, err := tokenize("400")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, token := range tokens {
		token.display()
	}
}
