package responses

import (
	"github.com/triostones/triostones-backend-gin/src/entity"
	"github.com/triostones/triostones-backend-gin/src/serializers"
	"github.com/triostones/triostones-backend-gin/src/utils"
)

type LoginResponse struct {
	serializers.User `json:"user"`
	AccessToken      string `json:"access_token"`
}

func (login *LoginResponse) Dump(user *entity.UserEntity) *LoginResponse {
	login.User.Dump(user)
	access_token, _ := utils.GenerateAccessToken(user.Email)
	login.AccessToken = access_token
	return login
}