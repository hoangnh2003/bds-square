package repo

type IUserRepository interface {
	GetUserByEmail(email string) bool
	// GetUsers() ([]*model.GoDbUser, error)
}

type userRepository struct {
	// q *query.Query
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		// q: query.Use(db),
	}
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}

