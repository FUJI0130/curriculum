// userdm/exist_by_name_domain_service.go
//
//go:generate mockgen -destination=../../mock/mockExistByNameDomainService/exist_by_name_domain_service_mock.go -package=mockExistByNameDomainService . ExistByNameDomainService

// package interfaces
package userdm

import (
	"context"
	"errors"
)

type ExistByNameDomainService struct {
	userRepo UserRepository
}

// func NewExistByNameDomainService(userRepo userdm.UserRepository) *existByNameDomainService {
func NewExistByNameDomainService(userRepo UserRepository) *ExistByNameDomainService {
	return &ExistByNameDomainService{userRepo: userRepo}
}

func (ds *ExistByNameDomainService) Exec(ctx context.Context, name string) (bool, error) {
	existingUser, err := ds.userRepo.FindByName(ctx, name)
	if err != nil {
		// if errors.Is(err, userdm.ErrUserNotFound) {
		if errors.Is(err, ErrUserNotFound) {
			return false, nil
		}
		return false, err
	}
	return existingUser != nil, nil
}
