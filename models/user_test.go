package models

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func getTestUser() User {
	userId := uuid.NewString()
	now := time.Now()
	password := "123456"
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return User{
		Base: Base{
			ID:        userId,
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: nil,
		},
		Email:     "test@test.com",
		Password:  string(passwordHash),
		Name:      "Name",
		Telephone: "+34666555444",
		Balance:   10000,
	}
}
func TestUserModel(t *testing.T) {
	u := getTestUser()
	assert.Equal(t, "100.00", u.GetBalanceFormatted(), "format incorrect")
	passwordEqual, err := u.CheckPassword("123456")
	if err != nil {
		t.Error(err)
	}
	assert.True(t, passwordEqual, "Password should be equal")
}

func TestUserModelTopUp(t *testing.T) {
	u := getTestUser()
	err := u.TopUpFromString("10.01")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, u.Balance, int64(11001))
	err = u.TopUpFromString("-10.01")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, u.Balance, int64(10000))
}
