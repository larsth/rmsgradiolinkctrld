package gps

import (
	"bytes"
	"fmt"

	"github.com/larsth/rmsgradiolinkctrld/errs"
	"github.com/larsth/rmsgradiolinkctrld/logging"
)

//filterGpsdJSON is a function that filters a gpsd JSON document.
//
//parseGpsdJSON ignores the following gpsd JSON document types:
//  - Everything but TPV gpsd JSON documents
//
//filterGpsdJSON logs all recieved gpsd JSON documents it ignores.
//Errors are logged, printed to STDERR, and returned.
func filterGpsdJSON(class string, payload []byte) (bool, error) {
	var buf bytes.Buffer

	if len(payload) == 0 {
		logging.Printf("%s", errs.ErrZeroLengthPayload.Error())
		return false, errs.ErrZeroLengthPayload
	}
	if len(class) == 0 {
		logging.Printf("%s", errs.ErrEmptyString.Error())
		return false, errs.ErrEmptyString
	}

	switch class {
	case "ATT":
		//ignore ATT reports from gpsd
		// - we IS using a compass, but it is unreachable from gpsd.
		//If an ATT gpsd JSON document type is recieved it is logged.
		fallthrough
	case "DEVICE":
		//ignore "class":"DEVICE" reports from gpsd
		// - we are not changeing anything with any device, but it is logged.
		fallthrough
	case "DEVICES":
		//ignore "class":"DEVICES" reports from gpsd
		// - we are not using the list of devices per gpsd server for anything
		// but logging, so a DEVICES class JSON document is only logged.
		fallthrough
	case "GST":
		//Currently pseudorange noise reports are not used, so thoose
		//reports are filtered out, but they are logged
		fallthrough
	case "POLL":
		//We are not transmitting a "?WATCH" command to gpsd, so we should not see
		//any "POOL"
		fallthrough
	case "PPS":
		//PPS is type that contains a message which is emitted each
		//time the gpsd server sees a PPS (Pulse Per Second) strobe
		// from a device.
		//
		//Currently this time-correction information is not used,
		//so a PPS gpsd JSON document is filtered out, but it is logged.
		fallthrough
	case "SKY":
		//A gpsd SKY JSON document/report, with its array of Sattelite, is not
		//currently used, but it is logged
		//log it
		fallthrough
	case "TOFF":
		//TOFF is a type that contains the message that is emitted on each cycle
		//and reports the offset between the host's clock time and the GPS time
		//at top of second (actually, when the first data for the reporting
		//cycle is received).
		//
		//TOFF gpsd JSON reports is not currently used.
		//The host´s clock time must be adjusted by another
		//priviledged program - not this program.
		fallthrough
	case "VERSION":
		//gps version information is only logged - for debugging usage only
		fallthrough
	case "WATCH":
		//gpspipe issues a "?WATCH" comamnd and we get a "WATCH" gpsd JSON
		//document inreturn, which we only logs - for debugging usage only

		logging.LogUnusedGpsdJSONDocument(class, payload)
		return true, nil
	case "SATTELITE":
		buf.WriteString("ERROR: An unexpected gpsd JSON document was recieved. ")
		buf.WriteString("A \"Sattelite\" gpsd JSON document should ")
		buf.WriteString("NOT come directly from gpsd (via gpspipe) - ")
		buf.WriteString(" it should be part of a SKY JSON document.")

		str := buf.String()
		logging.Printf("%s", str)
		err := fmt.Errorf("%s", str)
		return true, err
	}

	return false, nil // (not filtered, no error)
}
