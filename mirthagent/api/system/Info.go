package system

import (
	"encoding/xml"

	"github.com/NavigatingCancer/mirth-api/mirthagent/f"
)

type Info struct {
	JVMVersion     string `xml:"jvmVersion"`
	OSName         string `xml:"osName"`
	OSVersion      string `xml:"osVersion"`
	OSArchitecture string `xml:"osArchitecture"`
	DBName         string `xml:"dbName"`
	DBVersion      string `xml:"dbVersion"`
}

func InfoXmlParse(b []byte) (m Info) {
	xml.Unmarshal(b, &m)
	return m
}

func (Ω *System) Info() (chan Info, chan error) {
	c := make(chan Info, 1)
	ec := make(chan error, 1)
	go Ω.info(c, ec)
	return c, ec
}

func (Ω *System) info(c chan Info, ec chan error) {
	defer close(c)
	defer close(ec)
	req := Ω.Session.NewRequest().Get(Ω.Session.Paths.System.Info())
	r, b, e := req.EndBytes()
	if f.ResponseOrStatusErrors(ec, r, e, "System info could not be retrieved") {
		return
	}
	c <- InfoXmlParse(b)
}
