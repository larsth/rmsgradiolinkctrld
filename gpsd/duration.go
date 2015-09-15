package gpsd

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

//Duration is a type that embeds time.Duration, and has support for JSON.
//
//The duration is stored in seconds, if marshaled to JSON.
type Duration struct {
	time.Duration
}

//UnmarshalJSON can unmarshal a JSON description of itself.
//
//UnmarshalJSON implements interface encoding/json.Unmarshaler
func (d *Duration) UnmarshalJSON(data []byte) error {
	var (
		str           string
		hasUnitSuffix bool
		err           error
		duration      time.Duration
		buf           bytes.Buffer
	)

	str = strings.TrimSpace(string(data))
	buf.WriteString(str)

	//"300ms", "-1.5h" or "2h45m".
	//Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	hasUnitSuffix = ((-1 != strings.IndexAny(str, "ns")) &&
		(-1 != strings.IndexAny(str, "us")) &&
		(-1 != strings.IndexAny(str, "µs")) &&
		(-1 != strings.IndexAny(str, "ms")) &&
		(-1 != strings.IndexAny(str, "s")) &&
		(-1 != strings.IndexAny(str, "m")) &&
		(-1 != strings.IndexAny(str, "s")))

	if false == hasUnitSuffix {
		//assume seconds (written by the MarshalJSON method)
		buf.Write([]byte("s"))
	}

	if duration, err = time.ParseDuration(buf.String()); err != nil {
		return err
	}
	d.Duration = duration
	return nil
}

//MarshalJSON can marshal itself into valid JSON.
//
//MarshalJSON implements interface encoding/json.Marshaler
func (d *Duration) MarshalJSON() ([]byte, error) {
	f := d.Duration.Seconds()
	b := make([]byte, 0, 16)
	b = strconv.AppendFloat(b, f, byte('f'), -1, 64)
	return b, nil
}
