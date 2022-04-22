// Copyright © 2021 - 2022 vity <vityme@icloud.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package k

import (
	"encoding/xml"
	"io"
	"strings"
)

type xmlMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m xmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for k, v := range m {
		err = e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
		if err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

func (m *xmlMap) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	*m = xmlMap{}
	for {
		var e xmlMapEntry
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func XMLFromMap(m map[string]string, root string) string {
	xmlBuf, _ := xml.Marshal(xmlMap(m))
	return strings.ReplaceAll(string(xmlBuf), "xmlMap", root)
}

func XMLToMap(s string) map[string]string {
	m := make(map[string]string)
	_ = xml.Unmarshal([]byte(s), (*xmlMap)(&m))
	return m
}
