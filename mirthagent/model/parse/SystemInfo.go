package parse

import (
	"encoding/xml"

	"github.com/NavigatingCancer/mirth-api/mirthagent/model"
)

func SystemInfo(b []byte) (m model.SystemInfo) {
	xml.Unmarshal(b, &m)
	return m
}
