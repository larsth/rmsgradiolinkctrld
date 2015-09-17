package gps

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func AssertScanner_ReadFrom_N(got, want int64) error {
	var (
		format string = "Scanner.RedFrom, n output: Got: n='%d', Want: n='%d'"
	)
	if want != got {
		return fmt.Errorf(format, got, want)
	}
	return nil
}

func AssertScanner_ReadFrom_Error(got, want error) error {
	var (
		gStr   string
		wStr   string
		format string = "Scanner.RedFrom, error output: Got err='%s', Want: err='%s'"
	)
	gStr = fmt.Sprintf("%s", got)
	wStr = fmt.Sprintf("%s", want)
	if strings.EqualFold(wStr, gStr) == false {
		return fmt.Errorf(format, got, want)
	}
	return nil
}

func AssertScanner_ReadFrom_StructFieldIoReader(got, want io.Reader) error {
	format := "%s '%#v', Want struct field: err='%#v'"
	s1 := "Scanner.RedFrom, comparing io.Readers: Got struct field '"

	gStr := fmt.Sprintf("%#v", got)
	wStr := fmt.Sprintf("%#v", want)

	if strings.Compare(wStr, gStr) != 0 {
		return fmt.Errorf(format, s1, got, want)
	}
	return nil
}

func AssertScanner_ReadFrom_StructFieldC(got, want rune) error {
	panic("FIXME Llaves færdig")
	return nil
}

func TestScanner_ReadFromNilIoReader(t *testing.T) {
	var (
		s      *Scanner = NewScanner(&gst_testing)
		gotN   int64
		gotErr error
		r      io.Reader = io.Reader(nil)
		err    error
	)
	gotN, gotErr = s.ReadFrom(r)
	if err = AssertScanner_ReadFrom_N(gotN, 0); err != nil {
		t.Error(err.Error())
	}
	if err = AssertScanner_ReadFrom_Error(gotErr, ErrIoNilReader); err != nil {
		t.Error(err.Error())
	}
}

func TestScanner_ReadFromIoReader(t *testing.T) {
	var (
		s                *Scanner
		buf_struct_field bytes.Buffer
		buf_ioReader_arg bytes.Buffer
		r                io.Reader
		err              error
	)

	s = NewScanner(&gst_testing)

	buf_struct_field.WriteString(" ")
	s.ioReader = io.Reader(&buf_struct_field)

	buf_ioReader_arg.WriteString("\"class\":\"TPV\"}")
	r = io.Reader(&buf_ioReader_arg)

	_, _ = s.ReadFrom(r)

	if err = AssertScanner_ReadFrom_StructFieldIoReader(r, r); err != nil {
		t.Error(err.Error())
	}

	if err = AssertScanner_ReadFrom_StructFieldC(s.c, s.input[0]); err != nil {
		t.Error(err.Error())
	}
}
