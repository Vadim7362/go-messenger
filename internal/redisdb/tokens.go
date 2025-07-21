package redisdb

import (
	"time"
)

const RefreshTokenTTL = 7 * 24 * time.Hour

func SetRefreshToken(userID, token string) error {
	return Client.Set(ctx, userID, token, RefreshTokenTTL).Err()
}

func GetRefreshToken(userID string) (string, error) {
	return Client.Get(ctx, userID).Result()
}

func DeleteRefreshToken(userID string) error {
	return Client.Del(ctx, userID).Err()
}