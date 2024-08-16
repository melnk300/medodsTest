package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/melnk300/medodsTest/internal/jwt"
)

func Server() chi.Router {
	r := chi.NewRouter()

	r.Get("/{guid}", jwt.HandleCreateTokens)
	r.Get("/refresh", jwt.HandleRefreshToken)

	return r
}
