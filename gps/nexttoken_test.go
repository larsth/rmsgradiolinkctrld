package gps

//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"io"
//	"strings"
//	"testing"
//)

//func TestGPSCoordScanNilIoReader(t *testing.T) {
//	var (
//		coord *GPSCoord = new(GPSCoord)

//		gotSlice []byte
//		gotErr   error

//		wantSlice []byte = nil
//		wantErr   error  = ErrNilReader

//		gotSliceString  string
//		wantSliceString string

//		gotErrString  string
//		wantErrString string
//	)

//	coord.reader = nil
//	gotSlice, gotErr = coord.Scan()

//	if bytes.Equal(wantSlice, gotSlice) == false {
//		gotSliceString = fmt.Sprintf("%#v", gotSlice)
//		wantSliceString = fmt.Sprintf("%#v", wantSlice)
//		t.Errorf("Want slice: %s, Got slice: %s\n", wantSliceString, gotSliceString)
//	}

//	gotErrString = fmt.Sprintf("%#v", gotErr)
//	wantErrString = fmt.Sprintf("%#v", wantErr)
//	if strings.EqualFold(wantErrString, gotErrString) == false {
//		t.Errorf("Want error: %s, Got error: %s\n", wantErrString, gotErrString)
//	}
//}

//func TestGPSCoordScanNonNilErr(t *testing.T) {
//	var (
//		coord *GPSCoord = new(GPSCoord)

//		buf   *bytes.Buffer
//		input string = ""

//		gotSlice []byte
//		gotErr   error

//		wantSlice []byte = nil
//		wantErr   error  = io.EOF

//		gotSliceString  string
//		wantSliceString string

//		gotErrString  string
//		wantErrString string
//	)

//	buf = bytes.NewBufferString(input)
//	coord.reader = bufio.NewReader(buf)
//	gotSlice, gotErr = coord.Scan()

//	if bytes.Equal(wantSlice, gotSlice) == false {
//		gotSliceString = fmt.Sprintf("%#v", gotSlice)
//		wantSliceString = fmt.Sprintf("%#v", wantSlice)
//		t.Errorf("Want slice: %s, Got slice: %s\n", wantSliceString, gotSliceString)
//	}

//	gotErrString = fmt.Sprintf("%#v", gotErr)
//	wantErrString = fmt.Sprintf("%#v", wantErr)
//	if strings.EqualFold(wantErrString, gotErrString) == false {
//		t.Errorf("Want error: %s, Got error: %s\n", wantErrString, gotErrString)
//	}
//}

//func TestGPSCoordScanIsNotPrefixed(t *testing.T) {
//	var (
//		coord *GPSCoord = new(GPSCoord)

//		buf   *bytes.Buffer
//		input string = `{"class":"TPV"}\r\n`

//		gotReturnedSlice []byte

//		gotErr error

//		wantSlice []byte = nil
//		wantErr   error  = nil

//		gotString  string
//		wantString string
//	)

//	buf = bytes.NewBufferString(input)
//	coord.reader = bufio.NewReader(buf)

//	gotReturnedSlice, gotErr = coord.Scan()

//	if bytes.Equal(wantSlice, gotReturnedSlice) == false {
//		gotString = fmt.Sprintf("%#v", gotReturnedSlice)
//		wantString = fmt.Sprintf("%#v", wantSlice)
//		t.Errorf("Want returned slice: %s, Got returned slice: %s\n", wantString, gotString)
//	}

//	gotString = fmt.Sprintf("%#v", gotErr)
//	wantString = fmt.Sprintf("%#v", wantErr)
//	if strings.EqualFold(wantString, gotString) == false {
//		t.Errorf("Want error: %s, Got error: %s\n", wantString, gotString)
//	}
//}

//func TestGPSCoordScanIsPrefixed(t *testing.T) {
//	var (
//		coord *GPSCoord = new(GPSCoord)

//		buf   *bytes.Buffer
//		input string = `{"class":"TPV"}`

//		gotReturnedSlice  []byte
//		wantReturnedSlice []byte = nil

//		gotErr  error
//		wantErr error = nil

//		gotBufferedSlice  []byte
//		wantBufferedSlice []byte

//		gotString  string
//		wantString string
//	)

//	wantBufferedSlice = []byte(input)

//	buf = bytes.NewBufferString(input)
//	coord.reader = bufio.NewReader(buf)
//	gotReturnedSlice, gotErr = coord.Scan()

//	if bytes.Equal(wantReturnedSlice, gotReturnedSlice) == false {
//		gotString = fmt.Sprintf("%#v", gotReturnedSlice)
//		wantString = fmt.Sprintf("%#v", wantReturnedSlice)
//		t.Errorf("Want returned slice: %s, Got returned slice: %s\n", wantString, gotString)
//	}

//	gotBufferedSlice = coord.scanBuf.Bytes()
//	if bytes.Equal(wantBufferedSlice, gotBufferedSlice) == false {
//		gotString = fmt.Sprintf("%#v", gotBufferedSlice)
//		wantString = fmt.Sprintf("%#v", wantBufferedSlice)
//		t.Errorf("Want buffered slice: %s, Got buffered slice: %s\n", wantString, gotString)
//	}

//	gotString = fmt.Sprintf("%#v", gotErr)
//	wantString = fmt.Sprintf("%#v", wantErr)
//	if strings.EqualFold(wantString, gotString) == false {
//		t.Errorf("Want error: %s, Got error: %s\n", wantString, gotString)
//	}
//}

////func TestGPSCoordScanIsPrefixed(t *testing.T) {
////	var (
////		coord *GPSCoord = new(GPSCoord)

////		buf   *bytes.Buffer
////		input string = `{"class":"TPV"}\r\n`

////		gotReturnedSlice []byte

////		gotErr error

////		want      string = `{"class":"TPV"}`
////		wantSlice []byte = nil
////		wantErr   error  = nil

////		gotBufferedSlice  []byte
////		wantBufferedSlice []byte

////		gotString  string
////		wantString string
////	)

////	wantBufferedSlice = []byte(want)

////	buf = bytes.NewBufferString(input)
////	coord.reader = bufio.NewReader(buf)
////	gotReturnedSlice, gotErr = coord.Scan()

////	if bytes.Equal(wantSlice, gotReturnedSlice) == false {
////		gotString = fmt.Sprintf("%#v", gotReturnedSlice)
////		wantString = fmt.Sprintf("%#v", wantSlice)
////		t.Errorf("Want returned slice: %s, Got returned slice: %s\n", wantString, gotString)
////	}

////	gotBufferedSlice = coord.scanBuf.Bytes()
////	if bytes.Equal(wantBufferedSlice, gotBufferedSlice) == false {
////		gotString = fmt.Sprintf("%#v", gotBufferedSlice)
////		wantString = fmt.Sprintf("%#v", wantBufferedSlice)
////		t.Errorf("Want buffered slice: %s, Got buffered slice: %s\n", wantString, gotString)
////	}

////	gotString = fmt.Sprintf("%#v", gotErr)
////	wantString = fmt.Sprintf("%#v", wantErr)
////	if strings.EqualFold(wantString, gotString) == false {
////		t.Errorf("Want error: %s, Got error: %s\n", wantString, gotString)
////	}
////}

////func TestGPSCoordScanIsPrefixedThenNotPrefixedSequence(t *testing.T) {
////}
