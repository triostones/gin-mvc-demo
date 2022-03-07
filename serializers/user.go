package serializers

import (
	"github.com/triostones/gin-mvc-demo/models"
)

type User struct {
	ID        uint   `json:"id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	CreatedAt string `json:"created_at" binding:"required"`
	UpdatedAt string `json:"updated_at" binding:"required"`
}

func (user *User) Dump(userEntity *models.UserModel) *User {
	user.ID = userEntity.ID
	user.Username = userEntity.Username
	user.Email = userEntity.Email
	user.CreatedAt = userEntity.CreatedAt.Format("2006-01-02 15:04:05")
	user.UpdatedAt = userEntity.UpdatedAt.Format("2006-01-02 15:04:05")
	return user
}
