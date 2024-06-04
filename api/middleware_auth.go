package main

import (
	"fmt"
	"net/http"

	"github.com/basel2053/go-rss/internal/auth"
	"github.com/basel2053/go-rss/internal/database"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Error getting API key: %s", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %s", err))
			return
		}
		handler(w, r, user)
	}
}
