package jwt

import (
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
		Expires:  access.Expiration,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh.Value,
		Expires:  refresh.Expiration,
		HttpOnly: true,
	})

}
