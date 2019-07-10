package utils

import (
	"net"
	"net/http"
)

// GetIP is used to get the IP of the originator of the request
func GetIP(req *http.Request) net.IP {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil
	}

	iporig := net.ParseIP(ip)
	return iporig
}
