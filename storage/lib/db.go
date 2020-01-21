package lib

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"gopkg.in/rethinkdb/rethinkdb-go.v6/encoding"
	"log"
	airco2ntrol "storage/pb"
	"time"
)

type Document struct {
	Id        string
	Co2       float32
	Tmp       float32
	Timestamp time.Time
}

func DbConnection(
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
func FollowHistory(session *r.Session, sinceWhen time.Time, handleDoc func(doc airco2ntrol.AirQuality) (bool, error)) error {
	expectedDocuments := 10
	cursor, err := r.Table("airquality").
		OrderBy(r.OrderByOpts{Index: r.Desc("timestamp")}).
		Limit(expectedDocuments).
		Filter(func(row r.Term) r.Term { return row.Field("timestamp").Gt(sinceWhen) }).
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
			return err
		}
		proceed, err := handleDoc(makeProto(doc))
		if err != nil {
			return err
		}
		if !proceed {
			break
		}
	}
	if cursor.Err() != nil {
		return err
	}

	return nil
}

func FetchBatch(session *r.Session, from time.Time, to time.Time) (*airco2ntrol.Batch, error) {
	cursor, err := r.Table("airquality").
		Between(from, to, r.BetweenOpts{Index: "timestamp"}).
		OrderBy(r.OrderByOpts{Index: "timestamp"}).
		Run(session)
	if err != nil {
		panic(err)
	}

	var result airco2ntrol.Batch
	var rawDocument map[string]interface{}
	for cursor.Next(&rawDocument) {
		doc := Document{}
		err := encoding.Decode(&doc, rawDocument)
		if err != nil {
			return nil, err
		}

		d := makeProto(doc)
		result.Items = append(result.Items, &d)
	}
	if cursor.Err() != nil {
		return nil, err
	}

	log.Printf("fetched %v items", len(result.Items))
	return &result, nil
}

func makeProto(doc Document) airco2ntrol.AirQuality {
	pb := airco2ntrol.AirQuality{
		Timestamp: func() *timestamp.Timestamp {
			ts, err := ptypes.TimestampProto(doc.Timestamp)
			if err != nil {
				panic(err)
			}
			return ts
		}(),
		Tmp: doc.Tmp,
		Co2: doc.Co2,
	}
	return pb
}
