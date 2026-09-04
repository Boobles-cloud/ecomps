package services

import (
	"context"
	"errors"

	userstructs "ecomps.boobles.cloud/backend/internal/user/user_structs"
)

// Gets a user by the given id
func (s *UserService) GetUserById(ctx context.Context, userId uint) (userstructs.UserStruct, error) {

	if userId == 0 {
		return userstructs.UserStruct{}, errors.New("User Id cant be 0")
	}

	return s.userRepo.GetById(ctx, userId)
}

func (s *UserService) GetAllByTenantId(ctx context.Context, tenantId uint) ([]userstructs.UserStruct, error) {

	users := make([]userstructs.UserStruct, 0, 100)

	if tenantId == 0 {
		return users, errors.New("Tenant Id cant be 0")
	}

	return s.userRepo.GetAllByTenantId(ctx, tenantId)
}

func (s *UserService) CreateUser(ctx context.Context, user userstructs.UserStruct) (uint, error) {

	if user.UserName == "" {
		return 0, errors.New("Username is empty")
	}

	if user.UserMail == "" {
		return 0, errors.New("UserEmail is empty")
	}

	if user.UserPW == "" {
		return 0, errors.New("User pw is empty")
	}

	tmp, err := s.userRepo.GetUserByEmail(ctx, user.UserMail)

	if err == nil && tmp.UserMail == user.UserMail {
		return 0, errors.New("User is already registert")
	}

	return s.userRepo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user userstructs.UserStruct, filterName string) error {
	// TODO
	return errors.ErrUnsupported
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	// TODO
	return errors.ErrUnsupported
}
