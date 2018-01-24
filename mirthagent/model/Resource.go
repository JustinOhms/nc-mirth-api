package model

import (
	"encoding/xml"

	"github.com/NavigatingCancer/mirth-api/mirthagent/errors"
	"github.com/caimeo/console"
)

type resourceList struct {
	XMLName   xml.Name   `xml:"list"`
	Resources []Resource `xml:"com.mirth.connect.plugins.directoryresource.DirectoryResourceProperties"`
}

type Resource struct {
	//XMLName         xml.Name `xml:"com.mirth.connect.plugins.directoryresource.DirectoryResourceProperties"`
	Versionø        string `xml:"version,attr"`
	PluginPointName string `xml:"pluginPointName"`
	Type            string `xml:"type"`
	Id              string `xml:"id"`
	Name            string `xml:"name"`
	Description     string `xml:"description"`
	IsGlobal        bool   `xml:"includeWithGlobalScripts"`
	Directory       string `xml:"directory"`
	Recursive       bool   `xml:"directoryRecursion"`
}

func ResourcesFromXml(b []byte) []Resource {
	console.Always(string(b))
	l := resourceList{}
	xml.Unmarshal(b, &l)
	console.Always(l)
	return l.Resources
}

func ResourcesToXml(r []Resource) ([]byte, error) {
	l := resourceList{
		Resources: r,
	}
	x, e := xml.Marshal(l)
	errors.CheckErrorAndLog(e)
	return x, e
}

func NewResource() *Resource {
	r := Resource{
		Versionø:        "3.5.0",
		PluginPointName: "Directory Resource",
		Type:            "Directory",
		Recursive:       true,
		IsGlobal:        true,
	}
	return &r
}

func NewDefaultResource() *Resource {
	r := NewResource()
	r.Id = "Default Resource"
	r.Name = "[Default Resource]"
	return r
}

//	<com.mirth.connect.plugins.directoryresource.DirectoryResourceProperties version="3.5.0">
//		<pluginPointName>Directory Resource</pluginPointName>
//		<type>Directory</type>
//		<id>d7188444-6f80-4d78-9183-4e538260590d</id>
//		<name>Navigating Care Resources</name>
//		<description />
//		<includeWithGlobalScripts>true</includeWithGlobalScripts>
//		<directory>/mnt/mirth/data/javax</directory>
//		<directoryRecursion>true</directoryRecursion>
//	</com.mirth.connect.plugins.directoryresource.DirectoryResourceProperties>
