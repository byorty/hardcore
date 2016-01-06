package controller

type Param struct {
	Name string `xml:"name,attr"`
	Required bool `xml:"required"`
	From From `xml:"from,attr"`
	Kind string `xml:"type,attr"`
}
