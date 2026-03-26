package handlers

import "net/http"

func CreateShortLinkHandler(w http.ResponseWriter, r *http.Request) {

}
func GetShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://yandex.ru/", http.StatusMovedPermanently)
}

func RegisterHandlers() {
	Mux.HandleFunc("/", CreateShortLinkHandler)
	Mux.HandleFunc("/api", GetShortLinkHandler)
}
