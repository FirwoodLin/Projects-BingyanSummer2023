package util

import (
	"OnlineShop/model"
	"OnlineShop/model/response"
	"github.com/google/uuid"
)

func GenSessionId(userResponse *response.UserSignInResponse) (sessionID string, err error) {
	// user := model.ConvertUserRequestToUser(userReq)
	sessionID = uuid.New().String() // 使用 uuid 作为 sessinID
	if err := model.InsertSession(userResponse, sessionID); err != nil {
		return "", err
	}
	return sessionID, nil
}
