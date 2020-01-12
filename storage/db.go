package main

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"gopkg.in/rethinkdb/rethinkdb-go.v6/encoding"
	"log"
	"time"
)

type Document struct {
	Id        string
	Co2       float32
	Tmp       float32
	Timestamp time.Time
}

func EstablishConnection(
	address,
	username,
	password,
	database string) *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address:  address, // endpoint without http
		Username: username,
		Password: password,
		Database: database,
	})
	if err != nil {
		panic(err)
	}
	return session
}

// Query all documents of the requested past and keep following new ones.
func FollowHistory(session *r.Session, sinceHowLong time.Duration, handleDoc func(doc Document) bool) {
	expectedDocuments := sinceHowLong / (5 * time.Second) // we expect a data point every ~5s
	cursor, err := r.Table("airquality").
		OrderBy(r.OrderByOpts{Index: r.Desc("timestamp")}).
		Limit(expectedDocuments * 2). // we fetch twice as many documents as expected to be sure
		Filter(func(row r.Term) r.Term { return row.Field("timestamp").Gt(time.Now().Add(-sinceHowLong)) }).
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)
	if err != nil {
		panic(err)
	}
	var rawDocument map[string]interface{}
	for cursor.Next(&rawDocument) {
		log.Printf("raw: %v\n", rawDocument)
		doc := Document{}
		err := encoding.Decode(&doc, rawDocument["new_val"]) // since we work on a change feed we need key "nev_val"
		if err != nil {
			panic(err)
		}
		if !handleDoc(doc) {
			break
		}
	}
	if cursor.Err() != nil {
		panic(err)
	}
}
