package badges

type GitStar struct {
	commitCount   int
	reviewHandler handler
}

func NewGitStar(commitCount int, reviewHandler handler) *GitStar {
	return &GitStar{
		commitCount:   commitCount,
		reviewHandler: reviewHandler,
	}
}

func (b *GitStar) IsEnabled() bool {
	if b.commitCount > 100 && b.reviewHandler.IsEnabled() {
		return true
	}

	return false
}
