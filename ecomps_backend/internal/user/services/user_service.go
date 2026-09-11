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

// Updates a user
func (s *UserService) UpdateUser(ctx context.Context, user userstructs.UserStruct, filterName string) error {
	return s.userRepo.Update(ctx, user, "UserId")
}

// Deletes a user
func (s *UserService) DeleteUser(ctx context.Context, userId, tenantId uint) error {

	if userId == 0 {
		return errors.New("UserId cannot be 0")
	}

	tenant, err := s.tenantService.GetTenantById(ctx, tenantId)

	if err != nil {
		return err
	}

	// TODO: insert Tenant service
	// isThere := s.tenantService.CheckTenantDeletion()

	// If the user isnt a admin user we can just delete him
	// If he is a admin user and the tenant is in deletion we set the date of deletion to 2 Months
	if tenant.IsUserAdmin(userId) {
		return s.userRepo.CreateUserDeletionDate(ctx, userId)
	} else if !tenant.IsUserAdmin(userId) {
		s.DeleteUser(ctx, userId, tenantId)
	} else {
		return errors.New("User is still admin!")
	}

	return nil
}
