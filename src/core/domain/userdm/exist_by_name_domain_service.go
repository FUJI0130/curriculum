// userdm/exist_by_name_domain_service.go
//
//go:generate mockgen -destination=../../mock/mockUser/exist_by_name_domain_service_mock.go -package=mockUser . ExistByNameDomainService
package userdm

import (
	"context"
)

type ExistByNameDomainService interface {
	IsExist(ctx context.Context, name string) (bool, error)
}
type existByNameDomainService struct {
	userRepo UserRepository
}

func NewExistByNameDomainService(userRepo UserRepository) *existByNameDomainService {
	return &existByNameDomainService{userRepo: userRepo}
}

func (ds *existByNameDomainService) IsExist(ctx context.Context, name string) (bool, error) {
	existingUser, err := ds.userRepo.FindByName(ctx, name)
	if err != nil {
		return false, err
	}
	return existingUser != nil, nil
}
