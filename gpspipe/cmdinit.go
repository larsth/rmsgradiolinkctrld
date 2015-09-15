package gpspipe

import (
	"os/exec"

	"github.com/larsth/rmsgradiolinkctrld/logging"
)

func (c *Cmd) init_unsafe() error {
	if c.hasInited {
		return ErrCmdHadBeenInitialized
	}

	gpspipeWithPath, err := exec.LookPath("gpspipe")
	if err != nil {
		format := "FATAL ERROR: Cannot find gpsd client program `gpspipe`: \"%s"
		logging.Fatalf(format, err.Error())
	}

	c.cmd = exec.Command(gpspipeWithPath, gpspipeArgs...)
	c.hasInited = true
	return nil
}

func (c *Cmd) Init() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.init_unsafe()
}
