package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

//func (ipAddr IPAddr) String() string {
//	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
//}

func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	for i, val := range ip{
		s[i] = strconv.Itoa(int(val))
	}
	return strings.Join(s, ".")
}

func main() {
	addr := map[string]IPAddr{
		"lookback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addr {
		fmt.Printf("%v: %v\n", n, a)
	}
}
