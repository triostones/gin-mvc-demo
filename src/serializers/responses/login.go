package responses

import (
	"github.com/triostones/gin-mvc-demo/src/models"
	"github.com/triostones/gin-mvc-demo/src/serializers"
	"github.com/triostones/gin-mvc-demo/src/utils"
)

type LoginResponse struct {
	serializers.User `json:"user"`
	AccessToken      string `json:"access_token"`
}

func (login *LoginResponse) Dump(user *models.UserModel) *LoginResponse {
	login.User.Dump(user)
	access_token, _ := utils.GenerateAccessToken(user.ID, user.Email)
	login.AccessToken = access_token
	return login
}
