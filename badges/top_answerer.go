package badges

type TopAnswerer struct {
	countAnswering int
}

func NewTopAnswerer(countAnswering int) *TopAnswerer {
	return &TopAnswerer{
		countAnswering: countAnswering,
	}
}

func (b *TopAnswerer) IsEnabled() bool {
	if b.countAnswering > 300 {
		return true
	}

	return false
}
