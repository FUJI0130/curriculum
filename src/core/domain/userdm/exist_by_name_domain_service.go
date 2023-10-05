// userdm/exist_by_name_domain_service.go
//
//go:generate mockgen -destination=../../mock/mockExistByNameDomainService/exist_by_name_domain_service_mock.go -package=mockExistByNameDomainService . ExistByNameDomainService

// package interfaces
package userdm

import (
	"context"
	"errors"

	"github.com/FUJI0130/curriculum/src/core/support/customerrors"
)

type ExistByNameDomainService interface {
	Exec(ctx context.Context, name string) (bool, error)
}
type existByNameDomainService struct {
	// userRepo userdm.UserRepository
	userRepo UserRepository
}

// func NewExistByNameDomainService(userRepo userdm.UserRepository) *existByNameDomainService {
func NewExistByNameDomainService(userRepo UserRepository) *existByNameDomainService {
	return &existByNameDomainService{userRepo: userRepo}
}

func (ds *existByNameDomainService) Exec(ctx context.Context, name string) (bool, error) {
	existingUser, err := ds.userRepo.FindByName(ctx, name)

	if err != nil {

		if errors.Is(err, customerrors.WrapNotFoundError(err, "existByNameDomainService.Exec")) {
			return false, nil
		}
		return false, err
	}
	return existingUser != nil, nil
}
