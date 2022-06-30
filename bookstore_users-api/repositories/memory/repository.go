package memory

import (
	"fmt"
	"math/rand"

	"github.com/SantiagoBedoya/bookstore_users-api/users"
	"github.com/SantiagoBedoya/bookstore_utils/resterrors"
)

type memoryRepository struct {
	data []users.User
}

func NewMemoryRepository() users.Repository {
	return &memoryRepository{
		data: make([]users.User, 0),
	}
}

func (r *memoryRepository) CreateUser(user *users.User) (*users.User, *resterrors.RestError) {
	user.Id = int64(rand.Intn(100))
	r.data = append(r.data, *user)
	return user, nil
}

func (r *memoryRepository) GetUser(userId string) (*users.User, *resterrors.RestError) {
	for _, user := range r.data {
		if fmt.Sprint(user.Id) == userId {
			return &user, nil
		}
	}
	return nil, resterrors.NewBadRequestError(users.UserNotFoundErr.Error())
}

func (r *memoryRepository) SearchUserByEmail(email string) (*users.User, *resterrors.RestError) {
	for _, user := range r.data {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, resterrors.NewBadRequestError(users.UserNotFoundErr.Error())
}

func (r *memoryRepository) UpdateUser(userId string, user *users.User) (*users.User, *resterrors.RestError) {
	return nil, nil
}
func (r *memoryRepository) DeleteUser(userId string) (*users.User, *resterrors.RestError) {
	return nil, nil
}
func (r *memoryRepository) GetUsersByStatus(status string) ([]users.User, *resterrors.RestError) {
	return nil, nil
}
func (r *memoryRepository) GetUsers() ([]users.User, *resterrors.RestError) {
	return nil, nil
}
