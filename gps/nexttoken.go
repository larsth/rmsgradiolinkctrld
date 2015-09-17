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

func (s *Scanner) LINE(t *Token) {
	//	for len(s.input) > 0 {
	//		if isLetter(s.c) {
	//			s.p++
	//		} else {
	//			break
	//		}
	//	}
	//	t.Ident = TokenIdentLine
	//	if s.p == 0 {
	//		t.Error = io.EOF
	//		t.Runes = nil
	//	} else {
	//		t.Error = nil
	//		out := s.input[0:s.p]
	//		t.Runes = make([]rune, 0, len(out))
	//		copy(t.Runes, out)
	//		//reslice s.input:
	//		s.input = append(s.input[p:]...)
	//		s.p = 0
	//	}
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
			s.LINE(&t)
		} else {
			t.Error = fmt.Errorf("Unexpected character: '%v'\n\n", s.c)
			t.Ident = TokenIdentError
			t.Runes = nil
		}
	}
	return
}

//func (s *Scanner) readFromScanBuffer(t *Token) {
//	const i int = 0
//	var (
//		j         int  = 0
//		doReslice bool = false
//	)
//	if len(s.inputBuffer) == 0 {
//		t.Error = io.EOF
//		t.Ident = TokenIdentIoEOF
//	} else {
//		for j < len(s.inputBuffer) {
//			doReslice = true
//			if s.inputBuffer[j] == '\r' || s.inputBuffer[j] == '\n' {
//				if t.Ident == TokenIdentLine {
//					break
//				}
//				if s.inputBuffer[j] == '\r' {
//					t.Ident = TokenIdentSlashR
//					t.Value = []byte{'\r'}
//				}
//				if s.inputBuffer[j] == '\n' {
//					t.Ident = TokenIdentSlashN
//					t.Value = []byte{'\n'}
//				}
//			} else {
//				//add to ouput buffer:
//				s.outputbuffer = append(s.outputbuffer, s.inputBuffer[j])
//			}
//		}
//		if doReslice == true {
//			//cut bytes from index i (included) to index j (excluded):
//			s.inputBuffer = append(s.inputBuffer[:i], s.inputBuffer[j:]...)
//		}
//	}
//}

//func (s *Scanner) Scan1() (t Token) {
//	var (
//		p   []byte
//		n   int
//		err error
//	)
//	s.mutex.Lock()
//	defer s.mutex.Unlock()

//	if s.ioReader == nil {
//		t.Error = ErrIoNilReader
//		t.Ident = TokenIdentError
//		return
//	}

//	//1st: Read from the buffer:
//	s.readFromScanBuffer(&t)
//	if t.Ident != TokenIdentError && t.Ident != TokenIdentIoEOF &&
//		t.Ident != TokenIdentWasPrefixed {
//		//TokenIdentLine, or TokenIdentSlashN or TokenIdentSlashR
//		//was recieved from the buffer:
//		return
//	}

//	//2nd: Read from the s.ioReader:
//	p = make([]byte, 4096)
//	n, err = s.ioReader.Read(p)

//	if err != io.EOF {
//		t.Error = io.EOF
//		t.Ident = TokenIdentIoEOF
//		return
//	}
//	if err != nil {
//		t.Error = err
//		t.Ident = TokenIdentError
//		return
//	}

//	//n > 0:
//	s.scanBuf = append(s.scanBuf, p[0:n]...)

//	//Below:
//	//This is an invitation to call Scan again:
//	t.Ident = TokenIdentWasPrefixed

//	return
//}

//func (coord *GPSCoord) Scan0() (isPrefixed bool, token []byte, err error) {
//	var (
//		p           []byte = make([]byte, 256)
//		n           int
//		indexSlashN int
//		indexSlashR int
//		jsonDocN    []byte
//		jsonDoc     []byte
//	)

//
//		indexSlashN = bytes.IndexByte(coord.scanBuf, '\n')
//		if indexSlashN != -1 {
//			//coord.scanBuf has a '\n' at 'indexSlashN':
//			isPrefixed = false
//			jsonDocN = coord.scanBuf[0:indexSlashN]

//			indexSlashR = bytes.IndexByte(jsonDocN, '\r')
//			if indexSlashR != -1 {
//				//jsonDocN has a '\r' at 'indexSlashR':
//				jsonDoc = jsonDocN[0:indexSlashR]
//			} else {
//				//jsonDocN does not have a '\r' at 'indexSlashR':
//				jsonDoc = jsonDocN
//			}

//			token = make([]byte, len(jsonDoc))
//			copy(token, jsonDoc)
//			p = nil
//			jsonDocN = nil
//			jsonDoc = nil
//			//nuke the coord.scanBuf buffer:
//			coord.scanBuf = make([]byte, DefaultBufferSize)
//		} else {
//			isPrefixed = true
//			//Content had already been appended.
//		}
//		return isPrefixed, nil, nil //no token, and no error
//	}
//	//else: n == 0 is true:
//	return false, nil, nil //is not prefixed, no token, and no error
//}
