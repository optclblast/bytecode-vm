package lexer

type Token struct {
	Type    TType
	Literal string
}

type TType string

// todo types id
