package services

import (
	"github.com/triostones/gin-mvc-demo/dao"
	"github.com/triostones/gin-mvc-demo/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) Login(email string, password string) (*models.UserModel, error) {
	user, err := (&dao.UserDao{}).GetByEmail(email)
	if err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return user, nil
	}
	return user, nil
}
