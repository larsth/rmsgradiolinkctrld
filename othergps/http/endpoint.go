package othergpshttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/larsth/rmsgradiolinkctrld/cache"
	"github.com/larsth/rmsgradiolinkctrld/gps"
)

const EndPointPath = `/othergps`

//Endpoint is the HTTP endpoint for EndPointPath.
func Endpoint(w http.ResponseWriter, r *http.Request) {
	var (
		gpsCoord gps.GPSCoord
		p        []byte
		err      error
		buf      bytes.Buffer
	)
	//Below:
	//Discard the content from the request body, if not already empty at return
	defer io.Copy(ioutil.Discard, r.Body)

	w.Header().Add("Server", "rmsgradiolinkctrld/0.1 (Unix) (RaspberryPi1B+/Linux)")

	//caching rules for this HTTP response:
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Pragma", "no-cache")
	//'Expires' is a HTTP response header with a date and time value
	//in the past, which are for old HTTP protocol(<1.1) proxies and browsers:
	w.Header().Add("Expires", "Thu, 01 Jan 2001 01:23:45 GMT")

	if strings.Contains(r.Method, "GET") == true {
		// "GET" HTTP method request:

		//about 'cache.ThisGPS.GPSCoord()´:
		//	relative to the other end - this GPS coordinate is "id"="othergps",
		//	so fetch cache.ThisGPS.GPSCoord and set ID="othergps":
		gpsCoord = cache.ThisGPS.GPSCoord()
		gpsCoord.ID = "othergps"

		if p, err = json.Marshal(&gpsCoord); err != nil {
			buf.WriteString("500 Internal server error\r\n\r\n")
			buf.WriteString(err.Error())
			buf.WriteString("\r\n")
			http.Error(w, buf.String(), http.StatusInternalServerError)
		} else {
			buf.Write(p)
			w.WriteHeader(http.StatusOK)
			w.Write(buf.Bytes())
		}
		return
	}

	//implicit else/fallthrough:
	w.Header().Set("Allow", "GET") //the only allowed HTTP method is GET.
	http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
}
