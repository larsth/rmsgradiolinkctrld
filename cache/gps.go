package cache

import (
	"sync"

	"github.com/larsth/rmsgradiolinkctrld/gps"
)

var (
	ThisGPS  GPS
	OtherGPS GPS
)

type GPS struct {
	mutex sync.RWMutex
	//C        <-chan gps.GPSCoord
	gpsCoord gps.GPSCoord
}

func (g *GPS) GPSCoord() gps.GPSCoord {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.gpsCoord
}

func (g *GPS) SetGPSCoord(newGPSCoord gps.GPSCoord) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.gpsCoord = newGPSCoord
}
