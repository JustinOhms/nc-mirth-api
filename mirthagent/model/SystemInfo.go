package model

type SystemInfo struct {
	JVMVersion     string `xml:"jvmVersion"`
	OSName         string `xml:"osName"`
	OSVersion      string `xml:"osVersion"`
	OSArchitecture string `xml:"osArchitecture"`
	DBName         string `xml:"dbName"`
	DBVersion      string `xml:"dbVersion"`
}
