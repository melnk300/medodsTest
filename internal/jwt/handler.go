package jwt

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/melnk300/medodsTest/pkg/clientip"
	"net/http"
)

func HandleCreateTokens(w http.ResponseWriter, r *http.Request) {
	ip := clientip.ProcessClientIp(r)
	guid := chi.URLParam(r, "guid")
	access, refresh := CreateTokens(ip, guid)

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    access.Value,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh.Value,
		Expires:  refresh.Expiration,
		HttpOnly: true,
	})
}

func HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	accessCookie, err := r.Cookie("access_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "Access token is missing", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error retrieving access token", http.StatusBadRequest)
		return
	}

	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "Refresh token is missing", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error retrieving refresh token", http.StatusBadRequest)
		return
	}

	accessToken := accessCookie.Value
	refreshToken := refreshCookie.Value

	ip := clientip.ProcessClientIp(r)

	acToken, rfToken, err := ProcessTokens(accessToken, refreshToken, ip)
	if err != nil {
		switch err.Error() {
		case "invalid payload":
			http.Error(w, "Wrong data structure", http.StatusBadRequest)
			return
		case "invalid token":
			http.Error(w, "Not valid token", http.StatusUnauthorized)
			return
		case "different tokens":
			http.Error(w, "Not valid token pair", http.StatusUnauthorized)
			return
		case "token used":
			http.Error(w, "Token used before", http.StatusUnauthorized)
			return
		case "different ip":
			http.Error(w, "You're location is blocked", http.StatusUnauthorized)
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    acToken.Value,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    rfToken.Value,
		Expires:  rfToken.Expiration,
		HttpOnly: true,
	})
}
