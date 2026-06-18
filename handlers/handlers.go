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
	w.Header().Set("content-type", "text/plain")
	code := utils.GenerateShortCodeUrl(originalURL)
	url := "http://" + domain + "/" + code
	body := strings.ReplaceAll(url, " ", "")
	store.Store.Store(code, originalURL)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(body))
}

func getShorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	path := r.URL.Path
	shortUrl := "http://" + r.Host + string(path)
	val, exists := store.Store.Load(shortUrl)
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
