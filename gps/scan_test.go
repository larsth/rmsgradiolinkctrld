package gps

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
)

var tb_scan_table []*TestBenchData_func_scan = []*TestBenchData_func_scan{
	&TestBenchData_func_scan{
		InputBufferSize: int(-1),
		inputString:     "", // -> nil
		Want: CheckDataNSliceErr{
			N:     int64(0),
			Slice: "", /* -> nil */
			Err:   ErrBufferSizeNegativeOrZero,
		},
	},
	&TestBenchData_func_scan{
		InputBufferSize: int(0),
		inputString:     "", // -> nil
		Want: CheckDataNSliceErr{
			N:     int64(0),
			Slice: "", /* -> nil */
			Err:   ErrBufferSizeNegativeOrZero,
		},
	},
	&TestBenchData_func_scan{
		InputBufferSize: bufferSize,
		inputString:     "{\"class\":\"TPV\"}\r\n",
		Want: CheckDataNSliceErr{
			N:     int64(15),
			Slice: `{"class":"TPV"}`,
			Err:   nil,
		},
	},
	&TestBenchData_func_scan{
		InputBufferSize: bufferSize,
		inputString:     "\r\n",
		Want: CheckDataNSliceErr{
			N:     int64(0),
			Slice: "", /* -> nil */
			Err:   nil,
		},
	},
}

type CheckDataNSliceErr struct {
	N     int64
	Slice string
	Err   error
}

func (cd *CheckDataNSliceErr) SliceBytes() []byte {
	if len(cd.Slice) == 0 {
		return nil
	}
	return []byte(cd.Slice)
}

func (cd *CheckDataNSliceErr) SetSliceBytes(b []byte) {
	if len(b) == 0 {
		cd.Slice = ""
		return
	}
	cd.Slice = string(b)
}

type TestBenchData_func_scan struct {
	inputString     string
	InputBufferSize int
	Want            CheckDataNSliceErr
}

func (tbd *TestBenchData_func_scan) InputSlice() []byte {
	if len(tbd.inputString) == 0 {
		return nil
	}
	return []byte(tbd.inputString)
}

func (tbd *TestBenchData_func_scan) String() string {
	var (
		buf bytes.Buffer
	)

	buf.WriteString("\tTest data for testing a *GPSCoord.scan method:\n\n")

	buf.WriteString("\t\tInputBufferSize: `")
	buf.WriteString(strconv.Itoa(tbd.InputBufferSize))
	buf.WriteString("`\n")

	buf.WriteString("\t\tInputSlice: `")
	if tbd.InputSlice != nil {
		buf.Write(tbd.InputSlice())
		buf.WriteString("`\n")
	} else {
		buf.WriteString("<nil>`\n")
	}

	buf.WriteString("\t\tWantN: `")
	buf.WriteString(strconv.FormatInt(tbd.Want.N, 10))
	buf.WriteString("`\n")

	buf.WriteString("\t\tWantSlice: `")
	the_slice := tbd.Want.SliceBytes()
	if len(the_slice) > 0 {
		buf.Write(the_slice)
		buf.WriteString("`\n")
	} else {
		buf.WriteString("<nil>`\n")
	}

	buf.WriteString("\t\tWantErr: `")
	if tbd.Want.Err != nil {
		buf.WriteString(tbd.Want.Err.Error())
		buf.WriteString("`\n")
	} else {
		buf.WriteString("<nil>`\n")
	}
	buf.WriteString("\n\n")

	return buf.String()
}

func checkGpsCoordScanN(gotN, wantN int64) (msg string) {
	//test gotN vs test_item.WantN:
	if wantN != gotN {
		msg = fmt.Sprintf("\tGot N:=%d, Want N:=%d\n", gotN, wantN)
		return
	}

	msg = ""
	return
}

func checkGpsCoordScanSlice(gotSlice, wantSlice []byte) (msg string) {
	var (
		s_got  string
		s_want string
	)

	//test gotSlice vs test_item.WantSlice:
	if bytes.Equal(wantSlice, gotSlice) == false {
		if len(gotSlice) > 0 {
			s_got = string(gotSlice)
		} else {
			s_got = "<nil>"
		}

		if len(wantSlice) > 0 {
			s_want = string(wantSlice)
		} else {
			s_want = "<nil>"
		}

		msg = fmt.Sprintf("\tGot slice:=`%s`. Want slice:=`%s`", s_got, s_want)
		return
	}

	msg = ""
	return
}

func checkGpsCoordScanErr(gotErr, wantErr error) (msg string) {
	var (
		s_got  string
		s_want string
	)

	//test gotErr vs test_item.WantErr:
	if gotErr != nil {
		s_got = gotErr.Error()
	} else {
		s_got = "<nil>"
	}

	if wantErr != nil {
		s_want = wantErr.Error()
	} else {
		s_want = "<nil>"
	}

	if strings.Compare(s_want, s_got) != 0 {
		msg = fmt.Sprintf("\tGot error:=`%s`. Want error:=`%s`\n", s_got, s_want)
		return
	} else {
		msg = ""
		return
	}
}

func checkGPScoordscan(got *CheckDataNSliceErr, want *CheckDataNSliceErr) [3]string {
	var strs [3]string

	//test gotN vs test_item.WantN:
	strs[0] = checkGpsCoordScanN(got.N, want.N)

	//test gotSlice vs test_item.WantSlice:
	strs[1] = checkGpsCoordScanSlice(got.SliceBytes(), want.SliceBytes())

	//test gotErr vs test_item.WantErr:
	strs[2] = checkGpsCoordScanErr(got.Err, want.Err)

	return strs
}

