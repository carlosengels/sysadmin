package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	maxTTL     = 30    // Maximum number of hops to try
	timeoutSec = 2     // Timeout in seconds for each probe
	maxRetries = 3     // Number of retries per hop
	packetSize = 52    // Standard ICMP echo request size
)

type GeoIP struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	ISP         string  `json:"isp"`
	CountryCode string  `json:"countryCode"`
}

func getIPLocation(ip string) (*GeoIP, error) {
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoIP
	if err := json.Unmarshal(body, &geo); err != nil {
		return nil, err
	}

	// Add a small delay to respect rate limiting
	time.Sleep(100 * time.Millisecond)
	return &geo, nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("Usage: %s <host>\n", os.Args[0])
		os.Exit(1)
	}

	dest := flag.Arg(0)
	destAddr, err := net.ResolveIPAddr("ip4", dest)
	if err != nil {
		fmt.Printf("Could not resolve %s: %v\n", dest, err)
		os.Exit(1)
	}

	fmt.Printf("Tracing route to %s [%s]\n", dest, destAddr)
	fmt.Printf("Maximum hops: %d\n\n", maxTTL)

	for ttl := 1; ttl <= maxTTL; ttl++ {
		addr, rtt, err := probe(destAddr, ttl)
		if err != nil {
			fmt.Printf("%2d  * * *\n", ttl)
			continue
		}

		hostname := ""
		if names, err := net.LookupAddr(addr.String()); err == nil && len(names) > 0 {
			hostname = names[0]
		}

		// Get geolocation data
		geo, err := getIPLocation(addr.String())
		if err != nil || geo.Status != "success" {
			if hostname != "" {
				fmt.Printf("%2d  %s (%s)  %.3f ms\n", ttl, hostname, addr, rtt.Seconds()*1000)
			} else {
				fmt.Printf("%2d  %s  %.3f ms\n", ttl, addr, rtt.Seconds()*1000)
			}
		} else {
			location := fmt.Sprintf("%s, %s (%s)", geo.City, geo.Country, geo.CountryCode)
			if hostname != "" {
				fmt.Printf("%2d  %s (%s)  %.3f ms  [%s - %s]\n", ttl, hostname, addr, rtt.Seconds()*1000, location, geo.ISP)
			} else {
				fmt.Printf("%2d  %s  %.3f ms  [%s - %s]\n", ttl, addr, rtt.Seconds()*1000, location, geo.ISP)
			}
		}

		if addr.String() == destAddr.String() {
			break
		}
	}
}

func probe(dest *net.IPAddr, ttl int) (*net.IPAddr, time.Duration, error) {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	if err := conn.IPv4PacketConn().SetTTL(ttl); err != nil {
		return nil, 0, err
	}

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  ttl,
			Data: make([]byte, packetSize-8),
		},
	}

	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		return nil, 0, err
	}

	start := time.Now()
	if _, err := conn.WriteTo(msgBytes, dest); err != nil {
		return nil, 0, err
	}

	reply := make([]byte, 1500)
	err = conn.SetReadDeadline(time.Now().Add(timeoutSec * time.Second))
	if err != nil {
		return nil, 0, err
	}

	n, peer, err := conn.ReadFrom(reply)
	if err != nil {
		return nil, 0, err
	}

	duration := time.Since(start)
	rm, err := icmp.ParseMessage(1, reply[:n])
	if err != nil {
		return nil, 0, err
	}

	switch rm.Type {
	case ipv4.ICMPTypeTimeExceeded:
		return peer.(*net.IPAddr), duration, nil
	case ipv4.ICMPTypeEchoReply:
		return peer.(*net.IPAddr), duration, nil
	default:
		return nil, 0, fmt.Errorf("unexpected ICMP message type: %v", rm.Type)
	}
}
