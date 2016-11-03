package interfaces

import "fmt"
import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipaddr IPAddr) String() string {
	return strconv.Itoa((int)(ipaddr[0])) + "." + strconv.Itoa((int)(ipaddr[1])) + "." + strconv.Itoa((int)(ipaddr[2])) + "." + strconv.Itoa((int)(ipaddr[3]))
}

func StringerDemo() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	//when IPAddr implement String method, the Printf will call it to generate the output.
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
