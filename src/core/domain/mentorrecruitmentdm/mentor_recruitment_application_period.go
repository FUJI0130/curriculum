package mentorrecruitmentdm

import (
	"time"

	"github.com/cockroachdb/errors"
)

type ApplicationPeriod struct {
	from time.Time
	to   time.Time
}

func NewApplicationPeriod(from time.Time, to time.Time) (*ApplicationPeriod, error) {
	if from.After(to) {
		return nil, errors.New("application period is invalid: from should not be after to")
	}

	return &ApplicationPeriod{from: from, to: to}, nil
}

func (ap *ApplicationPeriod) From() time.Time {
	return ap.from
}

func (ap *ApplicationPeriod) To() time.Time {
	return ap.to
}
