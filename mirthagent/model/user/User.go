package user

import "encoding/xml"

type User struct {
	Id       string `xml:"id"`
	UserName string `xml:"username"`
}

func XmlParse(b []byte) (m User) {
	xml.Unmarshal(b, &m)
	return m
}

type Handler func(User)
