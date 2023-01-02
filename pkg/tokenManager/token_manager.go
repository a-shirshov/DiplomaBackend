package tokenManager

import (
	"Diploma/internal/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type tokenManager struct {

}

func NewTokenManager() *tokenManager{
	return &tokenManager{

	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
	   return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
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

func(tM *tokenManager) CreateToken(userId int) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}
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

func(tM *tokenManager) ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := verifyToken(r)
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

	   return &models.AccessDetails{
		  AccessUuid: accessUuid,
		  UserId:   int(userId),
	   }, nil
	}
	return nil, err
}