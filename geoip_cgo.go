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

// +build cgo

package geoip

//#include "GeoIP.h"
//#include "GeoIPCity.h"
//#include <stdio.h>
//#cgo LDFLAGS: -lGeoIP
import "C"

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

const (
	COUNTRY_EDITION                    = C.GEOIP_COUNTRY_EDITION
	CITY_EDITION_REV1                  = C.GEOIP_CITY_EDITION_REV1
	REGION_EDITION_REV1                = C.GEOIP_REGION_EDITION_REV1
	ISP_EDITION                        = C.GEOIP_ISP_EDITION
	ORG_EDITION                        = C.GEOIP_ORG_EDITION
	CITY_EDTION                        = C.GEOIP_CITY_EDITION_REV0
	REGTION_EDITION_REV0               = C.GEOIP_REGION_EDITION_REV0
	PROXY_EDTION                       = C.GEOIP_PROXY_EDITION
	ASNUM_EDITION                      = C.GEOIP_ASNUM_EDITION
	NETSPEED_EDITION                   = C.GEOIP_NETSPEED_EDITION
	DOMAIN_EDITION                     = C.GEOIP_DOMAIN_EDITION
	COUNTRY_EDITION_V6                 = C.GEOIP_COUNTRY_EDITION_V6
	LOCATIONA_EDITION                  = C.GEOIP_LOCATIONA_EDITION
	ACCURACYRADIUS_EDITION             = C.GEOIP_ACCURACYRADIUS_EDITION
	CITYCONFIDENCE_EDITION             = C.GEOIP_CITYCONFIDENCE_EDITION
	CITYCONFIDENCEDIST_EDITION         = C.GEOIP_CITYCONFIDENCEDIST_EDITION
	LARGE_COUNTRY_EDITION              = C.GEOIP_LARGE_COUNTRY_EDITION
	LARGE_COUNTRY_EDITION_V6           = C.GEOIP_LARGE_COUNTRY_EDITION_V6
	CITYCONFIDENCEDIST_ISP_ORG_EDITION = C.GEOIP_CITYCONFIDENCEDIST_ISP_ORG_EDITION
	CCM_COUNTRY_EDITION                = C.GEOIP_CCM_COUNTRY_EDITION
	ASNUM_EDITION_V6                   = C.GEOIP_ASNUM_EDITION_V6
	ISP_EDITION_V6                     = C.GEOIP_ISP_EDITION_V6
	ORG_EDITION_V6                     = C.GEOIP_ORG_EDITION_V6
	DOMAIN_EDITION_V6                  = C.GEOIP_DOMAIN_EDITION_V6
	LOCATIONA_EDITION_V6               = C.GEOIP_LOCATIONA_EDITION_V6
	REGISTRAR_EDITION                  = C.GEOIP_REGISTRAR_EDITION
	REGISTRAR_EDITION_V6               = C.GEOIP_REGISTRAR_EDITION_V6
	USERTYPE_EDITION                   = C.GEOIP_USERTYPE_EDITION
	USERTYPE_EDITION_V6                = C.GEOIP_USERTYPE_EDITION_V6
	CITY_EDITION_REV1_V6               = C.GEOIP_CITY_EDITION_REV1_V6
	CITY_EDITION_REV0_V6               = C.GEOIP_CITY_EDITION_REV0_V6
	NETSPEED_EDITION_REV1              = C.GEOIP_NETSPEED_EDITION_REV1
	NETSPEED_EDITION_REV1_V6           = C.GEOIP_NETSPEED_EDITION_REV1_V6
)

func (gi *GeoIP) IsIPv4Database() bool {
	switch gi.edition {
	case COUNTRY_EDITION,
		CITY_EDITION_REV1,
		REGION_EDITION_REV1,
		ISP_EDITION,
		ORG_EDITION,
		CITY_EDTION,
		REGTION_EDITION_REV0,
		PROXY_EDTION,
		ASNUM_EDITION,
		NETSPEED_EDITION,
		DOMAIN_EDITION,
		LOCATIONA_EDITION,
		ACCURACYRADIUS_EDITION,
		CITYCONFIDENCE_EDITION,
		CITYCONFIDENCEDIST_EDITION,
		LARGE_COUNTRY_EDITION,
		CITYCONFIDENCEDIST_ISP_ORG_EDITION,
		CCM_COUNTRY_EDITION,
		NETSPEED_EDITION_REV1,
		REGISTRAR_EDITION,
		USERTYPE_EDITION:
		return true
	}
	return false
}

