package timer

import "time"

type (
	timer struct {
		start, finish time.Time
	}

	Timer interface {
		Start()
		Stop()
		Timeit(func()) time.Duration
		Delta() time.Duration
	}
)

func (t *timer) Start() {
	t.start = time.Now()
}

func (t *timer) Stop() {
	t.finish = time.Now()
}

func (t *timer) Delta() time.Duration {
	return t.finish.Sub(t.start)
}

func (t *timer) Timeit(f func()) time.Duration {
	t.Start()
	f()
	t.Stop()
	return t.Delta()
}

func New() Timer {
	return &timer{}
}
