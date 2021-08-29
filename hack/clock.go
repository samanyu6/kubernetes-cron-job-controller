package hack

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (c realClock) Now() time.Time {
	return time.Now()
}
