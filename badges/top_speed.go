package badges

type TopSpeed struct {
	maxSpeed int
}

func NewTopSpeed(maxSpeed int) *TopSpeed {
	return &TopSpeed{
		maxSpeed: maxSpeed,
	}
}

func (b *TopSpeed) IsEnabled() bool {
	if b.maxSpeed > 100 {
		return true
	}

	return false
}
