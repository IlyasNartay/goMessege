package user

import "log"

// User represents a user entity
type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

// UserService provides user management functionality
type UserService interface {
	CreateUser(user *User) error
	GetUserByID(userID string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

// UserRepository provides methods to interact with the database for user management
type UserRepository interface {
	SaveUser(user *User) error
	GetUserByID(userID string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}

// UserServiceImpl implements UserService interface
type UserServiceImpl struct {
	repo   UserRepository
	logger *log.Logger
}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService(repo UserRepository, logger *log.Logger) *UserServiceImpl {
	return &UserServiceImpl{repo: repo, logger: logger}
}

// CreateUser creates a new user
func (s *UserServiceImpl) CreateUser(user *User) error {
	// Save user to the repository
	err := s.repo.SaveUser(user)
	if err != nil {
		s.logger.Printf("Failed to save user: %v", err)
		return err
	}
	return nil
}

// GetUserByID retrieves a user by ID
func (s *UserServiceImpl) GetUserByID(userID string) (*User, error) {
	// Retrieve user from the repository
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		s.logger.Printf("Failed to get user by ID %s: %v", userID, err)
		return nil, err
	}
	return user, nil
}

// GetUserByUsername retrieves a user by username
func (s *UserServiceImpl) GetUserByUsername(username string) (*User, error) {
	// Retrieve user from the repository
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		s.logger.Printf("Failed to get user by username %s: %v", username, err)
		return nil, err
	}
	return user, nil
}
