package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fairytale5571/halo_work_2/services/auth/pkg/server"
	"github.com/fairytale5571/halo_work_2/services/auth/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_auth(t *testing.T) {
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
			expectedBody: `{"message":"ok"}`,
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
			expectedBody: `{"message":"empty auth header"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			services := service.NewService()
			handlers := NewHandler(services)
			srv := server.InitServer()
			srv.SetHandler(handlers.InitHandlers())
			go func() {
				if err := srv.Run(); err != nil {
					log.Printf("error: %v", err)
				}
			}()

			r := gin.New()
			r.GET("/auth", handlers.auth, func(c *gin.Context) {
				c.Header(tt.header, tt.inputHeader)
			})
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/auth", nil)
			req.Header.Set(tt.header, tt.inputHeader)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}
