// Copyright 2011 Miek Gieben. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Chaos is a small program that prints the version.bind and hostname.bind
// for each address of the nameserver given as argument.
package main

import (
	"fmt"
	"log"

	"github.com/miekg-zdns/dns"
) 

func main() {
	m := new(dns.Msg)
	m.SetQuestion("miek.nl.", dns.TypeSOA)

	c := new(dns.Client)

	fmt.Println("first query")
	r, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.Fatalf("failed to exchange: %v", err)
	}
	if r == nil {
		log.Fatalf("response is nil")
	}
	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("failed to get an valid answer\n%v", r)
	}
	fmt.Println("successful: %v", r)

	// second DNS query
	fmt.Println("second query")
	r, _, err = c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.Fatalf("failed to exchange: %v", err)
	}
	if r == nil {
		log.Fatalf("response is nil")
	}
	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("failed to get an valid answer\n%v", r)
	}
	fmt.Println("successful: %v", r)
}