func (gi *GeoIP) IsIPv6Database() bool {
	switch gi.edition {
	case COUNTRY_EDITION_V6,
		LARGE_COUNTRY_EDITION_V6,
		ASNUM_EDITION_V6,
		ISP_EDITION_V6,
		ORG_EDITION_V6,
		DOMAIN_EDITION_V6,
		LOCATIONA_EDITION_V6,
		REGISTRAR_EDITION_V6,
		USERTYPE_EDITION_V6,
		CITY_EDITION_REV1_V6,
		CITY_EDITION_REV0_V6,
		NETSPEED_EDITION_REV1_V6:
		return true
	}
	return false
}

func (gi *GeoIP) IsCountryDatabase() bool {
	switch gi.edition {
	case COUNTRY_EDITION,
		COUNTRY_EDITION_V6,
		LARGE_COUNTRY_EDITION,
		LARGE_COUNTRY_EDITION_V6,
		CCM_COUNTRY_EDITION:
		return true
	}
	return false
}

func (gi *GeoIP) IsCityDatabase() bool {
	switch gi.edition {
	case CITY_EDITION_REV1,
		CITY_EDTION,
		CITYCONFIDENCE_EDITION,
		CITYCONFIDENCEDIST_EDITION,
		CITYCONFIDENCEDIST_ISP_ORG_EDITION,
		CITY_EDITION_REV1_V6,
		CITY_EDITION_REV0_V6:
		return true
	}
	return false
}

type GeoIP struct {
	gi      *C.GeoIP
	edition int
}

func checkedCString(goVal string) *C.char {
	cVal := C.CString(goVal)
	if cVal == nil {
		panic(syscall.ENOMEM)
	}
	return cVal
}

func New() (gi *GeoIP, err error) {
	gi = new(GeoIP)
	gi.gi = C.GeoIP_new(C.GEOIP_MEMORY_CACHE)
	if gi.gi == nil {
		err = errors.New("GeoIP_new failed")
		return
	}
	return
}

func Open(filename string) (gi *GeoIP, err error) {
	cFilename := checkedCString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	gi = new(GeoIP)
	gi.gi = C.GeoIP_open(cFilename, C.GEOIP_MEMORY_CACHE)
	if gi.gi == nil {
		err = errors.New("GeoIP_open failed")
		return
	}
	gi.edition = gi.DatabaseEdition()
	return
}

func (gi *GeoIP) DatabaseInfo() string {
	// this call returns a newly allocated CString
	info := C.GeoIP_database_info(gi.gi)
	defer C.free(unsafe.Pointer(info))
	goinfo := C.GoString(info)
	return goinfo
}

func (gi *GeoIP) DatabaseEdition() (code int) {
	return int(C.GeoIP_database_edition(gi.gi))
}

func (gi *GeoIP) DatabaseCreateTime() (time.Time, error) {
	info := gi.DatabaseInfo()
	infos := strings.Split(info, " ")
	if len(infos) < 2 {
		return time.Time{}, errors.New(fmt.Sprintf("unable to parse time when database info = %v", info))
	}
	if len(infos[1]) != 8 {
		return time.Time{}, errors.New(fmt.Sprintf("unable to parse date when database info = %v", info))
	}
	year, err := strconv.Atoi(infos[1][0:4])
	if err != nil {
		return time.Time{}, err
	}
	month, err := strconv.Atoi(infos[1][4:6])
	if err != nil {
		return time.Time{}, err
	}
	day, err := strconv.Atoi(infos[1][6:8])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func (gi *GeoIP) DbAvail(typ int) (avail bool) {
	return C.GeoIP_db_avail(C.int(typ)) == 1
}

func (gi *GeoIP) CountryCodeByIPv4(ip net.IP) (code string) {
	return gi.CountryCodeByIPNum(binary.BigEndian.Uint32(ip.To4()))
}

func (gi *GeoIP) CountryCodeByIPNum(ipnum uint32) (code string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_country_code_by_ipnum(gi.gi, C.ulong(ipnum)))
}

