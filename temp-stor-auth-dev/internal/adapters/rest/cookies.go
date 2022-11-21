package rest

import (
	"net/http"
	"time"
)

func clearCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "access",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})
}

func sendCookies(access string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    access,
		Expires:  time.Now().Add(time.Minute * 10),
		HttpOnly: true,
		Path:     "/",
	})
}
