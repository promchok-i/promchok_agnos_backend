package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a new router with the route registered
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/strong_password_steps", checkPasswordHandler)
	return router
}

// Helper function to perform a request and return the response recorder
func performRequest(router *gin.Engine, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestStrongPasswordStepsRouteSuccess(t *testing.T) {
	router := setupRouter()

	reqBody := `{"init_password":"Aa1aaabbbb"}`
	w := performRequest(router, http.MethodPost, "/api/strong_password_steps", reqBody)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"num_of_steps":2}`, w.Body.String())
}

func TestStrongPasswordStepsRouteBadRequest(t *testing.T) {
	router := setupRouter()

	reqBody := `{}`
	w := performRequest(router, http.MethodPost, "/api/strong_password_steps", reqBody)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":"Password is required"}`, w.Body.String())
}

func TestStrongPasswordStepsRouteWrongPath(t *testing.T) {
	router := setupRouter()

	reqBody := `{}`
	w := performRequest(router, http.MethodPost, "/api/strong_password_steps_wrong", reqBody)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