func (gi *GeoIP) CountryCodeByIPv6(ip net.IP) (code string) {
	cip := checkedCString(ip.String())
	defer C.free(unsafe.Pointer(cip))
	// this call returns a static CString
	return C.GoString(C.GeoIP_country_code_by_name_v6(gi.gi, cip))
}

func (gi *GeoIP) CountryCode3ByIPv4(ip net.IP) (code3 string) {
	return gi.CountryCode3ByIPNum(binary.BigEndian.Uint32(ip.To4()))
}

func (gi *GeoIP) CountryCode3ByIPNum(ipnum uint32) (code3 string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_country_code3_by_ipnum(gi.gi, C.ulong(ipnum)))
}

func (gi *GeoIP) CountryNameByIPv4(ip net.IP) (name string) {
	return gi.CountryNameByIPNum(binary.BigEndian.Uint32(ip.To4()))
}

func (gi *GeoIP) CountryNameByIPNum(ipnum uint32) (name string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_country_name_by_ipnum(gi.gi, C.ulong(ipnum)))
}

func (gi *GeoIP) RecordByIPv4(ip net.IP) (gir *GeoIPRecord) {
	return gi.RecordByIPNum(binary.BigEndian.Uint32(ip.To4()))
}

func (gi *GeoIP) RecordByIPNum(ipnum uint32) (gir *GeoIPRecord) {
	cGir := C.GeoIP_record_by_ipnum(gi.gi, C.ulong(ipnum))
	if cGir == nil {
		return nil
	}
	// this call frees all the CStrings in cGir
	defer C.GeoIPRecord_delete(cGir)
	gir = new(GeoIPRecord)
	gir.CountryCode = C.GoString(cGir.country_code)
	gir.CountryCode3 = C.GoString(cGir.country_code3)
	gir.CountryName = C.GoString(cGir.country_name)
	gir.Region = C.GoString(cGir.region)
	gir.City = latin1toUTF8([]byte(C.GoString(cGir.city)))
	gir.PostalCode = C.GoString(cGir.postal_code)
	gir.Latitude = float64(cGir.latitude)
	gir.Longitude = float64(cGir.longitude)
	gir.AreaCode = int(cGir.area_code)
	gir.ContinentCode = C.GoString(cGir.continent_code)
	return
}

func (gi *GeoIP) CityByIPNum(ipnum uint32) string {
	cGir := C.GeoIP_record_by_ipnum(gi.gi, C.ulong(ipnum))
	if cGir == nil {
		return ""
	}
	// this call frees all the CStrings in cGir
	defer C.GeoIPRecord_delete(cGir)
	return latin1toUTF8([]byte(C.GoString(cGir.city)))
}

func (gi *GeoIP) RecordByIPv6(ip net.IP) (gir *GeoIPRecord) {
	cip := checkedCString(ip.String())
	defer C.free(unsafe.Pointer(cip))
	cGir := C.GeoIP_record_by_addr_v6(gi.gi, cip)
	if cGir == nil {
		return nil
	}
	// this call frees all the CStrings in cGir
	defer C.GeoIPRecord_delete(cGir)
	gir = new(GeoIPRecord)
	gir.CountryCode = C.GoString(cGir.country_code)
	gir.CountryCode3 = C.GoString(cGir.country_code3)
	gir.CountryName = C.GoString(cGir.country_name)
	gir.Region = C.GoString(cGir.region)
	gir.City = latin1toUTF8([]byte(C.GoString(cGir.city)))
	gir.PostalCode = C.GoString(cGir.postal_code)
	gir.Latitude = float64(cGir.latitude)
	gir.Longitude = float64(cGir.longitude)
	gir.AreaCode = int(cGir.area_code)
	gir.ContinentCode = C.GoString(cGir.continent_code)
	return
}

func (gi *GeoIP) Delete() {
	C.GeoIP_delete(gi.gi)
}

func CodeByID(id int) (code string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_code_by_id(C.int(id)))
}

func Code3ByID(id int) (code3 string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_code3_by_id(C.int(id)))
}

func NameByID(id int) (name string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_name_by_id(C.int(id)))
}

func ContinentByID(id int) (continent string) {
	// this call returns a static CString
	return C.GoString(C.GeoIP_continent_by_id(C.int(id)))
}
