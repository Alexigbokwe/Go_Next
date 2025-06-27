package repository

type UserRepository struct{}

func (ur *UserRepository) FindAll() ([]string, error) {
	// Implementation for getting users from the database
	return []string{"user1", "user2"}, nil
}
