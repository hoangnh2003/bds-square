package service

import (
	"bds-square-backend/internal/repo"
	"bds-square-backend/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
	GetUsers() []string
}

type userService struct {
	userRepo repo.IUserRepository
	//...
}

// GetUsers implements IUserService.
func (us *userService) GetUsers() []string {
	panic("unimplemented")
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	return response.ErrCodeSuccess
}
