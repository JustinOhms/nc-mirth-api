package model

const defaultApiVersion = "3.5.0"

var _apiVersion string = defaultApiVersion

func apiVersion() string {
	return _apiVersion
}

func SetAPIVersion(v string) {
	_apiVersion = v
}
