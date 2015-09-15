//Package gpspipe calls a gpsd client program called gpspipe, and filters
//the output from the gpspipe program.
package gpspipe

import (
	"io"
	"os/exec"
	"sync"

	"github.com/larsth/rmsgradiolinkctrld/gps"
)

var (
	gpspipeArgs = []string{
		"-w", "-l", "-s", "localhost:gpsd",
	} // gpspipe -w -l localhost:gpsd
)

type Cmd struct {
	mutex     sync.Mutex
	hasInited bool
	cmd       *exec.Cmd
	cmdStdOut io.ReadCloser
	_         struct{}
}

func New() *Cmd {
	c := new(Cmd)

	//Below:
	//The error is ignored, because the error is always nil
	//the first time init_unsafe is called
	_ = c.init_unsafe()

	return c
}

func (c *Cmd) startGpsPipe() (err error) {
	if c.hasInited == false {
		return ErrCmdHadNotBeenInitialized
	}
	if c.cmdStdOut, err = c.cmd.StdoutPipe(); err != nil {
		return
	}
	if err = c.cmd.Start(); err != nil {
		return
	}
	return nil
}

//rx recieves a stream of bytes from gpspipe containing gpsd JSON documents.
func (c *Cmd) rx(coord *gps.GPSCoord) error {
	panic("FIXME: 'rx' metoden laves færdig")
	//FIXME 'rx' metoden laves færdig
	return nil
}
