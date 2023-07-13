package model

import (
	"OnlineShop/model/response"
	"errors"
	"log"
	"strconv"
	"time"
)

func InsertSession(userResponse *response.UserSignInResponse, sessionId string) (err error) {
	// 要存储的 session
	// session := map[string]string{"userID": strconv.FormatUint(uint64(userReq.UserID), 10), "expiresAt": strconv.FormatUint(uint64(time.Now().Add(time.Hour).Unix()), 10)}
	session := map[string]interface{}{
		"userID":    userResponse.ID,
		"expiresAt": time.Now().Add(time.Hour).Unix(),
		"isAdmin":   userResponse.IsAdmin}

	// session 的 HashKey
	sessionHashKey := "session:" + sessionId
	for k, v := range session {
		if err := RedisClient.HSet(RedisCtx, sessionHashKey, k, v).Err(); err != nil {
			log.Printf("[error]model-session:%v", err)
			return err
		}
	}
	return nil
}
func CheckSession(cookie string) (*response.UserSignInResponse, error) {
	// 检查 sessionID 是否存在
	userResponse := &response.UserSignInResponse{}
	sessionID := "session:" + cookie
	if RedisClient.Exists(RedisCtx, sessionID).Val() == 0 {
		return nil, errors.New("不存在 SessionID")
	}
	// 检验是否过期
	expiresAtStr := RedisClient.HGet(RedisCtx, sessionID, "expiresAt").Val()
	expiresAt, _ := strconv.ParseInt(expiresAtStr, 10, 64)
	if expiresAt < time.Now().Unix() {
		log.Printf("[info]model-SessionID已经过期")
		return nil, errors.New("SessionID已经过期")
	}
	userId64, _ := RedisClient.HGet(RedisCtx, sessionID, "userID").Uint64()
	userResponse.ID = uint(userId64)
	// 检查是否是管理员
	userResponse.IsAdmin, _ = RedisClient.HGet(RedisCtx, sessionID, "isAdmin").Bool()
	// 每次成功校验，进行 Session 的延期
	if err := RedisClient.HSet(RedisCtx, sessionID, "expiresAt", time.Now().Add(time.Hour).Unix()).Err(); err != nil {
		log.Printf("[error]model-CheckSession:延期错误%v\n", err.Error())
		return nil, err
	}
	return userResponse, nil
}
