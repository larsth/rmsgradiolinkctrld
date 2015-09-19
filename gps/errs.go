package gps

import "errors"

//var (
////	ErrGpsdJsonDocumentHadNotBeenParsed = errors.New(
////		"ERROR: The gpsd JSON document had not been parsed")
////	//'bufferSize' is not a const, because its size is set to a much lower
////	//value in unit tests
////ErrIoNilReader           = errors.New("Error: the io.Reader is nil")
////ErrBufferSizeLessThanOne = errors.New("The buffer size is less than 1")

////ErrNilBufioReader is a complaint about a *bufio.Reader is nil
////ErrNilBufioReader = errors.New("Error: The buffered I/O Reader is nil")
//)

var (
	ErrSameLocation                  error = errors.New("The 2 GPS coordinates are the same location.")
	ErrOtherGPSCoordIsNil            error = errors.New("The `other` *GPSCoord pointer is nil.")
	ErrThisGPSCoordHasZeroTimestamp  error = errors.New("The `this` TimeStamp is Zero: 0")
	ErrOtherGPSCoordHasZeroTimestamp error = errors.New("The `other`TimeStamp is Zero: 0")
	ErrNoPollFixes                         = errors.New("No POLL fixes (No TPV JSON documents)")
)
