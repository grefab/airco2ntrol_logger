package lib

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "storage/pb"
)

type server struct {
}

func RunServer(ip net.IP, port int) {
	lis, err := net.ListenTCP("tcp", &net.TCPAddr{IP: ip, Port: port})
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (s *server) GetSince(oldest *timestamp.Timestamp, stream pb.Storage_GetSinceServer) error {
	conn := EstablishConnection(
		"rethinkdb.isotronic.de",
		"sensor",
		"xxx",
		"homeautomation")
	defer conn.Close()

	ts, err := ptypes.Timestamp(oldest)
	if err != nil {
		panic(err)
	}
	FollowHistory(conn,
		ts,
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
