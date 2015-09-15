package gps

import (
	"bufio"
	"errors"
	"io"
)

var (
	ErrBufioReaderNotInitialized = errors.New("The bufio.Reader is not initialized")
	ErrIOReaderNotInitialized    = errors.New("The io.Reader is not initialized")
	ErrBufferSizeNegativeOrZero  = errors.New("The buffer size is negative or zero.")
)

func (coord *GPSCoord) scan(bufferSize int, r io.Reader) (int64, []byte, error) {
	var (
		line     []byte
		isPrefix bool
		err      error
		length   int
		slice    []byte
	)

	if bufferSize <= 0 {
		return 0, nil, ErrBufferSizeNegativeOrZero
	}
	if r == nil {
		return 0, nil, ErrIOReaderNotInitialized
	}

	if r != coord.ioReader {
		coord.ioReader = r
		coord.reader = bufio.NewReaderSize(coord.ioReader, bufferSize)
	}
	if coord.reader == nil {
		return 0, nil, ErrBufioReaderNotInitialized
	}

	line, isPrefix, err = coord.reader.ReadLine()
	if err != nil {
		return 0, nil, err
	}
	if isPrefix == true {
		//This is not the entire line, so add it to the coord.scanBuf bytes.Buffer
		coord.scanBuf.Write(line) //always successful, so error is ignored/unhandled
		return int64(len(line)), nil, nil
	} else {
		//isPrefix == false:

		// deferred coord.scanBuf cleaning
		defer coord.scanBuf.Reset()

		//	line is either the complete(unprefixed) line,
		//	or the last(prefixed) part of a line:
		if coord.scanBuf.Len() > 0 {
			//'line' is the last(prefixed) part of a line:

			length = coord.scanBuf.Len() + len(line)
			slice = make([]byte, length)
			slice = append(slice, coord.scanBuf.Bytes()...)
			slice = append(slice, line...)

			return int64(len(line)), slice, nil
		} else {
			//coord.scanBuf.Len() == 0:
			//	'line' is a complete(unprefixed) line:
			slice = line
			return int64(len(line)), slice, nil
		}
	}
}
