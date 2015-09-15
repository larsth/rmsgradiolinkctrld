//+build go1.5
package gps

var payloads map[string]string = map[string]string{
	"ATT":              ``,
	"DEVICE":           ``,
	"DEVICES":          ``,
	"ERROR":            ``,
	"GST":              ``,
	"POLL":             ``,
	"PPS":              ``,
	"Sattelite":        ``,
	"SKY":              ``,
	"TOFF":             ``,
	"TPV":              ``,
	"VERSION":          ``,
	"WATCH":            ``,
	"WRONG_CLASS_NAME": ``,
}

var list_of_payloads []string = []string{
	"ATT",
	"DEVICE",
	"DEVICES",
	"ERROR",
	"GST",
	"POLL",
	"PPS",
	"Sattelite",
	"SKY",
	"TOFF",
	"TPV",
	"VERSION",
	"WATCH",
	"WRONG_CLASS_NAME",
}
