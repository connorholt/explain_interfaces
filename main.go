package main

import (
	"fmt"
	"time"
)

// У нас есть сотрудники, и у них есть так называемые ачивки
// Каждый месяц человеку может выдаться разные ачивки:
//
// Больше всего комитов в гите
// Сделал самое большое количество сторипойнтов
// Закончил испытательный срок
// Без отпуска третий год
// Лучший ревьюир
//

type BadgeType string

const (
	TypeDeveloper = "developer"
	TypeSupport   = "support"

	BadgeGitStar    BadgeType = "git_star"
	BadgeStoryPoint BadgeType = "story_point"
	BadgeProbation  BadgeType = "probation"
	BadgeNoVacation BadgeType = "no_vacation"
	BadgeReviewer   BadgeType = "reviewer"
	// новые бейджи
	BadgeTopSpeed    BadgeType = "top_speed"
	BadgeTopAnswerer BadgeType = "top_answerer"
)

type User struct {
	TaskInDone     int
	CommitCount    int
	CountReview    int
	StartWorkAt    time.Time
	LastVocationAt time.Time

	MaxSpeed       int
	CountAnswering int

	Type string
}

func (u *User) HasBadge(badge BadgeType) bool {
	switch badge {
	case BadgeGitStar:
		return u.HasBadgeGitStar()
	case BadgeStoryPoint:
		return u.HasBadgeStoryPoint()
	case BadgeProbation:
		return u.HasBadgeProbation()
	case BadgeNoVacation:
		return u.HasBadgeNoVacation()
	case BadgeReviewer:
		return u.HasBadgeReviewer()
	case BadgeTopSpeed:
		return u.HasBadgeTopSpeed()
	case BadgeTopAnswerer:
		return u.HasBadgeTopAnswerer()
	default:
		return false
	}
}

func (u *User) HasBadgeGitStar() bool {
	if u.Type != TypeDeveloper {
		return false
	}

	if u.CommitCount > 100 && u.HasBadgeReviewer() {
		return true
	}

	return false
}

func (u *User) HasBadgeStoryPoint() bool {
	if u.Type != TypeDeveloper {
		return false
	}

	if u.TaskInDone > 10 {
		return true
	}

	return false
}

func (u *User) HasBadgeProbation() bool {
	if u.StartWorkAt.Unix() <= time.Now().Add(-1*time.Hour*24*3).Unix() {
		return true
	}

	return false
}

func (u *User) HasBadgeNoVacation() bool {
	if u.LastVocationAt.Unix() <= time.Now().Add(-1*time.Hour*24*10).Unix() {
		return true
	}

	return false
}

func (u *User) HasBadgeTopSpeed() bool {
	if u.Type != TypeSupport {
		return false
	}

	if u.MaxSpeed > 100 {
		return true
	}

	return false
}

func (u *User) HasBadgeTopAnswerer() bool {
	if u.Type != TypeSupport {
		return false
	}

	if u.CountAnswering > 300 {
		return true
	}

	return false
}

func (u *User) HasBadgeReviewer() bool {
	if u.Type != TypeDeveloper {
		return false
	}

	if u.CountReview > 20 {
		return true
	}

	return false
}

func main() {
	// с базы данных
	user := &User{
		TaskInDone:     10,
		CommitCount:    100,
		CountReview:    24,
		StartWorkAt:    time.Now(),
		LastVocationAt: time.Now(),
		Type:           TypeDeveloper,
	}

	badges := []BadgeType{
		BadgeGitStar,
		BadgeStoryPoint,
		BadgeProbation,
		BadgeNoVacation,
		BadgeReviewer,
	}

	for _, badge := range badges {
		fmt.Println(user.HasBadge(badge))
	}

	support := &User{
		MaxSpeed:       1,
		CountAnswering: 1000,
		StartWorkAt:    time.Now(),
		LastVocationAt: time.Now(),
		Type:           TypeSupport,
	}

	for _, badge := range badges {
		fmt.Println(support.HasBadge(badge))
	}
}
