//Package gpsd contains gpsd data structures, and the gps data structures
//that this program uses.
package gps

import "bytes"

//Type Data is the filtered gps data structure that the
//rmsgradiolinkctrld application is using.
//
//Compared to gpsd JSON data structures - this data structure is "class"-less
// (only 1 structure), which is continuously updated, and is for this reason
// made concurrency safe with a read-write mutex, because it is accessed
// concurrently by more than 1 go routine.
type Data struct {
	ThisGps  GPSCoord `json:"this_gps"`
	OtherGps GPSCoord `json:"other_gps"`
	_        struct{} // to prevent unkeyed literals
}

func (d *Data) String() string {
	var buf bytes.Buffer

	buf.WriteString("(*Data)\n")

	buf.WriteString("\t.ThisGPS")
	buf.WriteString(d.ThisGps.String())

	buf.WriteString("\t.OtherGPS")
	buf.WriteString(d.OtherGps.String())

	return buf.String()
}
