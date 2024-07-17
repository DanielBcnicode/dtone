package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWebTokenService(t *testing.T) {
	apiSecret := "12345678"
	lifeSpan := 10
	idTest := uuid.NewString()
	service, err := NewWebTokenService(apiSecret, lifeSpan)
	if err != nil {
		t.Fatal(err)
	}
	token, err := service.GenerateToken("email@email.com", idTest)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, token)
	assert.GreaterOrEqual(t, len(token), 10)
	assert.Nil(t, service.ValidateToken(token))
}
