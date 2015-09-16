package gps

import (
	"bytes"
	"errors"
	"io"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/larsth/rmsgradiolinkctrld/gpsd"
)

var (
	ErrSameLocation                  error = errors.New("The 2 GPS coordinates are the same location.")
	ErrOtherGPSCoordIsNil            error = errors.New("The `other` *GPSCoord pointer is nil.")
	ErrThisGPSCoordHasZeroTimestamp  error = errors.New("The `this` TimeStamp is Zero: 0")
	ErrOtherGPSCoordHasZeroTimestamp error = errors.New("The `other`TimeStamp is Zero: 0")
)

var (
	ErrNoPollFixes = errors.New("No POLL fixes (No TPV JSON documents)")
)

type GPSCoord struct {
	mutex        sync.Mutex `"json:-"`
	ioReader     io.Reader  `"json:-"`
	scanBuf      []byte     `"json:-"`
	jsonDocument []byte     `"json:-"`

	ID                string       `"json:id"`
	RecievedTimeStamp time.Time    `"json:datetime_iso3339"`
	Fix               gpsd.FixMode `"json:fix"` //FixNotSeen,FixNone,Fix2D,Fix3D

	Lat float64 `"json:lat"` //Lattitude - breddegrad
	Lon float64 `"json:lon"` //Longitude - længdegrad
	Alt float64 `"json:alt"` //Altitude - højden

	GpsdError         string    `"json:gpsd_error,omitempty"`
	GpsdErrorDateTime time.Time `"json:gpsd_error_datetime_iso3339,omitempty"`
	FetchUrl          string    `"json:fetch_url,omitempty"`
	FetchError        string    `"json:fetch_error,omitempty"`
}

func (this *GPSCoord) String() string {
	var buf bytes.Buffer

	buf.WriteString("\t\tRecievedTimeStamp: ")
	buf.WriteString(this.RecievedTimeStamp.UTC().Format(time.RFC3339))
	buf.WriteString("\n")

	buf.WriteString("\t\tFix: ")
	buf.WriteString(this.Fix.String())
	buf.WriteString("\n")

	buf.WriteString("\t\tLat: ")
	buf.WriteString(strconv.FormatFloat(this.Lat, byte('f'), -1, 64))
	buf.WriteString("\n")

	buf.WriteString("\t\tLon: ")
	buf.WriteString(strconv.FormatFloat(this.Lon, byte('f'), -1, 64))
	buf.WriteString("\n")

	buf.WriteString("\t\tGpsdError: ")
	buf.WriteString(this.GpsdError)
	buf.WriteString("\n")

	return buf.String()
}

//Bearing calculates the bearing (in degrees) between 2 GPS coordinates
//('this' relative to 'other').
func (this *GPSCoord) Bearing(other *GPSCoord) (float64, error) {
	var (
		dlon      float64
		x, x1, x2 float64
		y         float64
		radians   float64
		err       error
	)

	dlon = other.Lon - this.Lon
	y = math.Sin(dlon) * math.Cos(other.Lat)
	x1 = math.Cos(this.Lat) * math.Sin(other.Lat)
	x2 = math.Sin(this.Lat) * math.Cos(other.Lat)
	x = (x1 - x2) * math.Cos(dlon)

	if radians, err = rmsgAtan2(x, y); err != nil {
		return math.NaN(), err
	}
	return radians * (180.0 / math.Pi), nil
}

//Duration calculates the duration in recieve-datetime for 2
//GPS coordinates('this' relative to 'other').
//If the resulting Duration is negative, then 'other' is before 'this'.
func (this *GPSCoord) Duration(other *GPSCoord) (time.Duration, error) {
	if other == nil {
		return time.Duration(0), ErrOtherGPSCoordIsNil
	}
	if this.RecievedTimeStamp.IsZero() {
		return time.Duration(0), ErrThisGPSCoordHasZeroTimestamp
	}
	if other.RecievedTimeStamp.IsZero() {
		return time.Duration(0), ErrOtherGPSCoordHasZeroTimestamp
	}
	return this.RecievedTimeStamp.UTC().Sub(other.RecievedTimeStamp.UTC()), nil
}
