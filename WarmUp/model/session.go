package model

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func InsertSession(userResponse *UserSignInResponse, sessionId string) (err error) {
	// 要存储的 session
	// session := map[string]string{"userID": strconv.FormatUint(uint64(userReq.ID), 10), "expiresAt": strconv.FormatUint(uint64(time.Now().Add(time.Hour).Unix()), 10)}
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
func CheckSession(cookie string) (*UserSignInResponse, error) {
	// 检查 sessionID 是否存在
	userResponse := &UserSignInResponse{}
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
	// userIdStr := RedisClient.HGet(RedisCtx, sessionID, "userID").Val()
	// userId64, _ := strconv.ParseUint(userIdStr, 10, 64)
	userId64, _ := RedisClient.HGet(RedisCtx, sessionID, "userID").Uint64()

	userResponse.ID = uint(userId64)
	// 检查是否是管理员
	userResponse.IsAdmin, _ = RedisClient.HGet(RedisCtx, sessionID, "isAdmin").Bool()

	return userResponse, nil
}
