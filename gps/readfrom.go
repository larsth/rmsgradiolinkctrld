package gps

import (
	"errors"
	"io"
)

var (
	ErrGpsdJsonDocumentHadNotBeenParsed = errors.New(
		"ERROR: The gpsd JSON document had not been parsed")
	//'bufferSize' is not a const, because its size is set to a much lower
	//value in unit tests
	bufferSize = 4096
)

//ReadFrom is a function that reads a stream of gpsd JSON documents.
//
//ReadFrom stores some of the GPS data from the following gpsd
//JSON documents:
// POLL (stores only the 1st TPV item, and stores no SkyView items), TPV
//
//ReadFrom logs all recieved gpsd JSON document it recieves.
//Errors are returned, logged, and printed to STDERR.
//
//ReadFrom implements the io.ReadFrom interface.
func (coord *GPSCoord) ReadFrom(r io.Reader) (n int64, err error) {
	coord.mutex.Lock()
	defer coord.mutex.Unlock()

	if len(coord.jsonDocument) > 0 {
		return 0, ErrGpsdJsonDocumentHadNotBeenParsed
	}

	n, coord.jsonDocument, err = coord.scan(bufferSize, r)
	if err != nil {
		return n, err
	}

	return n, nil
}
