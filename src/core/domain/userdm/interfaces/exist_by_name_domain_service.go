// userdm/exist_by_name_domain_service.go
//
//go:generate mockgen -destination=../../mock/mockExistByNameDomainService/exist_by_name_domain_service_mock.go -package=mockExistByNameDomainService . ExistByNameDomainService

package interfaces

import (
	"context"
	"errors"

	"github.com/FUJI0130/curriculum/src/core/domain/userdm"
)

type ExistByNameDomainService interface {
	IsExist(ctx context.Context, name string) (bool, error)
}
type existByNameDomainService struct {
	userRepo userdm.UserRepository
}

func NewExistByNameDomainService(userRepo userdm.UserRepository) *existByNameDomainService {
	return &existByNameDomainService{userRepo: userRepo}
}

func (ds *existByNameDomainService) IsExist(ctx context.Context, name string) (bool, error) {
	existingUser, err := ds.userRepo.FindByName(ctx, name)
	if err != nil {
		if errors.Is(err, userdm.ErrUserNotFound) {
			return false, nil
		}
		return false, err
	}
	return existingUser != nil, nil
}
