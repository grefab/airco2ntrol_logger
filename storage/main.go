package main

import (
	flag "github.com/spf13/pflag"
	"log"
	"net"
	"regexp"
	"storage/lib"
	"strconv"
)

func main() {
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds)
	endpoint := flag.String("endpoint", "0.0.0.0:15800", "endpoint to listen to in server mode or to connect to in client mode")

	ip, port := parseEndpoint(endpoint)

	log.Printf("running server on %v:%v", ip, port)
	lib.RunServer(ip, port)
}

func parseEndpoint(endpoint *string) (net.IP, int) {
	parse := regexp.MustCompile(`^(.*):(.*)$`)
	matches := parse.FindAllStringSubmatch(*endpoint, -1)

	ip := net.ParseIP(matches[0][1])
	if ip == nil {
		log.Fatalf("not a valid ip: %v", matches[0][1])
	}

	port, err := strconv.Atoi(matches[0][2])
	if err != nil {
		log.Fatalf("not a valid port: %v: %v", matches[0][2], err)
	}
	return ip, port
}
