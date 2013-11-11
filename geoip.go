// Copyright 2013 The Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package geoip

type GeoIPRecord struct {
	CountryCode   string
	CountryCode3  string
	CountryName   string
	Region        string
	City          string
	PostalCode    string
	Latitude      float64
	Longitude     float64
	AreaCode      int
	ContinentCode string
}

//ISO_8859-1 to UTF8
//http://stackoverflow.com/questions/13510458/golang-convert-iso8859-1-to-utf8
func latin1toUTF8(latin1Buf []byte) string {
	buf := make([]rune, len(latin1Buf))
	for i, b := range latin1Buf {
		buf[i] = rune(b)
	}
	return string(buf)
}
