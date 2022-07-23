package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
    AccessUuid string
    UserId   int
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
	   return err
	}
	
	if !token.Valid {
	   return err
	}
	return nil
  }

func ExtractToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   //Make sure that the token method conform to "SigningMethodHMAC"
	   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	   }
	   return []byte(viper.GetString("ACCESS_TOKEN")), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
  }

func CreateToken(userId int) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(viper.GetString("ACCESS_TOKEN")))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(viper.GetString("REFRESH_TOKEN")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
	   return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
	   accessUuid, ok := claims["access_uuid"].(string)
	   if !ok {
		  return nil, errors.New("type problem")
	   }

	   userId, err := strconv.ParseInt(fmt.Sprintf("%.f",claims["user_id"]),10,0)
	   if err != nil {
		  return nil, err
	   } 

	   return &AccessDetails{
		  AccessUuid: accessUuid,
		  UserId:   int(userId),
	   }, nil
	}
	return nil, err
  }