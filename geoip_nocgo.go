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

// +build !cgo

package geoip

import (
	"errors"
	"net"
	"time"
)

func (gi *GeoIP) IsIPv4Database() bool {
	panic("geoip needs cgo")
}

func (gi *GeoIP) IsIPv6Database() bool {
	panic("geoip needs cgo")
}

func (gi *GeoIP) IsCountryDatabase() bool {
	panic("geoip needs cgo")
}

func (gi *GeoIP) IsCityDatabase() bool {
	panic("geoip needs cgo")
}

type GeoIP struct {
}

func New() (gi *GeoIP, err error) {
	return nil, errors.New("geoip needs cgo")
}

func Open(filename string) (gi *GeoIP, err error) {
	return nil, errors.New("geoip needs cgo")
}

func (gi *GeoIP) DatabaseInfo() string {
	panic("geoip needs cgo")
}

func (gi *GeoIP) DatabaseEdition() (code int) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) DatabaseCreateTime() (time.Time, error) {
	return time.Now(), errors.New("geoip needs cgo")
}

func (gi *GeoIP) DbAvail(typ int) (avail bool) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryCodeByIPv4(ip net.IP) (code string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryCodeByIPNum(ipnum uint32) (code string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryCodeByIPv6(ip net.IP) (code string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryCode3ByIPv4(ip net.IP) (code3 string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryCode3ByIPNum(ipnum uint32) (code3 string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryNameByIPv4(ip net.IP) (name string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CountryNameByIPNum(ipnum uint32) (name string) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) RecordByIPv4(ip net.IP) (gir *GeoIPRecord) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) RecordByIPNum(ipnum uint32) (gir *GeoIPRecord) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) CityByIPNum(ipnum uint32) string {
	panic("geoip needs cgo")
}

func (gi *GeoIP) RecordByIPv6(ip net.IP) (gir *GeoIPRecord) {
	panic("geoip needs cgo")
}

func (gi *GeoIP) Delete() {
	panic("geoip needs cgo")
}

func CodeByID(id int) (code string) {
	panic("geoip needs cgo")
}

func Code3ByID(id int) (code3 string) {
	panic("geoip needs cgo")
}

func NameByID(id int) (name string) {
	panic("geoip needs cgo")
}

func ContinentByID(id int) (continent string) {
	panic("geoip needs cgo")
}
