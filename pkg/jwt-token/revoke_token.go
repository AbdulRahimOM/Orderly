package jwttoken

import (
	"fmt"
	"orderly/internal/domain/constants"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type blacklistEntry struct {
	ExpireAt      time.Time
	ExceptionJTIs []string
}

const (
	expirationDuration = constants.DefaultTokenExpiry
)

var (
	tokenBlacklist = make(map[string]blacklistEntry)
	mutex          = &sync.Mutex{}
)

func init() {
	//start a cron job to clean up expired tokens
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", cleanupExpiredTokens)
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	c.Start()
	fmt.Println("Blacklisted-Tokens cleanup scheduled")
}

func RevokeExistingAuthToken(userID string) {
	mutex.Lock()
	defer mutex.Unlock()

	expirationTime := time.Now().Add(expirationDuration)
	tokenBlacklist[userID] = blacklistEntry{
		ExpireAt: expirationTime,
	}
}

func RegisterExceptionJTI(userID string, jti string) {
	mutex.Lock()
	defer mutex.Unlock()

	entry, exists := tokenBlacklist[userID]
	if !exists {
		return
	}

	entry.ExceptionJTIs = append(entry.ExceptionJTIs, jti)
	tokenBlacklist[userID] = entry
}

// Checks if a token is blacklisted
func isTokenBlacklisted(userID string, jti string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	entry, exists := tokenBlacklist[userID]
	if !exists {
		return false
	}

	for _, exceptionJTI := range entry.ExceptionJTIs {
		if exceptionJTI == jti {
			return false
		}
	}

	return true
}

// Cleans up expired tokens from the map
func cleanupExpiredTokens() {
	mutex.Lock()
	defer mutex.Unlock()

	for userID, entry := range tokenBlacklist {
		if entry.ExpireAt.Before(time.Now()) {
			delete(tokenBlacklist, userID)
		}
	}
}
