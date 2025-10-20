package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, router := gin.CreateTestContext(w)

	router.GET("/health", HealthCheckHandler)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(http.StatusOK, w.Code)

	expectedBody := `{"status":"UP"}`
	assert.JSONEq(expectedBody, w.Body.String())
}
