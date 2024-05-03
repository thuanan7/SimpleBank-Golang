package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayLoad(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

type MyCustomClaims struct {
	Payload
	jwt.RegisteredClaims
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrInvalidToken
	}
	return nil
}

func NewMyCustomClaims(payload *Payload) *MyCustomClaims {
	return &MyCustomClaims{
		Payload: *payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.IssueAt),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "simple_bank",
			Subject:   payload.Username,
			ID:        payload.ID.String(),
			Audience:  []string{"clients"},
		},
	}
}
