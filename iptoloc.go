/*
Package iptoloc is a API client to use http://ip-api.com/ for mapping an IP Address to
a Geographic Location.

The MIT License (MIT)

Copyright (c) 2015 RaviTeja

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.*/
package iptoloc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

//APIEndpoint - contains the API URL.
const (
	APIEndpoint = "http://ip-api.com/json/"
)

//GeoResponse - To capture the response from the APIEndpoint
type GeoResponse struct {
	AS          string  `json:"as"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	ISP         string  `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

//GetLoc - Takes an IP Address and returns a string, which contains City and Country
func GetLoc(ipaddr string) string {
	if isValidIP(ipaddr) {
		resp, err := http.Get(APIEndpoint + ipaddr)
		if err != nil {
			fmt.Println("Something wrong with contacting the API")
			os.Exit(1)
		}
		resp2, readerr := ioutil.ReadAll(resp.Body)
		if readerr != nil {
			fmt.Println("Something wrong with reading the response")
			os.Exit(1)
		}
		var res GeoResponse
		jsonerr := json.Unmarshal(resp2, &res)
		if jsonerr != nil {
			fmt.Println("Something wrong with unmarshalling the response")
			os.Exit(1)
		}
		if res.City == "" && res.Country == "" {
			return "<unknown-IP>"
		}
		return res.City + ", " + res.Country
	}
	return "Invalid IP Address"
}

//isValidIP - Takes a IP Address as string and returns true if it is valid or false otherwise
func isValidIP(ipaddr string) bool {
	//net.ParseIP returns nil if the parsed IP is not valid
	ip := net.ParseIP(ipaddr)
	return ip != nil
}
