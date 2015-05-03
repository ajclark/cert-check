package main

import (
	"./cert"
	"flag"
	"fmt"
	"os"
)

var numDays int
var domain string
var verbose bool

func init() {
	flag.IntVar(&numDays, "n", 30, "Number of days to alert on. Only print if this is triggered.")
	flag.StringVar(&domain, "d", "google.com", "domain name to alert on")
	flag.BoolVar(&verbose, "v", false, "always print number of days until expiration")
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func main() {
	expireDays := cert.CheckDate(domain, numDays)
	if expireDays <= numDays {
		fmt.Printf("domain: %s, expires in %d days\n", domain, expireDays)
		os.Exit(1)
	}
	if verbose {
		fmt.Printf("domain: %s, expires in %d days\n", domain, expireDays)
	}
}
