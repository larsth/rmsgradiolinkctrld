package gps

import (
	"encoding/json"
	"time"

	"github.com/larsth/rmsgradiolinkctrld/errs"
	"github.com/larsth/rmsgradiolinkctrld/gpsd"
	"github.com/larsth/rmsgradiolinkctrld/logging"
)

func (coord *GPSCoord) parseTPV(class string, jsonDocument []byte) error {
	var (
		t    time.Time
		terr error
		err  error
		TPV  gpsd.TPV
	)

	if len(jsonDocument) == 0 {
		return errs.ErrZeroLengthPayload
	}

	if err = json.Unmarshal(jsonDocument, &TPV); err != nil {
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