func hasFailedGpsCoordscan(strs [3]string) bool {
	var n int

	for _, s := range strs {
		if len(s) > 0 {
			n++
		}
	}

	return n > 0
}

//func TestGpsCoordScanWithoutPrefixFollowedByWithoutPrefix(t *testing.T) {
//	var (
//		buf      bytes.Buffer
//		ioReader io.Reader
//	)

//}

//TestGpsCoordScanWithPrefixFollowedByWithoutPrefix is a test function what 1st
//exercices the *GPScoord.scan method in such a way that it returns a
//isPrefix=true, and then again but in such a way that it returns a
//isPrefix=false.
//This test is done with a very low bufferSize (10).
func TestGpsCoordScanWithPrefixFollowedByWithoutPrefix(t *testing.T) {
	var (
		jsonDoc   []byte        = []byte("{\"class\":\"TPV\"}\r\n")
		reader    *bytes.Reader = bytes.NewReader(jsonDoc)
		coord     *GPSCoord     = new(GPSCoord)
		got       CheckDataNSliceErr
		the_slice []byte
		strs      [3]string

		test_item_wo_prefix *TestBenchData_func_scan = &TestBenchData_func_scan{
			InputBufferSize: 10,
			inputString:     "{\"class\":\"TPV\"}\r\n",
			Want: CheckDataNSliceErr{
				N:     int64(10),
				Slice: "{\"class\":\"",
				Err:   nil,
			},
		}
		//		test_item_w_prefix *TestBenchData_func_scan = &TestBenchData_func_scan{
		//			InputBufferSize: 10,
		//			inputString:     "{\"class\":\"TPV\"}\r\n",
		//			Want: CheckDataNSliceErr{
		//				N:     int64(5),
		//				Slice: "TPV\"}",
		//				Err:   nil,
		//			},
		//		}
	)

	//========================================================================
	//isPrefix == true test

	got.N, the_slice, got.Err = coord.scan(10, reader)
	got.SetSliceBytes(the_slice)
	strs = checkGPScoordscan(&got, &(test_item_wo_prefix.Want))
	//if (len(strs[0]) > 0) || (len(strs[1]) > 0) || (len(strs[2]) > 0) {
	if hasFailedGpsCoordscan(strs) == true {
		t.Fail()
		for _, str := range strs {
			if len(str) > 0 {
				t.Log(str)
			}
		}
		//Write the test data that triggered a fail.
		t.Errorf("%s", test_item_wo_prefix.String())
	}

	//	//========================================================================
	//	////isPrefix == false test
	//	wantN = 5
	//	wantSlice = []byte{"TPV\"}"}
	//	wantErr = nil

	//	got.SetSliceBytes(the_slice)
	//	strs = checkGPScoordscan(got, test_item_w_prefix.Want)
	////if (len(strs[0]) > 0) || (len(strs[1]) > 0) || (len(strs[2]) > 0) {
	//if hasFailedGpsCoordscan(strs) == true {
	//t.Fail()
	//			for _, str := range strs {
	//				if len(str) > 0 {
	//					t.Log(str)
	//				}
	//			}
	//			//Write the test data that triggered a fail.
	//			t.Errorf("%s", test_item.String())
	//	}
}

func TestGpsCoordScanNilReader(t *testing.T) {
	const wantN = 0
	var (
		coord     GPSCoord
		test_item *TestBenchData_func_scan = &TestBenchData_func_scan{
			InputBufferSize: bufferSize,
			Want: CheckDataNSliceErr{
				N:     int64(0),
				Slice: "",
				Err:   ErrIOReaderNotInitialized,
			},
		}
		got       CheckDataNSliceErr
		the_slice []byte
		strs      [3]string
	)

	//Special test case: The io.Reader provided to the scan function is nil.
	//	test_item.InputSlice == []byte{'_'}, so:
	got.N, the_slice, got.Err = coord.scan(test_item.InputBufferSize, io.Reader(nil))
	got.SetSliceBytes(the_slice)
	strs = checkGPScoordscan(&got, &(test_item.Want))
	//if (len(strs[0]) > 0) || (len(strs[1]) > 0) || (len(strs[2]) > 0) {
	if hasFailedGpsCoordscan(strs) == true {
		t.Fail()
		for _, str := range strs {
			if len(str) > 0 {
				t.Log(str)
			}
		}
		//Write the test data that triggered a fail.
		t.Errorf("%s", test_item.String())
	}
}

func TestGpsCoordScan(t *testing.T) {
	var (
		reader *bytes.Reader
		coord  GPSCoord

		got       CheckDataNSliceErr
		the_slice []byte
		strs      [3]string
	)

	for _, test_item := range tb_scan_table {
		//run function being tested:
		reader = bytes.NewReader(test_item.InputSlice())
		got.N, the_slice, got.Err = coord.scan(test_item.InputBufferSize, reader)
		got.SetSliceBytes(the_slice)
		strs = checkGPScoordscan(&got, &(test_item.Want))
		//if (len(strs[0]) > 0) || (len(strs[1]) > 0) || (len(strs[2]) > 0) {
		if hasFailedGpsCoordscan(strs) == true {
			t.Fail()
			for _, str := range strs {
				if len(str) > 0 {
					t.Log(str)
				}
			}
			//Write the test data that triggered a fail.
			t.Errorf("%s", test_item.String())
		}
	}
}
