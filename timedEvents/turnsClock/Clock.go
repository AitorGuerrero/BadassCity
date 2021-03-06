package turnsClock

import (
	"github.com/AitorGuerrero/BadassCity/timedEvents"
	"math"
)

type turn int

const daysPerTurn = 1

type Clock struct {
	tickers []*ticker
	timers []*timer
}

func DaysToTurns(d timedEvents.DaysAmount) turn {
	return turn(math.Ceil(float64(d) / float64(daysPerTurn)))
}

func (c *Clock) Next() {
	c.advanceTickers()
	c.advanceTimers()
}

func (c *Clock) Advance(a timedEvents.DaysAmount) {
	for i := 0; i < int(a); i++ {
		c.Next()
	}
}

func (c *Clock) AddTicker(days timedEvents.DaysAmount, callback func()) {
	c.tickers = append(c.tickers, &ticker{
		turnsForCallback: DaysToTurns(days),
		remainingTurns: DaysToTurns(days),
		callback: callback,
	})
}

func (c *Clock) AddTTimer(days timedEvents.DaysAmount, callback func()) {
	c.timers = append(c.timers, &timer{
		remainingTurns: DaysToTurns(days),
		callback: callback,
	})
}

func (c *Clock) advanceTickers() {
	for _, ticker := range c.tickers {
		ticker.next()
	}
}

func (c *Clock) advanceTimers() {
	for index, timer := range c.timers {
		expired := timer.next()
		if expired {
			c.timers = append(c.timers[:index], c.timers[index+1:]...)
		}
	}
}