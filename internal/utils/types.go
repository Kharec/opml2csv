package utils

type Outline struct {
	Flux    string `xml:"type,attr"`
	Text    string `xml:"text,attr"`
	Title   string `xml:"title,attr"`
	XmlUrl  string `xml:"xmlUrl,attr"`
	HtmlUrl string `xml:"htmlUrl,attr"`
}
