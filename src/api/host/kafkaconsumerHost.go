package host

import (
	"github.com/radyatamaa/loyalti-go-echo/src/api/host/consumer"
)
func StartKafka() {
	//consumer.NewReceiver()
	go consumer.NewEmployeeConsumer()
	go consumer.NewMerchantConsumer()
	go consumer.NewOutletConsumer()
	go consumer.NewProgramConsumer()
	go consumer.NewCardConsumer()
	go consumer.NewSpecialConsumer()
}
