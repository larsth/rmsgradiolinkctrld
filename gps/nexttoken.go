package gps

import (
	"fmt"
	"io"
	"unicode/utf8"
)

func isWS(c rune) bool {
	if c == ' ' || c == '\t' {
		return true
	}
	return false
}

func isLetter(c rune) bool {
	if c == '\r' || c == '\n' {
		return false
	}
	return !isWS(c)
}

//Move one character (reslice):
func (s *Scanner) consume() {
	var copyOfInput []rune

	if len(s.input) <= 1 {
		s.c = utf8.RuneError

		//Give s.input away:
		s.ggc.GiveRunes <- s.input
		s.input = nil // nil <- s.input[:0]
	} else {
		//len(s.input) > 1 is true:

		//Set the current character to the next character in input:
		s.c = s.input[1]

		//Consume the current character:
		copyOfInput = <-s.ggc.GetRunes
		copyOfInput = append(copyOfInput, s.input[1:]...)

		//Give s.input away, and let copyOfInput be input:
		s.ggc.GiveRunes <- s.input
		s.input = copyOfInput
		copyOfInput = nil
	}
}

//matchLINE matches a LINE token input
func (s *Scanner) matchLINE(t *Token) {
	for len(s.input) > 0 {
		if isLetter(s.c) {
			if len(t.Runes) == 0 {
				t.Runes = <-s.ggc.GetRunes
			}
			t.Runes = append(t.Runes, s.c)
			s.consume()
		} else {
			break
		}
	}

	if len(t.Runes) > 0 {
		t.Ident = TokenIdentLine
		t.Error = nil
		return
	}
	//else:
	return
}

func (s *Scanner) NextToken() (t Token) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.input) == 0 {
		t.Runes = append(t.Runes, utf8.RuneError)
		t.Error = io.EOF
		return
	}
	//else:
	switch s.c {
	case '\r':
		t.Error = nil
		t.Ident = TokenIdentSlashR
		t.Runes = nil
		s.consume()
	case '\n':
		t.Error = nil
		t.Ident = TokenIdentSlashN
		t.Runes = nil
		s.consume()
	default:
		if isLetter(s.c) {
			s.matchLINE(&t)
		} else {
			t.Error = fmt.Errorf("Unexpected character: '%v'\n\n", s.c)
			t.Ident = TokenIdentError
			t.Runes = nil
		}
	}
	return
}
