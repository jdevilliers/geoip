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
	"math/rand"
	"net"
	"strconv"
	"testing"
)

func generateIPv4String() string {
	bs := generateBytes(4)
	str := strconv.FormatUint(uint64(bs[0]), 10)
	for _, b := range bs[1:] {
		str += "."
		str += strconv.FormatUint(uint64(b), 10)
	}
	return str
}

func generateBytes(n int) (b []byte) {
	b = make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(255))
	}
	return
}

func getDBv4() (gi *GeoIP, gicity *GeoIP) {
	gi, err := Open(geoIPCountry)
	if err != nil {
		panic("Open(GeoIP.dat) failed")
	}

	gicity, err = Open(geoIPCity)
	if err != nil {
		panic("Open(GeoLiteCity.dat) failed")
	}
	return
}

func BenchmarkIPv4ParseOnly(b *testing.B) {
	ips := []string{}
	for i := 0; i < b.N; i++ {
		ips = append(ips, generateIPv4String())
	}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		net.ParseIP(ips[i])
	}
	b.StopTimer()
}

func BenchmarkIPv4ParseCountry(b *testing.B) {
	gi, gicity := getDBv4()
	defer gi.Delete()
	defer gicity.Delete()
	ips := []string{}
	for i := 0; i < b.N; i++ {
		ips = append(ips, generateIPv4String())
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ip4 := net.ParseIP(ips[i])
		gi.CountryCodeByIPv4(ip4)
	}
	b.StopTimer()
}

func BenchmarkIPv4NoParseCountry(b *testing.B) {
	gi, gicity := getDBv4()
	defer gi.Delete()
	defer gicity.Delete()
	ipbs := []uint32{}
	for i := 0; i < b.N; i++ {
		bs := generateBytes(4)
		ipb := binary.BigEndian.Uint32(bs)
		ipbs = append(ipbs, ipb)
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gi.CountryCodeByIPNum(ipbs[i])
	}
	b.StopTimer()
}

func BenchmarkIPv4ParseCity(b *testing.B) {
	gi, gicity := getDBv4()
	defer gi.Delete()
	defer gicity.Delete()

	ips := []string{}
	for i := 0; i < b.N; i++ {
		ips = append(ips, generateIPv4String())
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ip4 := net.ParseIP(ips[i])
		gicity.RecordByIPv4(ip4)
	}
	b.StopTimer()
}

func BenchmarkIPv4NoParseCityRecord(b *testing.B) {
	gi, gicity := getDBv4()
	defer gi.Delete()
	defer gicity.Delete()

	ipbs := []uint32{}
	for i := 0; i < b.N; i++ {
		bs := generateBytes(4)
		ipb := binary.BigEndian.Uint32(bs)
		ipbs = append(ipbs, ipb)
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gicity.RecordByIPNum(ipbs[i])
	}
	b.StopTimer()
}

func BenchmarkIPv4NoParseCityNoRecord(b *testing.B) {
	gi, gicity := getDBv4()
	defer gi.Delete()
	defer gicity.Delete()

	ipbs := []uint32{}
	for i := 0; i < b.N; i++ {
		bs := generateBytes(4)
		ipb := binary.BigEndian.Uint32(bs)
		ipbs = append(ipbs, ipb)
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		gicity.CityByIPNum(ipbs[i])
	}
	b.StopTimer()
}
