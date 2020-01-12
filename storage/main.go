package main

import (
	"log"
	"time"
)

func main() {
	conn := EstablishConnection(
		"rethinkdb.isotronic.de",
		"sensor",
		"xxx",
		"homeautomation")
	defer conn.Close()

	FollowHistory(conn,
		time.Hour,
		func(doc Document) bool {
			log.Printf("%v: %v\n", doc.Timestamp.In(time.Local), doc.Co2)
			return true // continue processing
		})
}
