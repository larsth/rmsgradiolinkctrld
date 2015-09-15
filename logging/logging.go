package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/larsth/rmsgradiolinkctrld/errs"
)

var isLogging bool

func init() {
	isLogging = true
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime | log.LUTC)
	log.SetPrefix("rmsgradiolinkctrld")
}

func LogUnusedGpsdJSONDocument(class string, payload []byte) {
	var logBuf bytes.Buffer
	if isLogging == true {
		if len(class) == 0 {
			return
		}
		if len(payload) == 0 {
			return
		}

		logBuf.WriteString("Logging an UNUSED \"class\":\"")
		logBuf.WriteString(class)
		logBuf.WriteString("\" gpsd JSON document.")
		logBuf.WriteString("\n\n\tJSON Payload is:\n\n\t\t")
		logBuf.Write(payload)

		log.Println(logBuf.String())
		fmt.Fprintln(os.Stdout, logBuf.String())
	}
}

func Printf(format string, v ...interface{}) error {
	var str string
	if isLogging == true {
		if len(format) == 0 {
			return fmt.Errorf("%s (\"%s\"-argument)", errs.ErrEmptyString.Error(), "format")
		}

		str = fmt.Sprintf(format, v)

		log.Println(str)
		fmt.Fprintln(os.Stderr, str)
	}
	return nil
}

func Fatalf(format string, v ...interface{}) {
	var str string

	if isLogging == true {
		str = fmt.Sprintf(format, v)

		fmt.Fprintln(os.Stderr, str)
		log.Fatalln(str)
	}
}

func SetIsLogging(v bool) {
	isLogging = v
}
