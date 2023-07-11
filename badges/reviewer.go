package badges

type Review struct {
	countReview int
}

func NewReview(countReview int) *Review {
	return &Review{
		countReview: countReview,
	}
}

func (b *Review) IsEnabled() bool {
	if b.countReview > 20 {
		return true
	}

	return false
}
