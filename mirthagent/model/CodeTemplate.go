package model

import "encoding/xml"

type CodeTemplates struct {
	XMLName   xml.Name       `xml:"list"`
	Templates []CodeTemplate `xml:",innerxml"`
}

type CodeTemplate struct {
	XMLName     xml.Name `xml:"codeTemplate"`
	Versionø    string   `xml:"version,attr"`
	XMLContentø string   `xml:",innerxml"`
}

type CodeLibraries struct {
	XMLName   xml.Name      `xml:"list"`
	Libraries []CodeLibrary `xml:",innerxml"`
}

type CodeLibrary struct {
	XMLName            xml.Name                     `xml:"codeTemplateLibrary"`
	Id                 string                       `xml:"id"`
	Name               string                       `xml:"name"`
	Revision           string                       `xml:"revision"`
	Description        string                       `xml:"description"`
	IncludeNewChannels bool                         `xml:"includeNewChannels"`
	EnabledChannelIds  []string                     `xml:"enabledChannelIds>string"`
	DisabledChannelIds []string                     `xml:"disabledChannelIds>string"`
	CodeTemplateRefs   []CodeLibraryCodeTemplateRef `xml:"codeTemplates>codeTemplate"`
}

type CodeLibraryCodeTemplateRef struct {
	XMLName  xml.Name `xml:"codeTemplate"`
	Id       string   `xml:"id"`
	Versionø string   `xml:"version,attr"`
}
