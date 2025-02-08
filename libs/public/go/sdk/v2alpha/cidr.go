package sdkv2alphalib

import (
	"errors"
	"fmt"
	"math"
	"net"

	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
)

const reservedIPCount = 10

// CIDRBlock represents a wrapper for the CIDR data structure, including associated network IP information.
type CIDRBlock struct {
	*typev2pb.CIDR

	netIP *net.IP
}

// NewCIDR creates a CIDRBlock instance by parsing the given CIDR string and calculating associated IP information.
// Returns a pointer to CIDRBlock and an error if the parsing fails.
func NewCIDR(cidr string) (*CIDRBlock, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	ones, bits := ipnet.Mask.Size()

	totalIPs := 1 << (bits - ones)

	// Reserve the following IPs
	var ips []string
	for i := 0; i < reservedIPCount; i++ {
		ips = append(ips, ip.String())
		incrementIP(ip)
	}

	// Prevent G115: integer overflow conversion int -> uint32 (gosec)
	if ones < 0 || ones > math.MaxUint32 {
		fmt.Println("Error: ones value is out of uint32 range")
		return nil, errors.New("error: ones value is out of uint32 range")
	}
	subnetBits := uint32(ones)

	// Prevent G115: integer overflow conversion int -> uint32 (gosec)
	if bits < 0 || bits > math.MaxUint32 {
		fmt.Println("Error: bits value is out of uint32 range")
		return nil, errors.New("error: bits value is out of uint32 range")
	}
	totalBits := uint32(bits)

	// Prevent G115: integer overflow conversion int -> uint64 (gosec)
	if totalIPs < 0 {
		fmt.Println("Error: totalIPs is negative, cannot convert to uint64")
		return nil, errors.New("error: totalIPs is negative, cannot convert to uint64")
	}
	totalIps := uint64(totalIPs) // Safe conversion

	return &CIDRBlock{
		CIDR: &typev2pb.CIDR{
			CidrNotation: ipnet.String(),
			FirstIp:      ip.String(),
			SubnetMask:   ipnet.Mask.String(),
			SubnetBits:   subnetBits,
			TotalIps:     totalIps,
			TotalBits:    totalBits,
			ReservedIps:  ips,
		},
		netIP: &ip,
	}, nil
}

// GetNthIP calculates and returns the nth IP address within the given CIDR block.
// It takes a CIDR string and an integer n, representing the desired offset from the first IP address in the block.
// If successful, it returns the nth IP address as a string and no error; otherwise, it returns an error.
func (c *CIDRBlock) GetNthIP(n int) (string, error) {
	ip := *c.netIP
	for i := 0; i < n; i++ {
		incrementIP(ip)
	}

	return ip.String(), nil
}

// incrementIP increments the given IP address by one, accounting for carry-over across byte boundaries.
func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
