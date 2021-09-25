package main

import (
	"fmt"
	"regexp"
)

var log string = `2021-09-21 22:10 awsbucket 10.10.100.1 -> 10.10.100.2`

func main() {
	p := `(?P<Day>\d{4}-\d{2}-\d{2}) (?P<Time>\d{2}\:\d{2}) (?P<BucketName>[a-z]+) (?P<SrcIp>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) -> (?P<DstIp>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`
	m, _ := regexp.Compile(p)

	matches := m.FindStringSubmatch(log)
	capture_groups := m.SubexpNames()
	for i, match := range matches {
		if i != 0 {
			fmt.Printf("%s == %s\n", capture_groups[i], match)
		}
	}
	//for i, match := range matches {
	//if i != 0 {
	//fmt.Println(match)
	//}
	//}
}
