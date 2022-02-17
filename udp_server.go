package main

import (
        "fmt"
        "net"
)

func startUDPServer(ip string, port int) {
        addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, port))
        ln, _ := net.ListenUDP("udp", addr)

        defer ln.Close()

        buf := make([]byte, 1024)
        for {
                n, cAddr, _ := ln.ReadFromUDP(buf)

                fmt.Printf("%s says: %s\n", cAddr, string(buf[:n]))
        }
}

func main() {
        ip := "10.191"
        port := 3333

        fmt.Printf("listening on udp://%s:%d ...\n", ip, port)
        startUDPServer(ip, port)
}
