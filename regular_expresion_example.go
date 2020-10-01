package main

import (
	"fmt"
	"regexp"
)

const IPRegex string = (`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
var IPMatcher, err = regexp.Compile(IPRegex)

// containsIP returns true if the string passed into it
// contains an IP address
func containsIP(str string) bool {
	return IPMatcher.MatchString(str)
}

// returns a slice of strings containing matches IPs
// if no matches were found, returns an empty slice
func getIPs(str string) []string{
	return IPMatcher.FindAllString(str, -1)
}

func main() {
	str_to_filter := `
	12.34.56.44
	234 244-1567
	test@gmail.com
	192.156.32.15
	lorem ipsum
	127.0.0.1
	`
	contains := containsIP(str_to_filter)
	fmt.Println("String contains IP:", contains)
	
	matches := getIPs(str_to_filter)
	for _, ip := range matches {
		fmt.Println(ip)
	}	
}
