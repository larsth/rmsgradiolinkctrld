package util

import (
	"bytes"
	"bufio"
)

func ReadLine(reader *bufio.Reader) ([]byte, error) {
	var (
		buf      bytes.Buffer
		line     []byte
		isPrefix bool
		err      error
	)

	//setup:
	if reader == nil {
		return nil, ErrNilBufioReader
	}
	//	p = <-s.ggc.GetBytes
	//	defer func(giveBytesChan chan []byte, p []byte) {
	//		giveBytesChan <- p
	//	}(s.ggc.GiveBytes, p)

	//reading, 1st time:
	line, isPrefix, err = s.reader.ReadLine()
	//A bufio.Reader.ReadLine() returns:
	//		[]byte != nil where len => 0, isPrefix=true|false, err == nil
	//	or
	//		[]byte == nil, isPrefix=true|false, err != nil

	//processing 1st ReadLine:
	if err != nil {
		//if err != nil, then line == nil is true, and isPrefix isn't relevant:
		return nil, err
	}
	// err == nil is true, so line != nil is true, and may have a length > 0:
	if len(line) > 0 {
		//'line' has a length > 0, so buffer 'line':
		buf.Write(line)
	}

	//If the line is prefixed, then we have an incomplete line, so read
	//until the line is read completely (is not prefixed):
	for isPrefix == true {
		line, isPrefix, err = s.reader.ReadLine()
		if err != nil {
			break
		}
		// err == nil is true, so line != nil is true, and may have a length > 0:
		if len(line) > 0 {
			//'line' has a length > 0, so buffer 'line':
			buf.Write(line)
		}
	}

	//returning:
	if buf.Len() > 0 && err == nil {
		return buf.Bytes(), nil
	}
	if buf.Len() > 0 && err != nil {
		return buf.Bytes(), err
	}
	//buf.len == 0 is true:
	if err != nil {
		return nil, err
	}
	//Below:
	//Make the compiler happy: The following source code line is unreachable,
	//because a bufio.Reader.ReadLine never return both a nil 'line' byte slice,
	//and an 'nil' error.
	//This means coverage for this function will never be 100%.
	return nil, nil
}
