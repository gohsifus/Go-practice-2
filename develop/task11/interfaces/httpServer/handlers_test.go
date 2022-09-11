package httpServer

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"task11/interfaces/httpServer/configs"
	"testing"
)

func TestHelloHandler(t *testing.T){
	s, _ := NewServer(configs.NewConfig())
	handler := http.HandlerFunc(s.helloHandler)
	mockW := httptest.NewRecorder()
	mockR, _ := http.NewRequest("GET", "/", nil)
	handler.ServeHTTP(mockW, mockR)
	assert.Equal(t, mockW, []byte("hello"))
}
