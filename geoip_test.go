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

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func init() {
	filepath.Walk("/usr/share/GeoIP/", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".dat") {
			return nil
		}
		if strings.Contains(path, "City") {
			geoIPCity = path
		} else {
			geoIPCountry = path
		}
		return nil
	})
}

var (
	geoIPCountry string
	geoIPCity    string
)

func TestLatin1(t *testing.T) {
	latin1Buf := []uint8{0xf8}
	s := latin1toUTF8(latin1Buf)
	if s != "ø" {
		t.Fatalf("%v", s)
	}
}

func TestGeoIP(t *testing.T) {
	gi, err := Open(geoIPCountry)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCountry)
	}
	defer gi.Delete()

	gicity, err := Open(geoIPCity)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCity)
	}
	defer gicity.Delete()

	ip := net.ParseIP("10.240.21.51")
	if gicity.RecordByIPv4(ip) != nil {
		t.Fatal("1")
	}

	ip = net.ParseIP("196.213.226.36")
	if gi.CountryCodeByIPv4(ip) != "ZA" {
		t.Fatal("2")
	}
	if gi.CountryCode3ByIPv4(ip) != "ZAF" {
		t.Fatal("3")
	}
	if gi.CountryNameByIPv4(ip) != "South Africa" {
		t.Fatal("4")
	}
	rec := gicity.RecordByIPv4(ip)
	if rec.CountryCode != "ZA" {
		t.Fatal("5")
	}
	if rec.CountryCode3 != "ZAF" {
		t.Fatal("6")
	}
	if rec.CountryName != "South Africa" {
		t.Fatal("7")
	}
	if int(rec.Latitude) != -26 {
		t.Fatalf("8: %v", rec.Latitude)
	}
	if int(rec.Longitude) != 28 {
		t.Fatal("9")
	}
	if rec.ContinentCode != "AF" {
		t.Fatal("10")
	}
}

func TestDatabaseCreateTime(t *testing.T) {
	gi, err := Open(geoIPCountry)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCountry)
	}
	defer gi.Delete()
	tt, err := gi.DatabaseCreateTime()
	if err != nil {
		panic(err)
	}
	if !tt.After(time.Date(1900, 10, 4, 0, 0, 0, 0, time.UTC)) || !tt.Before(time.Date(2100, 10, 4, 0, 0, 0, 0, time.UTC)) {
		t.Fatalf("weird date")
	}
}

func TestDatabaseInfo(t *testing.T) {
	gi, err := Open(geoIPCountry)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCountry)
	}
	defer gi.Delete()

	gicity, err := Open(geoIPCity)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCity)
	}
	defer gicity.Delete()

	fmt.Printf("Country\n")
	fmt.Printf("%v\n", gi.DatabaseInfo())
	fmt.Printf("City\n")
	fmt.Printf("%v\n", gicity.DatabaseInfo())
}

func TestDatabaseEdition(t *testing.T) {
	gi, err := Open(geoIPCountry)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCountry)
	}
	defer gi.Delete()

	gicity, err := Open(geoIPCity)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCity)
	}
	defer gicity.Delete()

	fmt.Printf("Country\n")
	fmt.Printf("%v\n", gi.DatabaseEdition())
	fmt.Printf("City\n")
	fmt.Printf("%v\n", gicity.DatabaseEdition())
	if !gi.IsIPv4Database() || !gi.IsCountryDatabase() || gi.IsIPv6Database() || gi.IsCityDatabase() {
		t.Fatalf("not and ip4 && country database")
	}
	if !gicity.IsIPv4Database() || gicity.IsCountryDatabase() || gicity.IsIPv6Database() || !gicity.IsCityDatabase() {
		t.Fatalf("not and ip4 && city database")
	}
}

func TestCountries(t *testing.T) {
	gi, err := Open(geoIPCountry)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCountry)
	}
	defer gi.Delete()

	//www.bbc.co.uk
	ip := net.ParseIP("212.58.246.91")
	fmt.Printf("Country = %v", gi.CountryCodeByIPv4(ip))

	//www.vastech.co.za
	ip = net.ParseIP("196.22.132.6")
	fmt.Printf("Country = %v", gi.CountryCodeByIPv4(ip))

	//www.asb.nl
	c := gi.CountryCodeByIPv4(net.ParseIP("194.247.30.31"))
	fmt.Printf("Country = %v", c)

	//www.asb.or.jp
	c = gi.CountryCodeByIPv4(net.ParseIP("61.197.168.228"))
	fmt.Printf("Country = %v", c)

	//guialocal.com.ar
	c = gi.CountryCodeByIPv4(net.ParseIP("200.80.35.138"))
	fmt.Printf("Country = %v", c)

	//www.asb.fr
	c = gi.CountryCodeByIPv4(net.ParseIP("62.73.5.199"))
	fmt.Printf("Country = %v", c)

	//www.asb.com.br
	c = gi.CountryCodeByIPv4(net.ParseIP("200.255.40.148"))
	fmt.Printf("Country = %v", c)

	//india.gov.in
	c = gi.CountryCodeByIPv4(net.ParseIP("164.100.56.201"))
	fmt.Printf("Country = %v", c)

	//www.deutschland.de
	c = gi.CountryCodeByIPv4(net.ParseIP("109.239.60.76"))
	fmt.Printf("Country = %v", c)

	//www.sweden.se
	c = gi.CountryCodeByIPv4(net.ParseIP("217.114.81.7"))
	fmt.Printf("Country = %v", c)
}

func TestBug14019(t *testing.T) {
	gi, err := Open(geoIPCity)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCity)
	}
	defer gi.Delete()
	ipnum := binary.BigEndian.Uint32(net.ParseIP("83.206.228.217").To4())
	c := gi.CityByIPNum(ipnum)
	city := "Le Kremlin-bicêtre"
	if c != city {
		t.Fatalf("Cities are encoded with Latin1")
	}
	c2 := gi.RecordByIPNum(ipnum)
	if c2.City != city {
		t.Fatalf("Cities are encoded with Latin1")
	}
}

func TestBug14030(t *testing.T) {
	gi, err := Open(geoIPCity)
	if err != nil {
		t.Fatalf("Open(%v) failed", geoIPCity)
	}
	defer gi.Delete()
	ipnum := binary.BigEndian.Uint32(net.ParseIP("83.206.228.217").To4())
	city := "Le Kremlin-bicêtre"
	done := make(chan int, 100)
	for i := 0; i < 100; i++ {
		go func() {
			c := gi.CityByIPNum(ipnum)
			if c != city {
				t.Fatalf("Cities are encoded with Latin1")
			}
			done <- 1
		}()
	}
	count := 0
	for count < 100 {
		<-done
		count++
	}
}
