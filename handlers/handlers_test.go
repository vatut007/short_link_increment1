package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/status", nil)
			w := httptest.NewRecorder()
			createShorten(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, test.want.code)
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.JSONEq(t, string(resBody), test.want.response)
			assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)
		})
	}
}
