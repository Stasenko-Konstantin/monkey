package src

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func newLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) nextToken() token {
	var tok token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case 0:
		tok = newToken(EOF, "")
	default:
		if isLetter(l.ch) {
			return l.readIdentifier()
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken[T byte | string](ttype tokenType, ch T) token {
	return token{ttype: ttype, literal: string(ch)}
}

func (l *Lexer) readIdentifier() token {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	tok := l.input[pos:l.position]
	return newToken(lookupIdent(tok), tok)
}

func (l *Lexer) readNumber() token {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	tok := l.input[pos:l.position]
	return newToken(INT, tok)
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}