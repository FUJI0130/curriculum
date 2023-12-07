package mentorrecruitmentdm

import "github.com/cockroachdb/errors"

type Budget struct {
	from uint32
	to   uint32
}

const (
	budgetMin = 1000
	budgetMax = 100000
)

func NewBudget(from uint32, to uint32) (*Budget, error) {
	if from < budgetMin {
		return nil, errors.New("invalid budget: from cannot be less than budgetMin")
	}

	if from > budgetMax {
		return nil, errors.New("invalid budget: from cannot exceed budgetMax")
	}

	if to < from {
		return nil, errors.New("invalid budget: to cannot be less than from")
	}

	if to > budgetMax {
		return nil, errors.New("invalid budget: to cannot exceed budgetMax")
	}

	return &Budget{from: from, to: to}, nil
}

func (b *Budget) From() uint32 {
	return b.from
}

func (b *Budget) To() uint32 {
	return b.to
}
