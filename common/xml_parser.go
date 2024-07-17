package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type SessionInfo struct {
	XMLName  xml.Name         `xml:"sessionInfo"`
	TimeZone string           `xml:"timeZone,attr"`
	Features []SessionFeature `xml:"features>feature"`
}

type SessionFeature struct {
	ID         uint32 `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	Type       string `xml:"type,attr"`
	Value      string `xml:"value,attr"`
	EndDate    string `xml:"endDate,attr"`
	UserNumber uint32 `xml:"userNumber,attr,omitempty"`
}

func main() {
	xmlData := `
<?xml version="1.0" encoding="UTF-8"?>
<sessionInfo timeZone="+08:00">
  <features>
    <feature id="15" name="BPM" type="ro" value="1" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="18" name="SDK_CIRCUIT" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="1" name="PASSIVE_EME" type="ro" value="1" endDate="2025-12-31 15:59:59"/>
    <feature id="11" name="PASSIVE_FDTD_GPU" type="ro" value="0" endDate="2025-12-31 15:59:59"/>
    <feature id="16" name="RCWA" type="ro" value="1" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="3" name="FDTD" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="12" name="CIRCUIT" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="3" name="PASSIVE_FDTD" type="ro" value="1" endDate="2025-12-31 15:59:59"/>
    <feature id="23" name="SDK_RCWA" type="ro" value="0" endDate="2025-12-31 15:59:59"/>
    <feature id="20" name="SDK_EME" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="10" name="ACTIVE_OEDEVICE" type="ro" value="1" endDate="2025-12-31 15:59:59"/>
    <feature id="1" name="EME" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="17" name="SDK_BPM" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="22" name="SDK_FDTD_GPU" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="11" name="FDTD_GPU" type="ro" value="1" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="21" name="SDK_FDTD" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="14" name="DDM" type="ro" value="1" endDate="2025-12-31 15:59:59" userNumber="198"/>
    <feature id="19" name="SDK_DDM" type="ro" value="0" endDate="2025-12-31 15:59:59" userNumber="198"/>
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
		fmt.Printf("ID: %d\n", id)
		fmt.Printf("  Name: %s\n", feature.Name)
		fmt.Printf("  Type: %s\n", feature.Type)
		fmt.Printf("  Value: %s\n", feature.Value)
		fmt.Printf("  EndDate: %s\n", feature.EndDate)
		fmt.Printf("  UserNumber: %d\n", feature.UserNumber)
		fmt.Println()
	}
}
