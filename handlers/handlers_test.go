package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlers(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "test create shorten",
			want: want{
				code:        200,
				response:    `{"status":"creaete"}`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func (t *testing.T)  {
			request := httptest.NewRequest(http.MethodGet, "/status", nil)
			w := httptest.NewRecorder()
			createShorten()
		})
	}
}
