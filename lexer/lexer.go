package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := Lexer{
		input: input,
	}

	l.readChar()

	return &l
}

func (l *Lexer) readChar() {
	// If I read past the input
	if l.readPosition >= len(l.input) {
		//set ch to EOF
		l.ch = 0
	} else {
		// update character
		l.ch = l.input[l.readPosition]
	}

	// update current position to match ch
	l.position = l.readPosition

	//  move to next position
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.SetStr(token.EQ, string(ch)+string(l.ch))
		} else {
			tok.SetCh(token.ASSIGN, l.ch)
		}
	case '+':
		tok.SetCh(token.PLUS, l.ch)
	case '-':
		tok.SetCh(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.SetStr(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok.SetCh(token.BANG, l.ch)
		}
	case '*':
		tok.SetCh(token.ASTERISK, l.ch)
	case '/':
		tok.SetCh(token.SLASH, l.ch)
	case '<':
		tok.SetCh(token.LT, l.ch)
	case '>':
		tok.SetCh(token.GT, l.ch)
	case ';':
		tok.SetCh(token.SEMICOLON, l.ch)
	case ',':
		tok.SetCh(token.COMMA, l.ch)
	case '(':
		tok.SetCh(token.LPAREN, l.ch)
	case ')':
		tok.SetCh(token.RPAREN, l.ch)
	case '{':
		tok.SetCh(token.LBRACE, l.ch)
	case '}':
		tok.SetCh(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			ident := l.readIndentifier()
			return token.CreateStr(token.LookupIdent(ident), ident)
		} else if isDigit(l.ch) {
			return token.CreateStr(token.INT, l.readNumber())
		} else {
			tok.SetCh(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// recognizes identifier and reads it
func (l *Lexer) readIndentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// recognizes number and reads it
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Look at next char without reading forward
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
