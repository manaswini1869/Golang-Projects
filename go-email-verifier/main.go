package main

import (
	"bufio"   // Importing bufio package for buffered I/O
	"fmt"     // Importing fmt package for formatted I/O
	"log"     // Importing log package for logging
	"net"     // Importing net package for network operations
	"os"      // Importing os package for operating system functionality
	"strings" // Importing strings package for string manipulation
)

// checkDomain checks various DNS records for a given domain and prints the results
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Performing MX record lookup
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Performing TXT record lookup
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println("Warning: TXT lookup failed for %v\n", err)
	}

	// Checking for SPF record
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			spfRecord = record
			hasSPF = true
			break
		}
	}

	// Performing DMARC record lookup
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println("Warning: DMARC lookup failed for %v\n", err)
	}

	// Checking for DMARC record
	for _, record := range dmarcRecords {
		hasDMARC = true
		dmarcRecord += "\n" + record
		break
	}

	// Printing the results
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text()) // Calling checkDomain function for each domain entered
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read from input: %v\n", err)
	}
}
