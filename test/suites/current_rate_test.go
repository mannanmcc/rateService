package suites

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	protos "github.com/mannanmcc/proto/rates/rate"
)

func TestCurrency(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

const (
	currencyGBP  = "GBP"
	currencyEURO = "EUR"
	currencyBDT  = "BDT"
)

var _ = Describe("Shopping cart", func() {

	When("A valid rpc request received", func() {
		It("should return rate", func() {
			expoectedRate := "1.10"
			rate, err := client.GetRate(context.Background(), &protos.RateRequest{BaseCurrency: currencyGBP, TargetCurrency: currencyEURO})
			Expect(err).To(BeNil())
			Expect(rate.GetRate()).To(Equal(expoectedRate))
		})
	})
})
