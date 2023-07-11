package users

import (
	"time"
)

type BadgeType string

const (
	TypeDeveloper = "developer"
	TypeSupport   = "support"
	TypeMarketing = "marketing"

	BadgeGitStar    BadgeType = "git_star"
	BadgeStoryPoint BadgeType = "story_point"
	BadgeProbation  BadgeType = "probation"
	BadgeNoVacation BadgeType = "no_vacation"
	BadgeReviewer   BadgeType = "reviewer"
	// новые бейджи
	BadgeTopSpeed    BadgeType = "top_speed"
	BadgeTopAnswerer BadgeType = "top_answerer"
)

type Developer struct {
	TaskInDone     int
	CommitCount    int
	CountReview    int
	StartWorkAt    time.Time
	LastVocationAt time.Time
}

type Support struct {
	MaxSpeed       int
	CountAnswering int
	StartWorkAt    time.Time
	LastVocationAt time.Time
}

type Marketing struct {
	MaxSpeed       int
	CountAnswering int
	StartWorkAt    time.Time
	LastVocationAt time.Time
}
