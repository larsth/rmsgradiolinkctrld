package gps

import (
	"encoding/json"
	"time"

	"github.com/larsth/rmsgradiolinkctrld/errs"
	"github.com/larsth/rmsgradiolinkctrld/gpsd"
	"github.com/larsth/rmsgradiolinkctrld/logging"
)

func (coord *GPSCoord) parsePOLL(class string, jsonDocument []byte) error {
	var (
		t    time.Time
		terr error
		err  error
		POLL gpsd.POLL
	)

	if len(jsonDocument) == 0 {
		return errs.ErrZeroLengthPayload
	}

	if err = json.Unmarshal(jsonDocument, &POLL); err != nil {
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
