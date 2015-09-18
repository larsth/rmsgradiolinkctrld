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
//	Func           func(testItem *TestBenchData_func_parseGpsdJSON) error
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
//	//		Func: func(testItem *TestBenchData_func_parseGpsdJSON) error {
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
//	//		Func: func(testItem *TestBenchData_func_parseGpsdJSON) error {
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
//	//		Func: func(testItem *TestBenchData_func_parseGpsdJSON) error {
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
//	//		Func: func(testItem *TestBenchData_func_parseGpsdJSON) error {
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
//		gotErr  error
//		data_err error
//		d        *Data
//	)
//	for _, testItem := range tb_parseGpsdJSON_table {
//		d = &Data{}
//		gotErr = d.parseGpsdJSON(testItem.Class, testItem.Payload)

//		if testItem.HasFunc == true {
//			if testItem.Func == nil {
//				t.Fail()
//				t.Fatal("'testItem.Func' is nil.\n", "Test data:\n\t%s\n", testItem)
//				return
//			}
//			if data_err = testItem.Func(testItem); data_err != nil {
//				t.Error("TEST ERROR: Got incorrect gpsd.Data structure: ", d, ".\n\nWant: ",
//					testItem.WantData)
//				t.Errorf("Test data:\n\t%s\n", testItem)
//				t.Fail()
//				return
//			}
//		}

//		if gotErr != nil {
//			if testItem.WantErr != nil {
//				if strings.Compare(gotErr.Error(), testItem.WantErr.Error()) != 0 {
//					t.Errorf("TEST ERROR: Got error: %s\nWant error: %s \n\twith content: %s",
//						gotErr.Error(), testItem.WantErrVarName, testItem.WantErr.Error())
//					t.Errorf("Test data:\n\t%s\n", testItem)
//					t.Fail()
//					return
//				}
//			}
//			if testItem.WantErr == nil {
//				t.Errorf("TEST ERROR: Got error: %s\nWant error: <nil> \n",
//					gotErr.Error())
//				t.Errorf("Test data:\n\t%s\n", testItem)
//				t.Fail()
//				return
//			}
//		}
//		if gotErr == nil {
//			if testItem.WantErr != nil {
//				t.Errorf("TEST ERROR: Got error: <nil>\nWant error: %s \n\twith content: %s",
//					testItem.WantErrVarName, testItem.WantErr.Error())
//				t.Errorf("Test data:\n\t%s\n", testItem)
//				t.Fail()
//				return
//			}
//		}
//	}
//}
