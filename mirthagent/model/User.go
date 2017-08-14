package model

import (
	"encoding/xml"
)

type User struct {
	Id       string `xml:"id"`
	UserName string `xml:"username"`
}

func UserFromXml(b []byte) (m User) {
	xml.Unmarshal(b, &m)
	return m
}
