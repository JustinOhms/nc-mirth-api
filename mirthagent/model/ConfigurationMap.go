package model

import "encoding/xml"

type ConfigurationMap struct {
	XMLName xml.Name         `xml:"map"`
	Entries []ConfigMapEntry `xml:",innerxml"`
}

type ConfigMapEntry struct {
	XMLName xml.Name `xml:"entry"`
	Key     string   `xml:"string"`
	Value   string   `xml:"com.mirth.connect.util.ConfigurationProperty>value"`
	Comment string   `xml:"com.mirth.connect.util.ConfigurationProperty>comment"`
}

/*
<map>
	<entry>
		<string>bi_api_host</string>
		<com.mirth.connect.util.ConfigurationProperty>
			<value>bi-api-staging.navigatingcare.com</value>
		</com.mirth.connect.util.ConfigurationProperty>
	</entry>
	<entry>
		<string>environment</string>
		<com.mirth.connect.util.ConfigurationProperty>
			<value>local</value>
			<comment>which environment we are running in local/staging/production</comment>
		</com.mirth.connect.util.ConfigurationProperty>
	</entry>
</map>
*/
