package gps

//import (
//	"bytes"
//	"fmt"
//	"strings"
//	"testing"

//	"github.com/larsth/rmsgradiolinkctrld/errs"
//// "github.com/larsth/rmsgradiolinkctrld/logging"
//)

//type TestBenchData_func_parseGpsdJSON struct {
//	Class          string
//	Payload        []byte
//	WantData       *Data
//	WantErr        error
//	WantErrVarName string
//	HasFunc        bool
//	Func           func(test_item *TestBenchData_func_parseGpsdJSON) error
//}

//func (tbd *TestBenchData_func_parseGpsdJSON) String() string {
//	var buf bytes.Buffer

//	buf.WriteString("\tClass: ")
//	buf.WriteString(tbd.Class)
//	buf.WriteString("\n")

//	if len(tbd.Payload) > 0 {
//		buf.WriteString("\tPayload: ")
//		buf.Write(tbd.Payload)
//		buf.WriteString("\n")
//	} else {
//		buf.WriteString("\tPayload: <nil>\n")
//	}

//	if tbd.WantData != nil {
//		buf.WriteString("\tWantData: ")
//		buf.WriteString(tbd.WantData.String())
//		buf.WriteString("\n")
//	} else {
//		buf.WriteString("\tWantData: <nil>\n")
//	}

//	if tbd.WantErr != nil {
//		buf.WriteString("\tWantErr: ")
//		buf.WriteString(tbd.WantErr.Error())
//		buf.WriteString("\n")
//	} else {
//		buf.WriteString("\tWantErr: <nil>\n")
//	}

//	buf.WriteString("\tWantErrVarName: ")
//	buf.WriteString(tbd.WantErrVarName)
//	buf.WriteString("\n")

//	buf.WriteString("\tHasFunc: ")
//	buf.WriteString(fmt.Sprintf("%v", tbd.HasFunc))
//	buf.WriteString("\n")

//	buf.WriteString("\tFunc: ")
//	buf.WriteString(fmt.Sprintf("%#v", tbd.Func))
//	buf.WriteString("\n")

//	return buf.String()
//}

//var tb_parseGpsdJSON_table []*TestBenchData_func_parseGpsdJSON = []*TestBenchData_func_parseGpsdJSON{
//	//len(payload) == 0 test data:
//	&TestBenchData_func_parseGpsdJSON{
//		Class:          "",
//		Payload:        nil,
//		WantData:       nil,
//		WantErr:        errs.ErrZeroLengthPayload,
//		WantErrVarName: `errs.ErrZeroLengthPayload`,
//		HasFunc:        false,
//		Func:           nil,
//	},
//	//len(payload) != 0 && //len(class) == 0 test data:
//	&TestBenchData_func_parseGpsdJSON{
//		Class:          "",
//		Payload:        []byte(`{"class":"TPV"}`),
//		WantData:       nil,
//		WantErr:        errs.ErrEmptyString,
//		WantErrVarName: `errs.ErrEmptyString`,
//		HasFunc:        false,
//		Func:           nil,
//	},
//	//class=ERROR
//	//	&TestBenchData_func_parseGpsdJSON{
//	//		Class:          "",
//	//		Payload:        []byte(""),
//	//		WantData:       &Data{},
//	//		WantErr:        nil,
//	//		WantErrVarName: `N/A`,
//	//		Func: func(test_item *TestBenchData_func_parseGpsdJSON) error {
//	//			var (
//	//				d *Data
//	//			)

//	//			return d, nil
//	//		},
//	//	},

//	//class=TPV
//	//	&TestBenchData_func_parseGpsdJSON{
//	//		Class:          "TPV",
//	//		Payload:        []byte(payloads["TPV"],
//	//		WantData:       &Data{},
//	//		WantErr:        nil,
//	//		WantErrVarName: `N/A`,
//	//		Func: func(test_item *TestBenchData_func_parseGpsdJSON) error {
//	//			var (
//	//				d *Data
//	//			)

//	//			return d, nil
//	//		},
//	//	},

//	//	&TestBenchData_func_parseGpsdJSON{
//	//		Class:          "",
//	//		Payload:        []byte(""),
//	//		WantData:       &Data{},
//	//		WantErr:        nil,
//	//		WantErrVarName: `N/A`,
//	//		Func: func(test_item *TestBenchData_func_parseGpsdJSON) error {
//	//			var (
//	//				d *Data
//	//			)

//	//			return d, nil
//	//		},
//	//	},
//	//	&TestBenchData_func_parseGpsdJSON{
//	//		Class:          "",
//	//		Payload:        []byte(""),
//	//		WantData:       &Data{},
//	//		WantErr:        nil,
//	//		WantErrVarName: `N/A`,
//	//		Func: func(test_item *TestBenchData_func_parseGpsdJSON) error {
//	//			var (
//	//				d *Data
//	//			)

//	//			return d, nil
//	//		},
//	//	},
//}

////TestFunc_parseGpsdJSON tests func parseGpsdJSON
//func TestFunc_parseGpsdJSON(t *testing.T) {
//	var (
//		got_err  error
//		data_err error
//		d        *Data
//	)
//	for _, test_item := range tb_parseGpsdJSON_table {
//		d = &Data{}
//		got_err = d.parseGpsdJSON(test_item.Class, test_item.Payload)

//		if test_item.HasFunc == true {
//			if test_item.Func == nil {
//				t.Fail()
//				t.Fatal("'test_item.Func' is nil.\n", "Test data:\n\t%s\n", test_item)
//				return
//			}
//			if data_err = test_item.Func(test_item); data_err != nil {
//				t.Error("TEST ERROR: Got incorrect gpsd.Data structure: ", d, ".\n\nWant: ",
//					test_item.WantData)
//				t.Errorf("Test data:\n\t%s\n", test_item)
//				t.Fail()
//				return
//			}
//		}

//		if got_err != nil {
//			if test_item.WantErr != nil {
//				if strings.Compare(got_err.Error(), test_item.WantErr.Error()) != 0 {
//					t.Errorf("TEST ERROR: Got error: %s\nWant error: %s \n\twith content: %s",
//						got_err.Error(), test_item.WantErrVarName, test_item.WantErr.Error())
//					t.Errorf("Test data:\n\t%s\n", test_item)
//					t.Fail()
//					return
//				}
//			}
//			if test_item.WantErr == nil {
//				t.Errorf("TEST ERROR: Got error: %s\nWant error: <nil> \n",
//					got_err.Error())
//				t.Errorf("Test data:\n\t%s\n", test_item)
//				t.Fail()
//				return
//			}
//		}
//		if got_err == nil {
//			if test_item.WantErr != nil {
//				t.Errorf("TEST ERROR: Got error: <nil>\nWant error: %s \n\twith content: %s",
//					test_item.WantErrVarName, test_item.WantErr.Error())
//				t.Errorf("Test data:\n\t%s\n", test_item)
//				t.Fail()
//				return
//			}
//		}
//	}
//}
