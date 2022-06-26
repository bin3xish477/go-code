package main

import (
        "log"
        "fmt"

        "github.com/google/gopacket/pcap"
)

func main() {
        devices, err := pcap.FindAllDevs()
        if err != nil {
                log.Printf(err.Error())
        }

        for _, dev := range devices {
                fmt.Printf("%s\n", dev.Name)
                for _, addr := range dev.Addresses {
                        fmt.Printf("\tIP: %s\n", addr.IP)
                        fmt.Printf("\tNetmask: %s\n", addr.Netmask)
                }
        }
}
