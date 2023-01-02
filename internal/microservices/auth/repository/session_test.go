package repository

import (
	"Diploma/internal/models"
	"strconv"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

type saveTokensTest struct {
	userId int
	td *models.TokenDetails
}

type FetchAuthTest struct {
	accessUuid string
	UserId int
}

type DeleteAuthTest struct {
	accessUuid string
	UserId int
}

var saveTokensTests = []saveTokensTest{
	{
		1, &models.TokenDetails{
			AccessToken: "accessToken_1",
			RefreshToken: "refreshToken_1",
		},
	},
}

var FetchAuthTests = []FetchAuthTest{
	{
		"accessUUID_1", 1,
	},
}

var DeleteAuthTests = []DeleteAuthTest{
	{
		"accessUUID_1", 1,
	},
}

func TestSaveTokens(t *testing.T) {
	testRedis := miniredis.RunT(t)

	redisMock := redis.NewClient(&redis.Options{
		Addr: testRedis.Addr(),
	})

	repositoryTest := NewSessionRepository(redisMock)
	for _, test := range saveTokensTests {
		test.td.AccessUuid = uuid.NewV4().String()
		test.td.RefreshUuid = uuid.NewV4().String()
		test.td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
		test.td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

		repositoryTest.SaveTokens(test.userId, test.td)

		testRedis.CheckGet(t, test.td.AccessUuid, strconv.Itoa(test.userId))
		testRedis.CheckGet(t, test.td.RefreshUuid, strconv.Itoa(test.userId))

		at := time.Unix(test.td.AtExpires, 0) 
		rt := time.Unix(test.td.RtExpires, 0)
		now := time.Now()

		testRedis.FastForward(at.Sub(now) + time.Second)
		if testRedis.Exists(test.td.AccessUuid) {
			t.Fatal("AtExpires should not exist")
		}

		testRedis.FastForward(rt.Sub(now) + time.Second)
		if testRedis.Exists(test.td.RefreshUuid) {
			t.Fatal("RtExpires should not exist")
		}
	}
}

func TestFetchAuth(t *testing.T) {
	testRedis := miniredis.RunT(t)

	redisMock := redis.NewClient(&redis.Options{
		Addr: testRedis.Addr(),
	})

	repositoryTest := NewSessionRepository(redisMock)
	for _, test := range FetchAuthTests {
		testRedis.Set(test.accessUuid, strconv.Itoa(test.UserId))
		resultUserId, err := repositoryTest.FetchAuth(test.accessUuid)
		assert.Equal(t, test.UserId, resultUserId)
		assert.Nil(t, err)
	}
}

func TestDeleteAuth(t *testing.T) {
	testRedis := miniredis.RunT(t)

	redisMock := redis.NewClient(&redis.Options{
		Addr: testRedis.Addr(),
	})

	repositoryTest := NewSessionRepository(redisMock)
	for _, test := range DeleteAuthTests {
		testRedis.Set(test.accessUuid, strconv.Itoa(test.UserId))
		repositoryTest.DeleteAuth(test.accessUuid)
		if testRedis.Exists(test.accessUuid) {
			t.Fatal("accessUuid should not exist")
		}
	}
}