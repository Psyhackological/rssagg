package main

import (
	"fmt"
	"net/http"

	"github.com/Psyhackological/rssagg/internal/auth"
	"github.com/Psyhackological/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(httpResponseWriter http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetAPIKey(request.Header)
		if err != nil {
			respondWithError(httpResponseWriter, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(request.Context(), apiKey)
		if err != nil {
			respondWithError(httpResponseWriter, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(httpResponseWriter, request, user)
	}
}
