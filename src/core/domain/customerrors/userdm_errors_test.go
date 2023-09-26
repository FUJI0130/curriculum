package customerrors_test

import (
	"errors"
	"testing"

	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
)

func TestErrorIs(t *testing.T) {
	err := customerrors.ErrUserNotFound()
	if !errors.Is(err, customerrors.ErrUserNotFound()) {
		t.Fatal("errors.Is did not recognize ErrUserNotFound")
	}
}
