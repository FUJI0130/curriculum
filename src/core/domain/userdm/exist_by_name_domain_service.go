//go:generate mockgen -destination=../../mock/mockExistByNameDomainService/exist_by_name_domain_service_mock.go -package=mockExistByNameDomainService . ExistByNameDomainService

package userdm

import (
	"context"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type ExistByNameDomainService interface {
	Exec(ctx context.Context, name string) (bool, error)
}
type existByNameDomainService struct {
	userRepo UserRepository
}

func NewExistByNameDomainService(userRepo UserRepository) *existByNameDomainService {
	return &existByNameDomainService{userRepo: userRepo}
}

func (ds *existByNameDomainService) Exec(ctx context.Context, name string) (bool, error) {
	existingUser, err := ds.userRepo.FindByName(ctx, name)
	if err != nil {

		if customerrors.IsNotFoundError(err) {
			return false, nil
		}
		return false, customerrors.WrapInternalServerError(err, "An error occurred in Exec when calling userRepo.FindByName")
	}
	return existingUser != nil, nil
}
