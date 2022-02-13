package server

import (
	"io"
	"log"
	"math/rand"

	pb "github.com/jzhang046/grpc-blackhole-go/blackhole"
)

func New() pb.BlackHoleServer {
	return &blackHoleServer{}
}

type blackHoleServer struct {
	pb.UnimplementedBlackHoleServer
}

func (s *blackHoleServer) ConsumeAll(stream pb.BlackHole_ConsumeAllServer) error {
	var count uint64 = 0
	for {
		bytes, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.ByteCount{Count: count})
		} else if err != nil {
			return err
		}
		count += uint64(len(bytes.Payload))
	}
}

func (s *blackHoleServer) EmitZeros(streamSize *pb.StreamSize, stream pb.BlackHole_EmitZerosServer) error {
	log.Printf("Received EmitZeros request [%v] \n", streamSize)
	data := make([]byte, streamSize.Length)
	for i := streamSize.Count; i > 0; i-- {
		if err := stream.Send(&pb.Bytes{Payload: data}); err != nil {
			return err
		}
	}
	return nil
}

func (s *blackHoleServer) EmitRandom(streamSize *pb.StreamSize, stream pb.BlackHole_EmitRandomServer) error {
	log.Printf("Received EmitRandom request [%v] \n", streamSize)
	for i := streamSize.Count; i > 0; i-- {
		data := make([]byte, streamSize.Length)
		rand.Read(data)
		if err := stream.Send(&pb.Bytes{Payload: data}); err != nil {
			return err
		}
	}
	return nil
}
