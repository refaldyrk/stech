package tests

import (
	"aidanwoods.dev/go-paseto"
	"bytes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"kreditplus-test/helper"
	"testing"
)

func TestGenerateRandomLimit(t *testing.T) {
	// Generate multiple random limits to test the range
	for i := 0; i < 1000; i++ {
		result := helper.GenerateRandomLimit()

		if result < 1000000 || result >= 10000000 {
			t.Errorf("GenerateRandomLimit() = %v; want a value between 1,000,000 and 10,000,000", result)
		}
	}
}

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	helper.Log(true, "Something went wrong", "/api/test", "user123", 500)

	logContent := buf.String()
	assert.Contains(t, logContent, "error")                // Log level should be 'error'
	assert.Contains(t, logContent, "/api/test")            // Endpoint
	assert.Contains(t, logContent, "Something went wrong") // Message
	assert.Contains(t, logContent, "user123")              // UserID
	assert.Contains(t, logContent, "500")                  // StatusCode
	assert.Contains(t, logContent, "timestamp")            // Timestamp
	assert.Contains(t, logContent, "caller")               // Caller info

	// Reset the buffer for the next test
	buf.Reset()

	// Test case 2: Log without error
	helper.Log(false, "Operation completed", "/api/test", "user456", 200)

	// Verify if info log contains expected fields
	logContent = buf.String()
	assert.Contains(t, logContent, "info")                // Log level should be 'info'
	assert.Contains(t, logContent, "/api/test")           // Endpoint
	assert.Contains(t, logContent, "Operation completed") // Message
	assert.Contains(t, logContent, "user456")             // UserID
	assert.Contains(t, logContent, "200")                 // StatusCode
	assert.Contains(t, logContent, "timestamp")           // Timestamp
	assert.Contains(t, logContent, "caller")              // Caller info
}

func TestGeneratePaseto(t *testing.T) {
	viper.Set("PASETO_AUDIENCE", "test_audience")
	viper.Set("PASETO_ISSUER", "test_issuer")
	viper.Set("PASETO_SUBJECT", "test_subject")
	viper.Set("PASETO_SECRET", "test_secret")

	key := paseto.NewV4SymmetricKey()

	claims := map[string]interface{}{
		"user_id": "user123",
	}

	token, err := helper.GeneratePaseto(key, claims)
	require.NoError(t, err)

	assert.NotEmpty(t, token)

	parsedClaims, err := helper.ValidatePaseto(key, token)
	require.NoError(t, err)

	assert.Equal(t, "user123", parsedClaims["user_id"])
}

func TestHashPassword(t *testing.T) {
	password := "mySecurePassword"
	hashedPassword, err := helper.HashPassword(password)
	require.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	hashedPassword2, err := helper.HashPassword(password)
	require.NoError(t, err)
	assert.NotEmpty(t, hashedPassword2)
	assert.NotEqual(t, hashedPassword, hashedPassword2)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mySecurePassword"
	hashedPassword, err := helper.HashPassword(password)
	require.NoError(t, err)

	isValid := helper.CheckPasswordHash(password, hashedPassword)
	assert.True(t, isValid, "Password should match the hash")

	invalidPassword := "wrongPassword"
	isValid = helper.CheckPasswordHash(invalidPassword, hashedPassword)
	assert.False(t, isValid, "Password should not match the hash")
}

func TestResponseAPI(t *testing.T) {
	successResponse := helper.ResponseAPI(true, 200, "Request berhasil", map[string]string{"key": "value"})
	assert.True(t, successResponse.Meta.Success)
	assert.Equal(t, 200, successResponse.Meta.Code)
	assert.Equal(t, "Request berhasil", successResponse.Meta.Message)
	assert.Equal(t, map[string]string{"key": "value"}, successResponse.Data)

	errorResponse := helper.ResponseAPI(false, 400, "Bad request", nil)
	assert.False(t, errorResponse.Meta.Success)
	assert.Equal(t, 400, errorResponse.Meta.Code)
	assert.Equal(t, "Bad request", errorResponse.Meta.Message)
	assert.Nil(t, errorResponse.Data)
}
