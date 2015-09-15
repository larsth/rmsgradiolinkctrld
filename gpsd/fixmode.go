package gpsd

//FixMode is a type used for indication no GPS fix, 2D GPS fix, and 3D GPS fix.
type FixMode byte

func (f FixMode) String() string {
	switch f {
	case FixNotSeen:
		return "Not Seen"
	case FixNone:
		return "None"
	case Fix2D:
		return "2D"
	case Fix3D:
		return "3D"
	}
	return "" //make compiler happy
}

const (
	//NoValueSeen means that there is no knowledge of what kind fix a GPS has.
	FixNotSeen FixMode = 0
	//NoFix means that the GPS hasn´t a fix.
	FixNone FixMode = 1
	//Mode2D means that the GPS has a 2D fix.
	Fix2D FixMode = 2
	//Mode3D means that the GPS has a 3D fix.
	Fix3D FixMode = 3
)
