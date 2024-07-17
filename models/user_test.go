package models

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestUserModel(t *testing.T) {
	userId := uuid.NewString()
	now := time.Now()
	password := "123456"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	u := User{
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
	assert.Equal(t, "100.00", u.GetBalanceFormatted(), "format incorrect")
	passwordEqual, err := u.CheckPassword(password)
	assert.True(t, passwordEqual, "Password should be equal")
}
