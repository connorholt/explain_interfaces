package badges

import (
	"time"
)

type Probation struct {
	startWorkAt time.Time
}

func NewProbation(startWorkAt time.Time) *Probation {
	return &Probation{
		startWorkAt: startWorkAt,
	}
}

func (b *Probation) IsEnabled() bool {
	if b.startWorkAt.Unix() <= time.Now().Add(-1*time.Hour*24*3).Unix() {
		return true
	}

	return false
}
