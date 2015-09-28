package paymentscenarios

import (
	"os"
	"testing"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

func TestCreate(t *testing.T) {
	var err error

	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")
	c := ticketmatic.NewClient(accountcode, accesskey, secretkey)

	paymentscenario, err := Create(c, &ticketmatic.PaymentScenario{
		Availability: &ticketmatic.PaymentscenarioAvailability{
			Saleschannels: []int64{
				1,
			},
			Usescript: false,
		},
		Expiryparameters: &ticketmatic.PaymentscenarioExpiryParameters{
			Daysafterordercreation: 24,
			Daysbeforeevent:        2,
		},
		Internalremark: "Testing",
		Name:           "Payment scenario test",
		Overdueparameters: &ticketmatic.PaymentscenarioOverdueParameters{
			Daysafterordercreation: 12,
			Daysbeforeevent:        5,
		},
		Paymentmethods: []int64{
			1,
			2,
		},
		Shortdescription: "Short test",
		Typeid:           2702,
	})
	if err != nil {
		t.Fatal(err)
	}

	if paymentscenario.Id == 0 {
		t.Errorf("Unexpected paymentscenario.Id, got %#v, expected different value", paymentscenario.Id)
	}

	if paymentscenario.Typeid != 2702 {
		t.Errorf("Unexpected paymentscenario.Typeid, got %#v, expected %#v", paymentscenario.Typeid, 2702)
	}

	if paymentscenario.Name != "Payment scenario test" {
		t.Errorf("Unexpected paymentscenario.Name, got %#v, expected %#v", paymentscenario.Name, "Payment scenario test")
	}

}