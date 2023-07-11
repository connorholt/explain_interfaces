package badges

import (
	"time"
)

type NoVocation struct {
	lastVocationAt time.Time
}

func NewNoVocation(lastVocationAt time.Time) *NoVocation {
	return &NoVocation{
		lastVocationAt: lastVocationAt,
	}
}

func (b *NoVocation) IsEnabled() bool {
	if b.lastVocationAt.Unix() <= time.Now().Add(-1*time.Hour*24*10).Unix() {
		return true
	}

	return false
}
