package gps

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func AssertScannerReadFromN(got, want int64) error {
	var (
		format = "Scanner.RedFrom, n output: Got: n='%d', Want: n='%d'"
	)
	if want != got {
		return fmt.Errorf(format, got, want)
	}
	return nil
}

func AssertScannerReadFromError(got, want error) error {
	var (
		gStr   string
		wStr   string
		format = "Scanner.RedFrom, error output: Got err='%s', Want: err='%s'"
	)
	gStr = fmt.Sprintf("%s", got)
	wStr = fmt.Sprintf("%s", want)
	if strings.EqualFold(wStr, gStr) == false {
		return fmt.Errorf(format, got, want)
	}
	return nil
}

func AssertScannerReadFromStructFieldIoReader(got, want io.Reader) error {
	format := "%s '%#v', Want struct field: err='%#v'"
	s1 := "Scanner.RedFrom, comparing io.Readers: Got struct field '"

	gStr := fmt.Sprintf("%#v", got)
	wStr := fmt.Sprintf("%#v", want)

	if strings.Compare(wStr, gStr) != 0 {
		return fmt.Errorf(format, s1, got, want)
	}
	return nil
}

func AssertScannerReadFromStructFieldC(got, want rune) error {
	panic("FIXME Llaves færdig")
	return nil
}

func TestScannerReadFromNilIoReader(t *testing.T) {
	var (
		s      = NewScanner(&gst_testing)
		gotN   int64
		gotErr error
		r      = io.Reader(nil)
		err    error
	)
	gotN, gotErr = s.ReadFrom(r)
	if err = AssertScannerReadFromN(gotN, 0); err != nil {
		t.Error(err.Error())
	}
	if err = AssertScannerReadFromError(gotErr, ErrIoNilReader); err != nil {
		t.Error(err.Error())
	}
}

func TestScannerReadFromStructFields(t *testing.T) {
	var (
		s              *Scanner
		bufStructField bytes.Buffer
		bufIoReaderArg bytes.Buffer
		r              io.Reader
		err            error
	)

	s = NewScanner(&gst_testing)

	bufStructField.WriteString(" ")
	s.ioReader = io.Reader(&bufStructField)

	bufIoReaderArg.WriteString("\"class\":\"TPV\"}")
	r = io.Reader(&bufIoReaderArg)

	_, _ = s.ReadFrom(r)

	if err = AssertScannerReadFromStructFieldIoReader(r, r); err != nil {
		t.Error(err.Error())
	}

	if err = AssertScannerReadFromStructFieldC(s.c, s.input[0]); err != nil {
		t.Error(err.Error())
	}
}
