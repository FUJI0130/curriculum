package mentorrecruitmentdm

import "errors"

type Budget struct {
	from int
	to   int
}

const budgetMin = 1000

func NewBudget(from int, to int) (*Budget, error) {
	if from < budgetMin {
		return nil, errors.New("invalid budget: from cannot be less than budgetMin")
	}

	if to < from {
		return nil, errors.New("invalid budget: to cannot be less than from")
	}

	return &Budget{from: from, to: to}, nil
}

func (b *Budget) From() int {
	return b.from
}

func (b *Budget) To() int {
	return b.to
}
