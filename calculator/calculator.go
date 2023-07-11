package calculator

import (
	"github.com/connorholt/explain_interfaces/badges"
	"github.com/connorholt/explain_interfaces/users"
)

// такой же, как в пакете badges
type handler interface {
	IsEnabled() bool
}

type Calculator struct {
	handlers map[users.BadgeType]handler
}

func (c *Calculator) HasBadge(badge users.BadgeType) bool {
	if handler, ok := c.handlers[badge]; ok {
		return handler.IsEnabled()
	}

	return false // or error
}

func GetCalculator(user interface{}) *Calculator {
	if user, ok := user.(users.Developer); ok {
		return &Calculator{handlers: developerBadges(user)}
	}

	if user, ok := user.(users.Support); ok {
		return &Calculator{handlers: supportBadges(user)}
	}

	return nil
}

func developerBadges(user users.Developer) map[users.BadgeType]handler {
	badgeReview := badges.NewReview(user.CountReview)
	badgeGitStar := badges.NewGitStar(user.CommitCount, badgeReview)
	badgeNoVocation := badges.NewNoVocation(user.LastVocationAt)
	badgeProbation := badges.NewProbation(user.StartWorkAt)
	badgeStoryPoint := badges.NewStoryPoint(user.TaskInDone)

	return map[users.BadgeType]handler{
		users.BadgeReviewer:   badgeReview,
		users.BadgeGitStar:    badgeGitStar,
		users.BadgeNoVacation: badgeNoVocation,
		users.BadgeProbation:  badgeProbation,
		users.BadgeStoryPoint: badgeStoryPoint,
	}
}

func supportBadges(user users.Support) map[users.BadgeType]handler {
	badgeTopSpeed := badges.NewTopSpeed(user.MaxSpeed)
	badgeTopAnswer := badges.NewTopAnswerer(user.CountAnswering)
	badgeNoVocation := badges.NewNoVocation(user.LastVocationAt)
	badgeProbation := badges.NewProbation(user.StartWorkAt)

	return map[users.BadgeType]handler{
		users.BadgeTopSpeed:    badgeTopSpeed,
		users.BadgeTopAnswerer: badgeTopAnswer,
		users.BadgeNoVacation:  badgeNoVocation,
		users.BadgeProbation:   badgeProbation,
	}
}

func marketingBadges(user users.Marketing) map[users.BadgeType]handler {
	return map[users.BadgeType]handler{}
}
