package handlers

import (
	"io"
	"net/http"
	"strings"
)

type ShortenRequest struct {
}

func createShorten(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "Content type must be text/plain", http.StatusUnsupportedMediaType)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	domain := r.Host
	originalURL := string(bodyBytes)
	w.Header().Set("content-type", "text/plain")
	url := "http://" + domain + "/" + originalURL + "ffffff"
	body := strings.ReplaceAll(url, " ", "")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(body))
}

func getShorten(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")
	http.Redirect(w, r, "https://practicum.yandex.ru/", http.StatusTemporaryRedirect)

}

func ShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getShorten(w, r)
	case http.MethodPost:
		createShorten(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RegisterHandlers() {
	Mux.HandleFunc("/", ShortLinkHandler)
}
