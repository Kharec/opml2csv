package utils

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func IsStructuredLine(line string) bool {
	requiredAttributes := []string{"type=", "text=", "title=", "xmlUrl=", "htmlUrl="}
	for _, attr := range requiredAttributes {
		if !strings.Contains(line, attr) {
			return false
		}
	}
	return true
}

func ParseOpmlLine(line string) (*Outline, error) {
	var outline Outline
	err := xml.Unmarshal([]byte(line), &outline)
	if err != nil {
		return nil, fmt.Errorf("failed to parse OPML line: %w", err)
	}
	return &outline, nil
}
