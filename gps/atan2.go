package gps

import (
	"fmt"
	"math"
)

//
//FIXME: let rmsgAtan2 wrap math.Atan2 - the wrapper contains error detection and reporting.
func rmsgAtan2(x, y float64) (float64, error) {
	if y > 0.0 {
		if x > 0.0 {
			return math.Atan((y / x)), nil
		}
		if x < 0.0 {
			return (math.Pi - math.Atan((-y / x))), nil
		}
		if x == 0.0 {
			return (math.Pi / 2), nil
		}
		return math.NaN(), fmt.Errorf("X is not <, > or equal to 0.0. x:=\"%d\"\n\n", x)
	}
	if y < 0.0 {
		if x > 0.0 {
			return (-1 * math.Atan((-y / x))), nil
		}
		if x < 0.0 {
			return (math.Atan((y / x)) - math.Pi), nil
		}
		if x == 0.0 {
			return ((3 * math.Pi) / 2), nil //270.0 grader og nil
		}
		return math.NaN(), fmt.Errorf("X is not <, > or equal to 0.0. x:=\"%d\"\n\n", x)
	}
	if y == 0.0 {
		if x > 0.0 {
			return 0.0, nil
		}
		if x < 0.0 {
			return math.Pi, nil //180 grader og nil
		}
		if x == 0.0 {
			return math.NaN(), ErrSameLocation
		}
		return math.NaN(), fmt.Errorf("X is not <, > or equal to 0.0. x:=\"%d\"\n\n", x)
	}
	// y is not <, >, or equal to 0.0!
	return math.NaN(), fmt.Errorf("Y is not <, > or equal to 0.0. y:=\"%d\"\n\n", y)
}
