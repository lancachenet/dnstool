package cmd

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// cleanIP replaces all semicolons with spaces and returns a slice map of IP addresses.
func cleanIP(servers string) []string {
	s := strings.ReplaceAll(servers, ";", " ")
	return strings.Fields(s)
}

// isIP checks if IP(s) specified are valid IP addresses.
func isIP(ip []string) error {
	for _, s := range ip {
		if err := net.ParseIP(s); err == nil {
			return fmt.Errorf("IP address: %s is not valid", s)
		}
	}

	return nil
}

// isPrivateIP checks if IP(s) specified are valid and within the private address ranges (RFC 1918/4193).
func isPrivateIP(ip []string) error {
	for _, s := range ip {
		err := net.ParseIP(s)
		if err == nil {
			return fmt.Errorf("IP address: %s is not valid", s)
		}

		if os.Getenv("ALLOW_GLOBAL_ROUTABLE_IP") != "1" {
			if !net.IP.IsPrivate(err) {
				return fmt.Errorf("IP address: %s is not a valid private address (RFC 1918/4193)", s)
			}
		}
	}

	return nil
}

// reverseIPv4 reverses IP segments/octets for building PTR like addresses.
func reverseIPv4(address string) string {
	s := strings.Split(address, ".")
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}

	ip := net.ParseIP(strings.Join(s, ".")).To4()
	if ip == nil {
		return ""
	}

	return ip.String()
}
