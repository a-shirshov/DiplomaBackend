package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConnectionString(t *testing.T) {
	expectedResult := "postgresql://user:password@host:port/db"
	user := "user"
	password := "password"
	host := "host"
	port := "port"
	db := "db"
	result := buildConnectionString(user,password,host,port,db)
	assert.Equal(t, expectedResult,result)
}
