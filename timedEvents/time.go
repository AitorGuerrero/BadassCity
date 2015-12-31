package timedEvents

import "time"

type daysAmount time.Duration

const (
	dayDuration = time.Duration(time.Second * 2)
	localPaymentPeriod = daysAmount(7)
	LocalPaymentDuration = daysAmount(time.Duration(localPaymentPeriod) * dayDuration)
)

func InitTicker (days daysAmount, callback func()) {
	ticker := time.NewTicker(time.Duration(days))
	go func() {
		for range ticker.C {
			callback()
		}
	}()
}