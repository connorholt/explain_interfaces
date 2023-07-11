package main

import (
	"github.com/connorholt/explain_interfaces/calculator"
	"github.com/connorholt/explain_interfaces/users"
)

func main() {
	developer := &users.Developer{
		TaskInDone:  100,
		CommitCount: 10,
	}

	calc := calculator.GetCalculator(developer)
	calc.HasBadge(users.BadgeNoVacation)

	// -------
	support := &users.Support{
		CountAnswering: 1000,
	}

	calc = calculator.GetCalculator(support)
	calc.HasBadge(users.BadgeNoVacation)
}
