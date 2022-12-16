package suites

import (
	"flag"
	"log"

	protos "github.com/mannanmcc/proto/rates/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	. "github.com/onsi/ginkgo"
)

func newBasicClient() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return nil, err
	}
	//defer conn.Close()
	return conn, nil
}

var (
	client     protos.RateServiceClient
	serverAddr = flag.String("addr", "api:50051", "The server address in the format of host:port")
)

var _ = BeforeSuite(func() {
	cnn, _ := newBasicClient()
	client = protos.NewRateServiceClient(cnn)
})
