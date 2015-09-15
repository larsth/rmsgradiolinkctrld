package othergpshttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/larsth/rmsgradiolinkctrld/cache"
	"github.com/larsth/rmsgradiolinkctrld/gps"
	"github.com/larsth/rmsgradiolinkctrld/logging"
)

//Client is a go routine that connects to another GPS via another
//rmsgradiolinkctrld server.
func Client(otherServer string, doExit <-chan struct{},
	tickerDuration time.Duration) {
	var (
		gpsCoord  gps.GPSCoord
		ticker    *time.Ticker
		transport *http.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   20 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     nil,
			TLSHandshakeTimeout: 5 * time.Second,
			//Use compression:
			DisableCompression: false,
			//TCP connections are reused:
			DisableKeepAlives:     false,
			ResponseHeaderTimeout: 5 * time.Second,
		}
		client *http.Client = &http.Client{
			Transport: transport,
			//Timeout duration:
			Timeout: 20 * time.Second,
		}
		buf      bytes.Buffer
		decoder  *json.Decoder
		fetchUrl string
		r        *http.Response
		err      error
	)

	buf.WriteString(otherServer)
	buf.WriteString(EndPointPath)
	fetchUrl = buf.String()

	ticker = time.NewTicker(tickerDuration)
	for {
		select {
		case _ = <-doExit:
			ticker.Stop()
			ticker = nil
			transport = nil
			client = nil
			return
		case _ = <-ticker.C:
			if r, err = client.Get(buf.String()); err != nil {
				logging.Printf("ERROR, othergps: fetching %s: %s\n",
					fetchUrl, err.Error())
				gpsCoord.FetchUrl = fetchUrl
				gpsCoord.FetchError = err.Error()
				cache.OtherGPS.SetGPSCoord(gpsCoord)
				continue //reloops in the for loop
			}

			if r.Body == nil {
				msg := "ERROR No HTTP body in reponse from otherservers /otherserver HTTP endpoint"
				logging.Printf(msg)
				gpsCoord.FetchUrl = fetchUrl
				gpsCoord.FetchError = err.Error()
				cache.OtherGPS.SetGPSCoord(gpsCoord)
				continue //reloops in the for loop
			}

			decoder = json.NewDecoder(r.Body)
			if err = decoder.Decode(&gpsCoord); err != nil {
				format := "ERROR during othergps gpsd JSON document decoding %s: %s\n"
				reportedErr := fmt.Errorf(format, fetchUrl, err.Error())
				logging.Printf("%s\n", reportedErr.Error())
				gpsCoord.FetchUrl = fetchUrl
				gpsCoord.FetchError = reportedErr.Error()
				cache.OtherGPS.SetGPSCoord(gpsCoord)
				continue //reloops in the for loop
			}

			//else: success!
			gpsCoord.FetchUrl = ""
			gpsCoord.FetchError = ""
			cache.OtherGPS.SetGPSCoord(gpsCoord)
		}

		if r != nil {
			io.Copy(ioutil.Discard, r.Body)
			r.Body.Close()
		}
	}
}
