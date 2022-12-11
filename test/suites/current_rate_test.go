package suites

import (
	"context"
	"flag"
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protos "github.com/mannanmcc/proto/rates/rate"
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
	serverAddr = flag.String("addr", "localhost:9001", "The server address in the format of host:port")
)

var _ = BeforeSuite(func() {
	cnn, _ := newBasicClient()
	client = protos.NewRateServiceClient(cnn)
})

func TestCurrency(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {
	When("A valid rpc request received", func() {
		It("should return rate", func() {
			expoectedRate := "1.1"
			rate, err := client.GetRate(context.Background(), &protos.RateRequest{BaseCurrency: "GBP", TargetCurrency: "GBP"})
			Expect(err).To(BeNil())
			Expect(rate.GetRate()).To(Equal(expoectedRate))
		})
	})
})
