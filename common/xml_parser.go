package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

type XmlDate time.Time

type SessionInfo struct {
	XMLName  xml.Name         `xml:"sessionInfo"`
	TimeZone string           `xml:"timeZone,attr"`
	Features []SessionFeature `xml:"features>feature"`
}

type SessionFeature struct {
	ID         uint32  `xml:"id,attr"`
	Name       string  `xml:"name,attr"`
	Type       string  `xml:"type,attr"`
	Value      string  `xml:"value,attr"`
	EndDate    XmlDate `xml:"endDate,attr"`
	UserNumber uint32  `xml:"userNumber,attr,omitempty"`
}

// UnmarshalText For parsing the malformed UTC datetime in xml
//  1. We must define an alias(XmlDate) for time.Time
//  2. To implement UnmarshalText for XmlDate
func (d *XmlDate) UnmarshalText(text []byte) error {
	t, err := time.Parse(dateFormat, string(text))
	if err != nil {
		return err
	}
	*d = XmlDate(t)
	return nil
}

func main() {
	xmlData := `
<?xml version="1.0" encoding="UTF-8"?>
<sessionInfo timeZone="+08:00">
  <features>
    <feature id="15" name="FEAT_A" type="ro" value="1" endDate="2025-12-31 15:59:59" userNumber="66"/>
    <feature id="18" name="FEAT_E" type="ro" value="0" endDate="2025-11-30 15:59:59" userNumber="66"/>
    <feature id="13" name="FEAT_B" type="ro" value="1" endDate="2025-10-31 15:59:59"/>
    <feature id="11" name="FEAT_D" type="ro" value="0" endDate="2025-12-24 15:59:59"/>
  </features>
</sessionInfo>
`

	var si SessionInfo
	err := xml.Unmarshal([]byte(xmlData), &si)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 id 到 SessionFeature 的映射
	idFeatureMap := make(map[uint32]SessionFeature)
	for _, feature := range si.Features {
		idFeatureMap[feature.ID] = feature
	}

	// 打印结果
	fmt.Printf("Session Name: %s\n", si.XMLName)
	fmt.Printf("Session TimeZone: %s\n", si.TimeZone)
	fmt.Println("Features:")
	for id, feature := range idFeatureMap {
		fmt.Printf("    ID: %d\n", id)
		fmt.Printf("    Name: %s\n", feature.Name)
		fmt.Printf("    Type: %s\n", feature.Type)
		fmt.Printf("    Value: %s\n", feature.Value)
		fmt.Printf("    EndDate: %s\n", time.Time(feature.EndDate).Format(dateFormat))
		fmt.Printf("    UserNumber: %d\n", feature.UserNumber)
		fmt.Println()
	}
}
