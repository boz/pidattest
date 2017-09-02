package worker

import (
	"log"
	"net"
	"time"

	context "golang.org/x/net/context"

	"github.com/boz/circumspect/proto"

	grpc "google.golang.org/grpc"
)

func Run(log *log.Logger, path string) error {

	log.Printf("connecting to %v...", path)

	conn, err := grpc.Dial(path, grpc.WithInsecure(), grpc.WithDialer(dialer))
	if err != nil {
		log.Printf("error connecting to %v: %v", path, err)
		return err
	}
	defer conn.Close()

	client := proto.NewWorkloadClient(conn)

	_, err = client.Register(context.Background(), &proto.Request{})
	if err != nil {
		log.Printf("error registering: %v", err)
		return err
	}

	log.Print("ok")

	return nil
}

func dialer(addr string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", addr, timeout)
}
