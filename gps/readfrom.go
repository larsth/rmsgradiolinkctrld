package gps

import (
	"bytes"
	"io"
)

//ReadFrom implements the io.ReaderFrom interface.
//
//ReadFrom reads data from r until EOF or error.
//The return value n is the number of bytes read.
//Any error - except io.EOF - encountered during the
//read is also returned.
//
//The io.Copy function uses an io.ReaderFrom interface, if available.
//
func (s *Scanner) ReadFrom(r io.Reader) (n int64, err error) {
	var p []byte
	var nInt int
	var pRunes []rune

	//setup
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if r == nil {
		return 0, ErrIoNilReader
	}
	if r != s.ioReader {
		s.ioReader = r
	}
	p = <-s.ggc.GetBytes
	defer func(giveBytesChan chan []byte, p []byte) {
		giveBytesChan <- p
	}(s.ggc.GiveBytes, p)

	//reading
	nInt, err = s.ioReader.Read(p)
	n = int64(nInt)
	if n > 0 {
		pRunes = bytes.Runes(p[0:n])
		if s.input == nil {
			s.input = <-s.ggc.GetRunes
		}
		s.input = append(s.input, pRunes...)
		if len(s.input) > 0 {
			//prime s.c with the 1st character:
			s.c = s.input[0]
		}
	}
	if err == io.EOF {
		//don't return io.EOF:
		//	replace err=io.EOF with err=nil:
		return n, nil
	}
	//else:
	//n and err was set by s.ioReader.Read(p), where err != io.EOF
	return
}
