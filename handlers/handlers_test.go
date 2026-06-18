package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_create_handlers(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "test create shorten",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://www.yandex.ru"))
			request.Header.Set("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			createShorten(w, request)
			res := w.Result()
			assert.Equal(t, res.StatusCode, test.want.code)
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Regexp(t, `^http://example\.com/[A-Za-z0-9+/]{8}$`, string(resBody))
			assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)
		})
	}
}

func Test_get_handlers(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "test create shorten",
			want: want{
				code:        307,
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://www.yandex.ru"))
			request.Header.Set("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			createShorten(w, request)
			res := w.Result()
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			gw := httptest.NewRecorder()
			getRequest := httptest.NewRequest(http.MethodGet, string(resBody), nil)
			getRequest.Header.Set("Content-Type", "text/plain")
			getShorten(gw, getRequest)
			gres := gw.Result()
			defer gres.Body.Close()
			_, errgres := io.ReadAll(gres.Body)
			require.NoError(t, errgres)
			assert.Equal(t, test.want.code, gres.StatusCode)
		})
	}
}
