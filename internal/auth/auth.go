package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"uuid"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const ctxKey = "auth.user_id"

type TokenManager struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{secret: []byte(secret), ttl: ttl}
}

// this issues a new token for the current user, and signs it with the HS256 method
func (m *TokenManager) Issue(userID uuid.UUID) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   userID.String(),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.ttl)),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
}

// parse the given token and verify the user id and return either error or the uuid or the user
func (m *TokenManager) Parse(token string) (uuid.UUID, error) {
	claims := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return m.secret, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return uuid.Nil(), fmt.Errorf("invalid token: %w", err)
	}
	return uuid.Parse(claims.Subject)
}

// get the id from the request and return the cut prefix
func (m *TokenManager) fromRequest(c *gin.Context) (uuid.UUID, bool) {
	token, ok := strings.CutPrefix(c.GetHeader("Authorization"), "Bearer")
	if !ok || token == "" {
		return uuid.Nil(), false
	}
	id, err := m.Parse(token)
	if err != nil {
		return uuid.Nil(), false
	}
	return id, true
}

// rejects invalid tokens from the request
func Required(m *TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := m.fromRequest(ctx)
		if !ok {
			log.Println("some internal error occured")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request",
			})
			ctx.Abort()
			return
		}
		ctx.Set(ctxKey, id)
		ctx.Next()
	}
}

// Optional lets anonymous requests through but identifies logged-in users, so
// public GET endpoints can still return "liked_by_me".
func Optional(m *TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		if id, ok := m.fromRequest(c); ok {
			c.Set(ctxKey, id)
		}
		c.Next()
	}
}

// UserID returns the authenticated user, or uuid.Nil for anonymous requests.
func UserID(c *gin.Context) uuid.UUID {
	if v, ok := c.Get(ctxKey); ok {
		if id, ok := v.(uuid.UUID); ok {
			return id
		}
	}
	return uuid.Nil()
}
