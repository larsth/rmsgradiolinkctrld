package gps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/larsth/rmsgradiolinkctrld/errs"
	"github.com/larsth/rmsgradiolinkctrld/gpsd"
	"github.com/larsth/rmsgradiolinkctrld/logging"
)

const iso8601 = "2015-09-01T12:34:56+00:00"

func createJSONParserError(err error, class string) error {
	var (
		logBuf bytes.Buffer
	)

	logBuf.WriteString("JSON parser error: \"")
	logBuf.WriteString(err.Error())

	logBuf.WriteString("\"\n\nClass was: ")
	logBuf.WriteString(class)
	logBuf.WriteString("\n")

	return fmt.Errorf(logBuf.String())
}

func (coord *GPSCoord) parseTPV(class string) error {
	var (
		t    time.Time
		terr error
		err  error
		TPV  gpsd.TPV
	)

	if err = json.Unmarshal(coord.jsonDocument, &TPV); err != nil {
		return createJSONParserError(err, class)
	}
	t, terr = time.Parse(iso8601, TPV.Time)
	if terr == nil {
		coord.RecievedTimeStamp = t
		msg := "time.Parse error parsing a TPV.Time with ISO8601 formatting. Error"
		logging.Printf("%s: '%s'", msg, terr.Error())
	} else {
		coord.RecievedTimeStamp = time.Now().UTC()
	}
	coord.Fix = TPV.Fix
	coord.Lat = TPV.Lat
	coord.Lon = TPV.Lon
	coord.Alt = TPV.Alt

	return nil
}

func (coord *GPSCoord) parsePOLL(class string) error {
	var (
		t    time.Time
		terr error
		err  error
		POLL gpsd.POLL
	)

	if err = json.Unmarshal(coord.jsonDocument, &POLL); err != nil {
		return createJSONParserError(err, class)
	}
	if len(POLL.Fixes) > 0 {
		t, terr = time.Parse(iso8601, POLL.Fixes[0].Time)
		if terr == nil {
			coord.RecievedTimeStamp = t
			msg := "time.Parse error parsing a TPV.Time with ISO8601 formatting."
			logging.Printf("%s: '%s'", msg, terr.Error())
		} else {
			coord.RecievedTimeStamp = time.Now().UTC()
		}
		coord.Fix = POLL.Fixes[0].Fix
		coord.Lat = POLL.Fixes[0].Lat
		coord.Lon = POLL.Fixes[0].Lon
		return nil
	} else {
		return ErrNoPollFixes
	}
}

//parse is a function that parses a gpsd JSON document.
//
//parse stores some or all of the GPS data from the following gpsd
//JSON documents:
// ERROR, POLL (stores only the 1st TPV item, and stores no SkyView items), TPV
//
//parse logs all recieved gpsd JSON document it recieves.
//Errors are returned, logged, and printed to STDERR.
func (coord *GPSCoord) Parse() error {
	var (
		isFiltered bool
		class      Class
		logBuf     bytes.Buffer
		err        error
	)

	coord.mutex.Lock()
	defer coord.mutex.Unlock()

	if len(coord.jsonDocument) == 0 {
		return errs.ErrZeroLengthPayload
	}

	//This deferred function sets coord.JsonDocument = nil, after a 'return', so
	//coord.ReadFrom can write yet another gpsd JSON document into
	//coord.jsonDocument
	defer func() {
		coord.jsonDocument = nil
		//FIXME replace this with a release of a byte slice to a memory buffer,
		//where we first need to fill the soon-to-be-released byte slice with
		// '\0's, and then create a new slice header from the old slice with len=0,
		//and a capacity equal to the capacity of the old slice.
	}()

	//Get the type of gpsd JSON document ("class":...)
	if err = json.Unmarshal(coord.jsonDocument, (&class)); err != nil {
		return err
	}

	if len(class.Class) == 0 {
		return errs.ErrEmptyString
	}

	isFiltered, err = filterGpsdJSON(class.Class, coord.jsonDocument)
	if err != nil {
		return err
	}
	if isFiltered == true {
		return nil //nothing to do: the gpsd JSON document was filtered
	}

	//The jsonDocument was not filtered, so ...
	switch class.Class {
	case "POLL":
		if err = coord.parsePOLL(class.Class); err != nil {
			return err
		}
		return nil //A POLL.TPV[0] document was inserted
	case "TPV":
		if err = coord.parseTPV(class.Class); err != nil {
			return err
		}
		return nil //A TPV document was inserted
	}

	logBuf.WriteString("UNHANDLED gpsd JSON document\n")
	logBuf.WriteString("\tClass: ")
	logBuf.WriteString(class.Class)
	logBuf.WriteString("\n\n\tPayload: ")
	logBuf.Write(coord.jsonDocument)
	logBuf.WriteString("\n\n")

	return fmt.Errorf(logBuf.String())
}
