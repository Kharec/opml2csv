package utils

func BuildCSVLine(outline *Outline) []string {
	return []string{outline.Flux, outline.Text, outline.Title, outline.XmlUrl, outline.HtmlUrl}
}

func BuildCSVHeader() []string {
	return []string{"feed", "text", "title", "xmlUrl", "htmlUrl"}
}
