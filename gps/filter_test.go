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
	// TestBenchData_func_filterGpsdJSON structure,
	for _, item := range tb_filterGpsdJSON_table {
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

type TestBenchData_func_filterGpsdJSON struct {
	Class          string
	Payload        []byte
	WantIsFiltered bool
	WantErr        error
	WantErrVarName string
}

func (tb *TestBenchData_func_filterGpsdJSON) String() string {
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

var tb_filterGpsdJSON_table []*TestBenchData_func_filterGpsdJSON = []*TestBenchData_func_filterGpsdJSON{
	//len(payload) == 0 test data:
	&TestBenchData_func_filterGpsdJSON{
		Class:          "",
		Payload:        nil,
		WantIsFiltered: false,
		WantErr:        errs.ErrZeroLengthPayload,
		WantErrVarName: `errs.ErrZeroLengthPayload`,
	},
	//len(payload) != 0 && //len(class) == 0 test data:
	&TestBenchData_func_filterGpsdJSON{
		Class:          "",
		Payload:        []byte(`"class":"ATT"`),
		WantIsFiltered: false,
		WantErr:        errs.ErrEmptyString,
		WantErrVarName: `errs.ErrEmptyString`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "ATT",
		Payload:        []byte(`"class":"ATT"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "DEVICE",
		Payload:        []byte(`"class":"DEVICE"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "DEVICES",
		Payload:        []byte(`"class":"DEVICES"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "GST",
		Payload:        []byte(`"class":"GST"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "POLL",
		Payload:        []byte(`"class":"POLL"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "PPS",
		Payload:        []byte(`"class":"PPS"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "SKY",
		Payload:        []byte(`"class":"SKY"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "TOFF",
		Payload:        []byte(`"class":"TOFF"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "VERSION",
		Payload:        []byte(`"class":"VERSION"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "WATCH",
		Payload:        []byte(`"class":"WATCH"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `<nil>`,
	},
	&TestBenchData_func_filterGpsdJSON{
		Class:          "SATTELITE",
		Payload:        []byte(`"class":"SATTELITE"`),
		WantIsFiltered: true,
		WantErr:        nil,
		WantErrVarName: `N/A`,
	},
	&TestBenchData_func_filterGpsdJSON{
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
		got_isFiltered bool
		got_err        error
	)
	//silence logging:
	logging.SetIsLogging(false)

	for _, test_item := range tb_filterGpsdJSON_table {
		got_isFiltered, got_err = filterGpsdJSON(test_item.Class, test_item.Payload)
		if got_isFiltered != test_item.WantIsFiltered {
			t.Errorf("TEST ERROR. %s %v %s %v",
				"Unexpected 'isFiltered' bool value.\n Got: ",
				got_isFiltered, ". Want: ", test_item.WantIsFiltered)
			t.Errorf("\n\nTestdata: %s\n", test_item)
			t.Fail()
			continue
		}
		if (got_err != nil) && (test_item.WantErr != nil) {
			if strings.Compare(got_err.Error(), test_item.WantErr.Error()) != 0 {
				t.Errorf("TEST ERROR. Got another error: %s. Want: '%s' %s '%s'\n",
					got_err.Error(),
					test_item.WantErrVarName,
					" with the content: ",
					test_item.WantErr.Error())
				t.Errorf("\n\nTestdata: %s\n", test_item)
				t.Fail()
				continue
			}
		}
		if (got_err == nil) && (test_item.WantErr != nil) {
			t.Errorf("TEST ERROR. Got a nil error. Want: '%s' %s '%s'",
				test_item.WantErrVarName,
				" with the content: ",
				test_item.WantErr.Error())
			t.Errorf("\n\nTestdata: %s\n", test_item)
			t.Fail()
			continue
		}
		if (got_err != nil) && (test_item.WantErr == nil) {
			t.Errorf("TEST ERROR. Got another error: '%s'. Want: a nil error.\n",
				got_err.Error())
			t.Errorf("\n\nTestdata: %s\n", test_item)
			t.Fail()
			continue
		}
		//got_err and test_item.WantErr is either both nil, or both are
		//the same error (contains the same error string)
	}
}

//BenchmarkFunc_filterGpsdJSON benchmarks func filterGpsdJSON
func BenchmarkFunc_filterGpsdJSON(b *testing.B) {
	var (
		length int = len(tb_filterGpsdJSON_table)
		i      int = 0
	)

	//silence logging:
	logging.SetIsLogging(false)

	for j := 0; j < b.N; j++ {
		_, _ = filterGpsdJSON(
			tb_filterGpsdJSON_table[i].Class,
			tb_filterGpsdJSON_table[i].Payload)

		if (i + 1) == length {
			i = 0
		} else {
			i++
		}
	}
}
