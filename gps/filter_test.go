//+build go1.5

package gps

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/larsth/rmsgradiolinkctrld/errs"
	"github.com/larsth/rmsgradiolinkctrld/logging"
)

func init() {
	var buf bytes.Buffer

	//create correct WantErr and WantErrVarName for the
	// TestBenchDataFuncFilterGpsdJSON structure,
	for _, item := range tbFilterGpsdJSONTable {
		//where Class is "WRONG_CLASS_NAME":
		if strings.Compare("WRONG_CLASS_NAME", item.Class) == 0 {
			item.WantErr = fmt.Errorf("%s :'%s' (%v)",
				"FATAL ERROR: UNKNOWN gpsd JSON document",
				item.Class, item.Payload)
			item.WantErrVarName = "N/A"
		}
		//where class is "SATTELITE":
		if strings.Compare("SATTELITE", item.Class) == 0 {
			buf.WriteString("ERROR: An unexpected gpsd JSON document was recieved. ")
			buf.WriteString("A \"Sattelite\" gpsd JSON document should ")
			buf.WriteString("NOT come directly from gpsd (via gpspipe) - ")
			buf.WriteString(" it should be part of a SKY JSON document.")

			str := buf.String()
			item.WantErr = fmt.Errorf("%s", str)
			item.WantErrVarName = "N/A"
		}
	}
}

type TestBenchDataFuncFilterGpsdJSON struct {
	Class          string
	Payload        []byte
	WantIsFiltered bool
	WantErr        error
	WantErrVarName string
}

func (tb *TestBenchDataFuncFilterGpsdJSON) String() string {
	var buf bytes.Buffer

	buf.WriteString("\n\tClass: \t")
	buf.WriteString(tb.Class)
	buf.WriteString("\n\tPayload: \t")
	buf.Write(tb.Payload)
	buf.WriteString("\n\tWantIsFiltered: \t")
	buf.WriteString(fmt.Sprintln(tb.WantIsFiltered))
	buf.WriteString("\n\tWantErr: \t")
	buf.WriteString(fmt.Sprintln(tb.WantErr))
	buf.WriteString("\n\tWantErrVarName: \t")
	buf.WriteString(tb.WantErrVarName)

	return buf.String()
}

var tbFilterGpsdJSONTable = []*TestBenchDataFuncFilterGpsdJSON{
	//len(payload) == 0 test data:
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "",
		Payload:        nil,
		WantIsFiltered: false,
		WantErr:        errs.ErrZeroLengthPayload,
		WantErrVarName: `errs.ErrZeroLengthPayload`,
	},
	//len(payload) != 0 && //len(class) == 0 test data:
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "",
		Payload:        []byte(`"class":"ATT"`),
		WantIsFiltered: false,
		WantErr:        errs.ErrEmptyString,
		WantErrVarName: `errs.ErrEmptyString`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "ATT",
		Payload:        []byte(`"class":"ATT"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "DEVICE",
		Payload:        []byte(`"class":"DEVICE"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "DEVICES",
		Payload:        []byte(`"class":"DEVICES"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "GST",
		Payload:        []byte(`"class":"GST"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "POLL",
		Payload:        []byte(`"class":"POLL"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "PPS",
		Payload:        []byte(`"class":"PPS"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "SKY",
		Payload:        []byte(`"class":"SKY"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "TOFF",
		Payload:        []byte(`"class":"TOFF"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "VERSION",
		Payload:        []byte(`"class":"VERSION"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "WATCH",
		Payload:        []byte(`"class":"WATCH"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "SATTELITE",
		Payload:        []byte(`"class":"SATTELITE"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `N/A`,
	},
	&TestBenchDataFuncFilterGpsdJSON{
		Class:          "TPV",
		Payload:        []byte(`"class":"TPV"`),
		WantIsFiltered: false,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
}

//TestFunc_filterGpsdJSON tests func filterGpsdJSON
func TestFunc_filterGpsdJSON(t *testing.T) {
	var (
		gotIsFiltered bool
		gotErr        error
	)
	//silence logging:
	logging.SetIsLogging(false)

	for _, testItem := range tbFilterGpsdJSONTable {
		gotIsFiltered, gotErr = filterGpsdJSON(testItem.Class, testItem.Payload)
		if gotIsFiltered != testItem.WantIsFiltered {
			t.Errorf("TEST ERROR. %s %v %s %v",
				"Unexpected 'isFiltered' bool value.\n Got: ",
				gotIsFiltered, ". Want: ", testItem.WantIsFiltered)
			t.Errorf("\n\nTestdata: %s\n", testItem)
			t.Fail()
			continue
		}
		if (gotErr != nil) && (testItem.WantErr != nil) {
			if strings.Compare(gotErr.Error(), testItem.WantErr.Error()) != 0 {
				t.Errorf("TEST ERROR. Got another error: %s. Want: '%s' %s '%s'\n",
					gotErr.Error(),
					testItem.WantErrVarName,
					" with the content: ",
					testItem.WantErr.Error())
				t.Errorf("\n\nTestdata: %s\n", testItem)
				t.Fail()
				continue
			}
		}
		if (gotErr == nil) && (testItem.WantErr != nil) {
			t.Errorf("TEST ERROR. Got a nil error. Want: '%s' %s '%s'",
				testItem.WantErrVarName,
				" with the content: ",
				testItem.WantErr.Error())
			t.Errorf("\n\nTestdata: %s\n", testItem)
			t.Fail()
			continue
		}
		if (gotErr != nil) && (testItem.WantErr == nil) {
			t.Errorf("TEST ERROR. Got another error: '%s'. Want: a nil error.\n",
				gotErr.Error())
			t.Errorf("\n\nTestdata: %s\n", testItem)
			t.Fail()
			continue
		}
		//gotErr and testItem.WantErr is either both nil, or both are
		//the same error (contains the same error string)
	}
}

//BenchmarkFunc_filterGpsdJSON benchmarks func filterGpsdJSON
func BenchmarkFuncFilterGpsdJSON(b *testing.B) {
	var (
		length = len(tbFilterGpsdJSONTable)
		i      int
	)

	//silence logging:
	logging.SetIsLogging(false)

	for j := 0; j < b.N; j++ {
		_, _ = filterGpsdJSON(
			tbFilterGpsdJSONTable[i].Class,
			tbFilterGpsdJSONTable[i].Payload)

		if (i + 1) == length {
			i = 0
		} else {
			i++
		}
	}
}
