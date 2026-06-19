package handlers

import (
	"io"
	"net/http"
	"shortener/store"
	"shortener/utils"
	"strings"
)

type ShortenRequest struct {
}

func createShorten(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "Content type must be text/plain", http.StatusUnsupportedMediaType)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	domain := r.Host
	originalURL := string(bodyBytes)
	if strings.TrimSpace(originalURL) == "" {
		http.Error(w, "Url is empty", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "text/plain")
	code := utils.GenerateShortCodeUrl(originalURL)
	url := "http://" + domain + "/" + code
	store.Store.Store(code, originalURL)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(url))
}

func getShorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	code := strings.TrimPrefix(r.URL.Path, "/")
	val, exists := store.Store.Load(code)
	if !exists {
		http.Error(w, "Short url not found", http.StatusBadRequest)
		return
	}
	originalUrl := val.(string)
	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}

func ShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getShorten(w, r)
	case http.MethodPost:
		createShorten(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}

func RegisterHandlers() {
	Mux.HandleFunc("/", ShortLinkHandler)
}
