// userdm/exist_by_name_domain_service.go
//
//go:generate mockgen -destination=../../mock/mockExistByNameDomainService/exist_by_name_domain_service_mock.go -package=mockExistByNameDomainService . ExistByNameDomainService

// package interfaces
package userdm

import (
	"context"
	"errors"
	"log"

	"github.com/FUJI0130/curriculum/src/core/domain/customerrors"
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
	log.Printf("Actual type of the error: %T", err)

	log.Printf("after FindByName  exist_by_name_domain_service  ErrUserNotFound  err is : %s ", err)
	// if _, ok := err.(*customerrors.UserNotFoundError); ok {
	// 	log.Printf("exist_by_name_domain_service  ErrUserNotFound  err is : %s ", err)
	// 	return false, nil
	// }

	if err != nil {
		log.Printf("Expected UserNotFound error: %T - %v", customerrors.ErrUserNotFound(), customerrors.ErrUserNotFound())
		if err == customerrors.ErrUserNotFound() {
			log.Printf("Direct comparison succeeded!")
		} else {
			log.Printf("Direct comparison failed!")
		}

		if errors.Is(err, customerrors.ErrUserNotFound()) {
			// if customerrors.Is(err, customerrors.ErrUserNotFound()) {
			log.Printf("exist_by_name_domain_service  ErrUserNotFound  err is : %s ", err)
			return false, nil
		}
		return false, err
	}
	return existingUser != nil, nil
}
