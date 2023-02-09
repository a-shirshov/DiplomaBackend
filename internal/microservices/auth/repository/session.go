package repository

import (
	"Diploma/internal/models"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type passwordRedeemStruct struct {
	RedeemCode int
	NumberOfFailedAttempts int
	AcceptedToChange bool
}

func newPasswordRedeem(redeemCode int) (*passwordRedeemStruct) {
	return &passwordRedeemStruct{
		RedeemCode: redeemCode,
		NumberOfFailedAttempts: 0,
		AcceptedToChange: false,
	}
}

type SessionRepository struct {
	redis *redis.Client
}

func NewSessionRepository (redis *redis.Client) (*SessionRepository) {
	return &SessionRepository{
		redis: redis,
	}
}

func (sR *SessionRepository) SaveTokens(userId int, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) 
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := sR.redis.Set(td.AccessUuid, strconv.Itoa(userId), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := sR.redis.Set(td.RefreshUuid, strconv.Itoa(userId), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	
	return nil
}

func (sR *SessionRepository) FetchAuth(accessUuid string) (int, error) {
	userid, err := sR.redis.Get(accessUuid).Result()
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

func (sR *SessionRepository) SavePasswordRedeemCode(email string, redeemCode int) error {
	redeemStruct := newPasswordRedeem(redeemCode)
	redeemStructJSON, err := json.Marshal(redeemStruct)
	if err != nil {
		return err
	}
	return sR.redis.Set(email, redeemStructJSON, time.Hour).Err()
}

func (sR *SessionRepository) CheckRedeemCode(email string, redeemCode int) error {
	redeemStructJSON, err := sR.redis.Get(email).Result()
	if err != nil {
		return errors.New("something went wrong")
	}
	var redeemStruct passwordRedeemStruct
	err = json.Unmarshal([]byte(redeemStructJSON), &redeemStruct)
	if err != nil {
		return err
	}

	if (redeemCode == redeemStruct.RedeemCode) {
		redeemStruct.AcceptedToChange = true
		saveJSON, err := json.Marshal(redeemStruct)
		if err != nil {
			return err
		}
		err = sR.redis.Set(email, saveJSON, time.Hour).Err()
		if err != nil {
			return errors.New("something went wrong")
		}
		return nil
	}

	redeemStruct.NumberOfFailedAttempts++
	if redeemStruct.NumberOfFailedAttempts >= 3 {
		_, err := sR.redis.Del(email).Result()
		if err != nil {
			return errors.New("something went wrong")
		}
		return errors.New("redeem code has been expired")
	}

	saveJSON, err := json.Marshal(redeemStruct)
	if err != nil {
		return err
	}

	err = sR.redis.Set(email, saveJSON, time.Hour).Err()
	if err != nil {
		return err
	}

	return errors.New("wrong redeem code")
}

func (sR *SessionRepository) CheckAccessToNewPassword(email string) (bool) {
	isAccepted := false
	redeemStructJSON, err := sR.redis.Get(email).Result()
	if err != nil {
		return isAccepted
	}

	var redeemStruct passwordRedeemStruct
	err = json.Unmarshal([]byte(redeemStructJSON), &redeemStruct)
	if err != nil {
		return isAccepted
	}

	if redeemStruct.AcceptedToChange {
		_, err := sR.redis.Del(email).Result()
		if err != nil {
			return false
		}
		return redeemStruct.AcceptedToChange
	}

	return false
}