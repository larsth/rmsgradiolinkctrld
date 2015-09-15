package gps

//Type Class is used to determin the value of the "class" field in gpsd
//JSON documents.
type Class struct {
	//Always? Yes. Type: string
	Class string `json:"class"`
}
