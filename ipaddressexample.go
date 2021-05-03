package main

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"log"
)

func main() {
	_, ipna, _ := iplib.ParseCIDR("192.168.4.0/22")
	fmt.Println(ipna.Subnet(24))   // []iplib.Net{ 192.168.4.0/24, 192.168.5.0/24,
	//              192.168.6.0/24, 192.168.7.0/24 }
	ipnb, _ := ipna.Supernet(21) // 192.168.0.0/21
	fmt.Println(ipnb.String())
	ipnc := ipna.PreviousNet(21)   // 192.168.0.0/21
	fmt.Println(ipnc.String())
	ipnd := ipna.NextNet(21)       // 192.168.8.0/21
	fmt.Println(ipnd.String())

	transitnet, transitpool, err := iplib.ParseCIDR("10.0.0.0/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transitnet.To4())
	fmt.Println(transitpool.String())
	ip1 := transitpool.NextNet(31)
	fmt.Println(ip1.String())
	ip2 := transitpool.NextNet(31)
	fmt.Println(ip2.String())
	poolarray := transitpool.Enumerate(31, 0)
	fmt.Println(poolarray)
	fmt.Println("Enumerated list")
	for _, value := range poolarray{
		fmt.Println(value)
	}
	// New net from existing, not sure if useful yet, subnet is next
	ip3 := iplib.NewNet(transitpool.FirstAddress(), 31)
	fmt.Println(ip3.String())

	// Use this to subnet a /24 to /31
	ip4, _ := transitpool.Subnet(31)
	for _, network := range ip4{
		fmt.Println(network.String())
		fmt.Println(network.FirstAddress(), ", ", network.LastAddress())
		fmt.Println(fmt.Sprintf("%d.%d.%d.%d", network.Mask[0], network.Mask[1], network.Mask[2], network.Mask[3]))

	}
}
