package gps

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/larsth/rmsgradiolinkctrld/errs"
)

const iso8601 = "2015-09-01T12:34:56+00:00"

func createJSONParserError(err error, class string) error {
	var (
		logBuf bytes.Buffer
	)

	logBuf.WriteString("JSON parser error with: class=\"")
	logBuf.WriteString(class)
	logBuf.WriteString("\", JSON parser error: \"")
	logBuf.WriteString(err.Error())
	logBuf.WriteString("\"\n\n")

	return fmt.Errorf(logBuf.String())
}

//Parse is a function that parses a gpsd JSON document.
//
//parse stores some or all of the GPS data from the following gpsd
//JSON documents:
// ERROR, POLL (stores only the 1st TPV item, and stores no SkyView items), TPV
//
//Parse logs all recieved gpsd JSON document it recieves.
//Errors are returned, logged, and printed to STDERR.
func (coord *GPSCoord) Parse(jsonDocument []byte) error {
	var (
		isFiltered bool
		class      Class
		logBuf     bytes.Buffer
		err        error
	)

	coord.mutex.Lock()
	defer coord.mutex.Unlock()

	if len(jsonDocument) == 0 {
		return errs.ErrZeroLengthPayload
	}

	//Get the type of gpsd JSON document ("class":...)
	if err = json.Unmarshal(jsonDocument, (&class)); err != nil {
		return err
	}

	if len(class.Class) == 0 {
		return errs.ErrEmptyString
	}

	isFiltered, err = filterGpsdJSON(class.Class, jsonDocument)
	if err != nil {
		return err
	}
	if isFiltered == true {
		return nil //nothing to do: the gpsd JSON document was filtered
	}

	//The jsonDocument was not filtered, so ...
	switch class.Class {
	case "POLL":
		if err = coord.parsePOLL(class.Class, jsonDocument); err != nil {
			return err
		}
		return nil //A POLL.TPV[0] document was inserted
	case "TPV":
		if err = coord.parseTPV(class.Class, jsonDocument); err != nil {
			return err
		}
		return nil //A TPV document was inserted
	}

	logBuf.WriteString("UNHANDLED gpsd JSON document\n")
	logBuf.WriteString("\tClass: ")
	logBuf.WriteString(class.Class)
	logBuf.WriteString("\n\n\tPayload: ")
	logBuf.Write(jsonDocument)
	logBuf.WriteString("\n\n")

	return fmt.Errorf(logBuf.String())
}
