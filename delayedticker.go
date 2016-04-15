package delayedticker

import "time"

type DelayedTicker struct {
	C    chan time.Time
	done chan struct{}
}

func NewDelayedTicker(firstDelay, interval time.Duration) *DelayedTicker {
	c := make(chan time.Time, 1)
	ticker := &DelayedTicker{
		C:    c,
		done: make(chan struct{}, 1),
	}
	go ticker.run(firstDelay, interval)
	return ticker
}

func (t *DelayedTicker) Stop() {
	t.done <- struct{}{}
}

func (t *DelayedTicker) run(firstDelay, interval time.Duration) {
	timer := time.NewTimer(firstDelay)
	select {
	case ti := <-timer.C:
		t.sendTime(ti)
	case <-t.done:
		timer.Stop()
		return
	}

	ticker := time.NewTicker(interval)
	for {
		select {
		case ti := <-ticker.C:
			t.sendTime(ti)
		case <-t.done:
			ticker.Stop()
			return
		}
	}
}

func (t *DelayedTicker) sendTime(ti time.Time) {
	select {
	case t.C <- ti:
	default:
	}
}
