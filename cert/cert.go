package cert

import (
	"crypto/tls"
	"net"
	"time"
)

func CheckDate(domain string, numDays int) (days int) {
	// Set connect timeout to 2 seconds
	ipConn, err := net.DialTimeout("tcp", domain+":443", 2000*time.Millisecond)
	if err != nil {
		panic(err)
	} else {
		// Make TLS connection if TCP connection worked
		config := tls.Config{ServerName: domain}
		conn := tls.Client(ipConn, &config)
		defer conn.Close()

		// Validate TLS handshake
		err := conn.Handshake()
		if err != nil {
			panic(err)
		} else {

			// Print out certificate information for good domains only
			state := conn.ConnectionState()
			v := state.PeerCertificates

			// Calculate the number of days remaining before certificate expires
			expireDate := v[0].NotAfter.Sub(v[0].NotBefore)
			expireDays := int(expireDate.Hours() / 24)
			return expireDays
		}
	}
}
