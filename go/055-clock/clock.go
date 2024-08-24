package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	totalMinutes := (h*60 + m) % (24 * 60)

	for totalMinutes < 0 {
		totalMinutes += 24 * 60
	}

	return Clock{
		hour:   totalMinutes / 60,
		minute: totalMinutes % 60,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
