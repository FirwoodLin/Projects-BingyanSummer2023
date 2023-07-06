package model

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func InsertSession(userReq *UserRequset, sessionId string) (err error) {
	// 要存储的 session
	// session := map[string]string{"userID": strconv.FormatUint(uint64(userReq.ID), 10), "expiresAt": strconv.FormatUint(uint64(time.Now().Add(time.Hour).Unix()), 10)}
	session := map[string]interface{}{"userID": userReq.ID, "expiresAt": time.Now().Add(time.Hour).Unix()}

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
func CheckSession(cookie string) (userId uint, err error) {
	// 检查 sessionID 是否存在
	sessionID := "session:" + cookie
	if RedisClient.Exists(RedisCtx, sessionID).Val() == 0 {
		return 0, errors.New("不存在 SessionID")
	}
	// 检验是否过期
	expiresAtStr := RedisClient.HGet(RedisCtx, sessionID, "expiresAt").Val()
	expiresAt, _ := strconv.ParseInt(expiresAtStr, 10, 64)
	if expiresAt < time.Now().Unix() {
		log.Printf("[info]model-SessionID已经过期")
		return 0, errors.New("SessionID已经过期")
	}
	userIdStr := RedisClient.HGet(RedisCtx, sessionID, "userID").Val()
	userId64, _ := strconv.ParseUint(userIdStr, 10, 64)
	userId = uint(userId64)
	return userId, nil
}
