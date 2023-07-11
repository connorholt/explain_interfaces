package badges

type StoryPoint struct {
	taskInDone int
}

func NewStoryPoint(taskInDone int) *StoryPoint {
	return &StoryPoint{
		taskInDone: taskInDone,
	}
}

func (b *StoryPoint) IsEnabled() bool {
	if b.taskInDone > 10 {
		return true
	}

	return false
}
