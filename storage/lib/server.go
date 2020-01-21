package lib

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "storage/pb"
	"time"
)

type server struct {
}

func RunServer(ip net.IP, port int) {
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})

	lis, err := net.ListenTCP("tcp", &net.TCPAddr{IP: ip, Port: port})
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func protoTimestampToTime(protoTimestamp *timestamp.Timestamp) time.Time {
	ts, err := ptypes.Timestamp(protoTimestamp)
	if err != nil {
		panic(err)
	}
	return ts
}

func connDetails() (address,
	username,
	password,
	database string) {
	return "rethinkdb.isotronic.de",
		"sensor",
		"xxx",
		"homeautomation"
}

func (s *server) GetBatch(ctx context.Context, timeFrame *pb.TimeFrame) (*pb.Batch, error) {
	log.Printf("fetching batch between [%v, %v]", protoTimestampToTime(timeFrame.From), protoTimestampToTime(timeFrame.To))
	conn := DbConnection(connDetails())
	defer conn.Close()

	from := protoTimestampToTime(timeFrame.From)
	to := protoTimestampToTime(timeFrame.To)
	batch, err := FetchBatch(conn, from, to)
	if err != nil {
		log.Printf("GetBatch(): %v", err)
		return nil, err
	}

	return batch, nil
}

func (s *server) GetSince(oldest *timestamp.Timestamp, stream pb.Storage_GetSinceServer) error {
	log.Printf("GetSince(%v)", oldest)
	conn := DbConnection(connDetails())
	defer conn.Close()

	ts, err := ptypes.Timestamp(oldest)
	if err != nil {
		return err
	}

	err = FollowHistory(conn,
		ts,
		func(doc pb.AirQuality) (bool, error) {
			err := stream.Send(&doc)
			if err != nil {
				return false, err
			}
			return true, err
		})
	if err != nil {
		log.Printf("GetSince(): %v", err)
		return err
	}

	return nil
}
