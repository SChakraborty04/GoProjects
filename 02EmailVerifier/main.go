package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the domain name you want to check: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err:=scanner.Err();err!=nil{
		panic(err)
	}
}

func checkDomain(domain string) {
	var hasMx,hasSPF,hasDMARC bool
	var spfRecord,dmarkRecord string

	mxRecord,err:=net.LookupMX(domain)
	if err!=nil{
		log.Printf("No MX record found for %s\n",domain)
	}
	if len(mxRecord)>0{
		hasMx=true
	}
	txtRecords,err:=net.LookupTXT(domain)
	if err!=nil{
		log.Printf("No TXT record found for %s\n",domain)
	}
	for _,record:=range txtRecords{
		if strings.HasPrefix(record,"v=spf1"){
			hasSPF=true
			spfRecord=record
			break
		}
	}
	dmarcRecords,err:=net.LookupTXT("_dmarc."+domain)
	if err!=nil{
		log.Printf("No DMARC record found for %s\n",domain)
	}
	for _,record:=range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC=true
			dmarkRecord=record
			break
		}
	}
	fmt.Printf("Input Domain: %s\nHas MX record: %v\nHas SPF Record: %v\nSPF record: %v\nHas DMARC Record: %v\nDMARC Record: %v\n",domain,hasMx,hasSPF,spfRecord,hasDMARC,dmarkRecord)
}