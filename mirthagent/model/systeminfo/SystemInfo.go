package systeminfo

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

func XmlParse(b []byte) (m SystemInfo) {
	xml.Unmarshal(b, &m)
	return m
}

type Handler func(SystemInfo)
