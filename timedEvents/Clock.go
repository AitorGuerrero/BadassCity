package timedEvents

type DaysAmount int

const (
	Day = DaysAmount(1)
	Week = Day * 7
)

type Clock interface {
	AddTicker(days DaysAmount, callback func())
	AddTTimer(days DaysAmount, callback func())
}