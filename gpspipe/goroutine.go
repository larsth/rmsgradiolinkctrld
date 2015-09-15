package gpspipe

import (
	"time"

	"github.com/larsth/rmsgradiolinkctrld/gps"
)

type GPSCoordGoRoutineArgs struct {
	GPSPipeTickerDuration time.Duration

	//'DoneChan' must be a buffered channel with a capacity of 1
	DoneChan chan struct{}

	//unbuffered channel
	DoExitChan chan struct{}

	//unbuffered channel
	GpsDataChan chan gps.GPSCoord

	//'ErrOutChan' must be a buffered channel with a capacity of 1
	ErrOutChan chan error
}

//StreamData is a go routine which filters, parses, and finally stream
//gpsd JSON documents that had been recieved from the gpspipe program.
//
//The actual gpsd JSON document lexical analysis, filtering, and parsing happens
//in the gpsd package.
func (c *Cmd) StreamData(args GPSCoordGoRoutineArgs) {
	var (
		gptDuration time.Duration
		gptTicker   *time.Ticker
		s           struct{}
		err         error
		coord       gps.GPSCoord
	)

	if err = c.startGpsPipe(); err != nil {
		args.DoneChan <- s
		args.ErrOutChan <- err
		return
	}

	//setup:
	if args.GPSPipeTickerDuration.Nanoseconds() < (time.Millisecond * 8).Nanoseconds() {
		gptDuration = time.Millisecond * 8000 //force use of 8ms
	} else {
		gptDuration = args.GPSPipeTickerDuration
	}

	gptTicker = time.NewTicker(gptDuration)

	//work:
	//This is a go routine, so loop for infinity
	// - until a message on the chanArgs.DoExitChan channel had been recieved:
	for {
		select {
		case _ = <-args.DoExitChan:
			args.DoneChan <- s
			args.ErrOutChan <- error(nil)
			return
		case _ = <-gptTicker.C: // 'gptDuration' amount of time had elapsed, so run:
			if err = c.rx(&coord); err != nil {
				args.ErrOutChan <- err
			} else {
				//Th result of recieving the sended gpsd.GPSCoord should be
				//copied into the 'This' struct field of type gpsd.GPSCoord:
				args.GpsDataChan <- coord //'coord' is a value, so it is copied
			}
		}
	}
}
