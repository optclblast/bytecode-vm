package lexer

type Lexer struct {
	p      int
	read_p int
	ch     rune
	chars  []rune
}

// New creates a new instace of a Lexer from input string that will be processed
func New(in string) *Lexer {
	return &Lexer{chars: []rune(in)}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespaces()

}

// Reads and stores next char
func (l *Lexer) readChar() {
	if l.read_p >= len(l.chars) {
		l.ch = rune(0)
	} else {
		l.ch = l.chars[l.read_p]
	}

	l.p = l.read_p
	l.read_p++
}

func (l *Lexer) skipWhitespaces() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
}
