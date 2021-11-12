package lexer

import (
	"fmt"
	"go/token"
	"unicode"
)

// isSpace reports whether r is a space character.
func isSpace(r int) bool {
	switch r {
	case ' ', '\t', '\r':
		return true
	}
	return false
}

func isAlpha(r int) bool {
	return r == '_' || unicode.IsLetter(rune(r))
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r int) bool {
	return r == '_' || unicode.IsLetter(rune(r)) || unicode.IsDigit(rune(r))
}

func (a Item) equal(b Item) bool {
	if a.Token != b.Token {
		return false
	}
	if a.Literal != b.Literal {
		return false
	}
	if a.Pos != token.NoPos && b.Pos != token.NoPos {
		if a.Pos != b.Pos {
			return false
		}
	}

	return true
}

func (i Item) String() string {
	return fmt.Sprintf("{%v:%q}", i.Token, i.Literal)
}
