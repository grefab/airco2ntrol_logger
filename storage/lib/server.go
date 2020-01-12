package lib

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "storage/pb"
	"time"
)

type server struct {
}

func RunServer(ip net.IP, port int, async bool) {
	lis, err := net.ListenTCP("tcp", &net.TCPAddr{IP: ip, Port: port})
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})

	run := func() {
		err := s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}

	if async {
		go run()
	} else {
		run()
	}
}

func (s *server) GetSince(oldest *timestamp.Timestamp, stream pb.Storage_GetSinceServer) error {
	conn := EstablishConnection(
		"rethinkdb.isotronic.de",
		"sensor",
		"S3nsor#D4ta",
		"homeautomation")
	defer conn.Close()

	FollowHistory(conn,
		time.Hour,
		func(doc pb.AirQuality) bool {
			log.Printf("%v\n", proto.MarshalTextString(&doc))
			err := stream.Send(&doc)
			if err != nil {
				panic(err)
			}
			return true
		})

	return nil
}
