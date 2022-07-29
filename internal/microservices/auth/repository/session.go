package repository

import (
	"Diploma/utils"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type SessionRepository struct {
	redis *redis.Client
}

func NewSessionRepository (redis *redis.Client) (*SessionRepository) {
	return &SessionRepository{
		redis: redis,
	}
}

func (sR *SessionRepository) SaveTokens(userId int, td *utils.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) 
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := sR.redis.Set(td.AccessUuid, userId, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := sR.redis.Set(td.RefreshUuid, userId, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	
	return nil
}

func (sR *SessionRepository) FetchAuth(accessToken string) (int, error) {
	userid, err := sR.redis.Get(accessToken).Result()
	if err != nil {
	   return 0, err
	}
	userID, _ := strconv.Atoi(userid)
	return userID, nil
}

func (sR *SessionRepository) DeleteAuth(accessUuid string) error {
	_, err := sR.redis.Del(accessUuid).Result()
	if err != nil {
		return err
	}
	return nil
}