package model

import (
	"encoding/xml"
)

type SystemInfo struct {
	JVMVersion     string `xml:"jvmVersion"`
	OSName         string `xml:"osName"`
	OSVersion      string `xml:"osVersion"`
	OSArchitecture string `xml:"osArchitecture"`
	DBName         string `xml:"dbName"`
	DBVersion      string `xml:"dbVersion"`
}

func SystemInfoFromXml(b []byte) (m SystemInfo) {
	xml.Unmarshal(b, &m)
	return m
}
