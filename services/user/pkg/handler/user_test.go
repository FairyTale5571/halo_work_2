package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fairytale5571/halo_work_2/services/user/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_user(t *testing.T) {
	tests := []struct {
		name         string
		header       string
		expectedBody string
		expectedCode int
	}{
		{
			name:         "Ok",
			header:       "Username",
			expectedCode: http.StatusOK,
			expectedBody: `{"service":"user-microservice"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewHandler()
			srv := server.InitServer()
			srv.SetHandler(handlers.InitHandlers())
			go func() {
				if err := srv.Run(); err != nil {
					log.Printf("error: %v", err)
				}
			}()

			r := gin.New()
			r.GET("/microservice/name", handlers.microserviceName)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/microservice/name", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

func Test_profile(t *testing.T) {
	tests := []struct {
		name         string
		header       string
		inputHeader  string
		expectedBody string
		expectedCode int
	}{
		{
			name:         "Ok",
			header:       "Username",
			inputHeader:  "user-name",
			expectedCode: http.StatusOK,
			expectedBody: `{"age":"30","dob":"01/01/1990","phone":"1234567890","username":"user-name"}`,
		},
		{
			name:         "Wrong Input",
			header:       "Username",
			inputHeader:  "123",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"message":"invalid username"}`,
		},
		{
			name:         "No Header",
			header:       "",
			inputHeader:  "user-name",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"message":"invalid username"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewHandler()
			srv := server.InitServer()
			srv.SetHandler(handlers.InitHandlers())
			go func() {
				if err := srv.Run(); err != nil {
					log.Printf("error: %v", err)
				}
			}()

			r := gin.New()
			r.GET("/user/profile", handlers.profile, func(c *gin.Context) {
				c.Header(tt.header, tt.inputHeader)
			})
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/user/profile", nil)
			req.Header.Set(tt.header, tt.inputHeader)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}
