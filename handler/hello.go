package handler

import (
	"context"
	"io"
	"time"

	"go-micro.dev/v4/metadata"

	log "go-micro.dev/v4/logger"

	pb "go-micro-study/proto"
)

type Hello struct{}

func (e *Hello) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	meta, _ := metadata.FromContext(ctx)
	log.Infof("Received Hello.Call request: %v meta:%v", req, meta["id"])

	rsp.Msg = "Received " + req.Name
	return nil
}

func (e *Hello) ClientStream(ctx context.Context, stream pb.Hello_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Hello) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Hello_ServerStreamStream) error {
	log.Infof("Received Hello.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Hello) BidiStream(ctx context.Context, stream pb.Hello_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
