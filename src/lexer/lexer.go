package lexer

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
	EOF                  = 8
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

func Tokenize(source string) ([]Token, error) {
	var tokens []Token
	src := strings.Split(source, "")
	for i := 0; i < len(src); i++ {
		if src[i] == "(" {
			tokens = append(tokens, Token{src[i], LeftParen})
		} else if src[i] == ")" {
			tokens = append(tokens, Token{src[i], RightParen})
		} else if src[i] == "+" || src[i] == "-" || src[i] == "*" || src[i] == "/" {
			tokens = append(tokens, Token{src[i], BinaryOp})
		} else if src[i] == "=" {
			tokens = append(tokens, Token{src[i], Equals})
		} else {
			//For multicharacter tokens,
			if isNum(src[i]) {
				num := ""
				for i < len(src) && isNum(src[i]) {
					num += src[i]
					i++
				}
				tokens = append(tokens, Token{num, Number})
			} else if isAlpha(src[i]) {
				ident := ""
				for i < len(src) && isAlpha(src[i]) {
					ident += src[i]
					i++
				}

				//reserved keywords
				reserved := Keywords[ident]
				if reserved == 0 {
					tokens = append(tokens, Token{ident, Identifier})
				} else {
					tokens = append(tokens, Token{ident, Keywords[ident]})
				}
			} else if isSkippable(src[i]) {
				continue
			} else {
				return nil, fmt.Errorf("Unknown character found: %s", src[i])
			}
		}

	}
	tokens = append(tokens, Token{"EOF", EOF})
	return tokens, nil
}
