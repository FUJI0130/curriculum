// userdm/exist_by_name_domain_service.go

package userdm

import (
	"context"
)

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
